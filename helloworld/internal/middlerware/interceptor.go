package middlerware

import (
	"context"
	"runtime"

	"github.com/game1991/layout/helloworld/internal/pkg/ecode"
	"github.com/game1991/layout/helloworld/pkg/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcUseUnaryIncerceptor  一元拦截器中间件
func GrpcUseUnaryIncerceptor(ui ...grpc.UnaryServerInterceptor) []grpc.UnaryServerInterceptor {
	return ui
}

// GrpcUseStreamIncerceptor 流式拦截器中间件
func GrpcUseStreamIncerceptor(ui ...grpc.StreamServerInterceptor) []grpc.StreamServerInterceptor {
	return ui
}

// GrpcRecovery ...
func GrpcRecovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		panicked := true
		defer func() {
			if rerr := recover(); rerr != nil || panicked {
				buf := make([]byte, 64<<10) //nolint:gomnd
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				log.Errorf("%s invoke panic %v: %+v\n%s\n", info.FullMethod, rerr, req, buf)
				err = status.Errorf(codes.Internal, "%v", rerr)
			}
		}()

		resp, err = handler(ctx, req)
		panicked = false
		return
	}

}

// GrpcError 针对service层的error进行封装为标准grpc error
func GrpcError() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		resp, err = handler(ctx, req)
		if err != nil {
			err = ecode.ConvertToGrpcErr(err)
			log.Errorf("%s err:%v\n", info.FullMethod, err)
		}
		return
	}
}
