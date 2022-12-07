package message

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	message.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	message.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
