package appcountry

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
	"google.golang.org/grpc"
)

type Server struct {
	appcountry.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appcountry.RegisterMiddlewareServer(server, &Server{})
}
