//nolint:nolintlint,dupl
package appcountry

import (
	"context"

	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/appcountry"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	appcountrymgrapi "github.com/NpoolPlatform/g11n-manager/api/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-middleware/pkg/appcountry"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCountry")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = appcountrymgrapi.Validate(in.GetInfo())
	if err != nil {
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcountry", "crud", "Create")

	info, err := appcountry1.CreateCountry(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCountry", "error", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountryResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	for _, info := range in.GetInfos() {
		if err := appcountrymgrapi.Validate(info); err != nil {
			return &npool.CreateCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	infos, err := appcountry1.CreateCountries(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateCountries", "error", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountriesResponse{
		Infos: infos,
	}, nil
}
