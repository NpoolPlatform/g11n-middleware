package country

import (
	"github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
	"google.golang.org/grpc"
)

type Server struct {
	country.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	country.RegisterMiddlewareServer(server, &Server{})
}
