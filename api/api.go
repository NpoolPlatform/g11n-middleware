package api

import (
	"context"

	g11n "github.com/NpoolPlatform/message/npool/g11n/mw/v1"

	"github.com/NpoolPlatform/g11n-middleware/api/appcountry"
	"github.com/NpoolPlatform/g11n-middleware/api/applang"
	"github.com/NpoolPlatform/g11n-middleware/api/country"
	"github.com/NpoolPlatform/g11n-middleware/api/lang"
	"github.com/NpoolPlatform/g11n-middleware/api/message"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	g11n.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	g11n.RegisterMiddlewareServer(server, &Server{})
	applang.Register(server)
	lang.Register(server)
	appcountry.Register(server)
	country.Register(server)
	message.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := g11n.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}
