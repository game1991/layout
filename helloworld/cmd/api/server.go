package api

import (
	"context"
	"errors"
	v1 "helloworld/api/proto/v1"
	"helloworld/build"
	"helloworld/internal/conf"
	"helloworld/internal/controller"
	"helloworld/internal/pkg/store"
	"helloworld/internal/server"
	"helloworld/internal/service"
	"helloworld/pkg/log"
	"helloworld/pkg/uuid"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra" // for cobra.Command
)

// StartCmd cmd args
var StartCmd = &cobra.Command{

	Use:          "serve",
	Short:        "Run the gRPC hello-world server",
	Example:      "cmd serve -d ../configs",
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("Recover error : %v", err)
			}
		}()

		log.InitLog()
		//TODO 读取配置文件

		app, cleanup, err := wireApp(&conf.Server{Addr: ":8081"}, store.Config{
			DSN: "root:root@(127.0.0.1:3306)/helloworld?charset=utf8mb4&parseTime=True&loc=Local",
		})
		if err != nil {
			log.Errorf("wireApp error : %v", err)
			return err
		}

		defer cleanup()

		// start and wait for stop signal
		if err := app.Run(); err != nil {
			panic(err)
		}

		return nil
	},
}

var configToml string

func init() {
	// 启动命令加载config，具体加载规则见utils/config/config.go 中viper的使用
	StartCmd.Flags().StringVarP(&configToml, "dir", "d", "../configs", "config file")
}

// APP ...
type APP struct {
	ctx    context.Context
	sigs   []os.Signal
	cancel func()
	server *http.Server
	addr   string
}

// NewApp .
func NewApp(
	conf *conf.Server,
	greeterSrv *service.GreeterSrv,
	userSrv *service.UserSrv,
	opts ...Option,
) *APP {
	// default option
	o := &options{
		ctx:     context.Background(),
		sigs:    []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		id:      uuid.UUID32(),
		name:    build.Name,
		version: build.Version,
	}
	// apply the option
	for _, opt := range opts {
		opt(o)
	}

	grpcServer := server.NewGRPCServer(conf)
	/***** 注册你的grpc服务 *****/
	v1.RegisterGreeterServer(grpcServer, greeterSrv)
	v1.RegisterUserServiceServer(grpcServer, userSrv)
	reflection.Register(grpcServer)
	// 初始化一个空Gin路由
	engine := gin.Default()
	/***** 添加你的api路由吧 *****/
	controller.NewHandler(userSrv).InstallHandler(o.ctx, engine)

	handler, err := server.Serve(grpcServer, engine)
	if err != nil {
		log.Panicf("server.Serve err:%v\n", err)
	}

	ctx, cancel := context.WithCancel(o.ctx)

	return &APP{
		ctx:    ctx,
		sigs:   o.sigs,
		cancel: cancel,
		server: &http.Server{Addr: conf.GetAddr(), Handler: handler},
		addr:   conf.GetAddr(),
	}
}

func newApp(
	conf *conf.Server,
	greeterSrv *service.GreeterSrv,
	userSrv *service.UserSrv,
) *APP {
	return NewApp(
		conf,
		greeterSrv,
		userSrv,
		Name(build.Name),
		Version(build.Version),
	)
}

// Run .
func (a *APP) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)

	eg.Go(func() error {
		<-ctx.Done() // wait for stop signal
		log.Info("[HTTP] server stopping")
		return a.server.Shutdown(ctx)
	})

	eg.Go(func() error {
		if err := a.server.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case <-c:
			return a.Stop()
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *APP) Stop() error {
	// TODO:服务下线

	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
