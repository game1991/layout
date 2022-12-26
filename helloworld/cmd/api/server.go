package api

import (
	"context"
	"errors"
	v1 "helloworld/api/proto/v1"
	"helloworld/build"
	"helloworld/internal/conf"
	"helloworld/internal/controller"
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

		//TODO 读取配置文件
		if err := conf.InitConfig(); err != nil {
			panic("InitConfig" + err.Error())
		}
		logConf, err := conf.Log()
		if err != nil {
			log.Panic("logConf", log.FieldErr(err))
		}
		log.InitLogger(*logConf)
		serverConf, err := conf.GetServer()
		if err != nil {
			log.Panic("serverConf", log.FieldErr(err))
		}
		mysqlConf, err := conf.MySQL()
		if err != nil {
			log.Panic("mysql", log.FieldErr(err))
		}
		app, cleanup, err := wireApp(serverConf, *mysqlConf)
		if err != nil {
			log.Errorf("wireApp error : %v", err)
			return err
		}
		defer cleanup()

		log.Infof("服务启动中 Listening and serving HTTP on %v", serverConf.GetAddr())
		build.BuildInfo()
		log.Infof("SystemInfo:[%v]", build.SystemInfo)
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
	handler *controller.Handler,
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
	/*
		如果启动了 gprc 反射服务，那么就可以通过 reflection 包提供的反射服务查询 gRPC 服务或调用 gRPC 方法。
		grpcurl 是 Go 语言开源社区开发的工具,可以用来验证grpc的服务。
	*/
	reflection.Register(grpcServer)

	/***** 注册你的grpc服务 *****/
	v1.RegisterGreeterServer(grpcServer, greeterSrv)
	v1.RegisterUserServiceServer(grpcServer, userSrv)

	// 初始化一个空Gin路由
	engine := gin.Default()
	/***** 添加你的api路由吧 *****/
	v1Engine := engine.Group("v1")
	handler.APIHandler(o.ctx, v1Engine)
	handler.SYSHandler(o.ctx, engine)

	// 使用grpc+gin模式同一个端口提供服务
	httpHandler, err := server.Serve(grpcServer, engine)
	if err != nil {
		log.Panicf("server.Serve err:%v\n", err)
	}

	ctx, cancel := context.WithCancel(o.ctx)

	return &APP{
		ctx:    ctx,
		sigs:   o.sigs,
		cancel: cancel,
		server: &http.Server{Addr: conf.GetAddr(), Handler: httpHandler},
		addr:   conf.GetAddr(),
	}
}

func newApp(
	conf *conf.Server,
	handler *controller.Handler,
	greeterSrv *service.GreeterSrv,
	userSrv *service.UserSrv,
) *APP {
	return NewApp(
		conf,
		handler,
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
