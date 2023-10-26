package country

import (
	"context"

	country1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/country"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCountry(ctx context.Context, in *npool.GetCountryRequest) (*npool.GetCountryResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountry",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountry",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCountryResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithConds(in.GetConds()),
		country1.WithOffset(in.GetOffset()),
		country1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountries",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountriesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetCountries(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountries",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCountriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
