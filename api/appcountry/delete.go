//nolint:nolintlint,dupl
package appcountry

import (
	"context"

	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-middleware/pkg/appcountry"

	"github.com/google/uuid"
)

func (s *Server) DeleteCountry(ctx context.Context, in *npool.DeleteCountryRequest) (*npool.DeleteCountryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCountry")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteCountry", "ID", in.GetID(), "error", err)
		return &npool.DeleteCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcountry", "mw", "Delete")

	info, err := appcountry1.DeleteCountry(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteCountry", "error", err)
		return &npool.DeleteCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCountryResponse{
		Info: info,
	}, nil
}
