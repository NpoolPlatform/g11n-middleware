package country

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	country.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	country.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
