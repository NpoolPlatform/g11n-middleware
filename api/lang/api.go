package lang

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	lang.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	lang.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
