package appcountry

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appcountry.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appcountry.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
