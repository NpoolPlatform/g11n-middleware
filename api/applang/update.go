//nolint:nolintlint,dupl
package applang

import (
	"context"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	applangmgrtracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/applang"
	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"github.com/google/uuid"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateLang")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = applangmgrtracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateLang", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "applang", "crud", "Update")

	info, err := applang1.UpdateLang(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateLang", "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateLangResponse{
		Info: info,
	}, nil
}
