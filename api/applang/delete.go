//nolint:nolintlint,dupl
package applang

import (
	"context"

	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/applang"

	"github.com/google/uuid"
)

func (s *Server) DeleteLang(ctx context.Context, in *npool.DeleteLangRequest) (*npool.DeleteLangResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteLang")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteLang", "ID", in.GetID(), "error", err)
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "applang", "mw", "Delete")

	info, err := applang1.DeleteLang(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteLang", "error", err)
		return &npool.DeleteLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteLangResponse{
		Info: info,
	}, nil
}
