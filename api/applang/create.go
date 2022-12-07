//nolint:nolintlint,dupl
package applang

import (
	"context"

	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/applang"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	applangmgrapi "github.com/NpoolPlatform/g11n-manager/api/applang"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/applang"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateLang")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = applangmgrapi.Validate(in.GetInfo())
	if err != nil {
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "applang", "crud", "Create")

	info, err := applang1.CreateLang(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateLangs(ctx context.Context, in *npool.CreateLangsRequest) (*npool.CreateLangsResponse, error) {
	return nil, nil
}
