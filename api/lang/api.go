package lang

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
	"google.golang.org/grpc"
)

type Server struct {
	lang.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	lang.RegisterMiddlewareServer(server, &Server{})
}
