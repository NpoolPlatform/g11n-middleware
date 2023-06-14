package applang

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	applang.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	applang.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
