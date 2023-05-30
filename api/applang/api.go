package applang

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
	"google.golang.org/grpc"
)

type Server struct {
	applang.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	applang.RegisterMiddlewareServer(server, &Server{})
}
