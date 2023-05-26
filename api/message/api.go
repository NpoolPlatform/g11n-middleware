package message

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"google.golang.org/grpc"
)

type Server struct {
	message.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	message.RegisterMiddlewareServer(server, &Server{})
}
