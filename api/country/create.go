//nolint:nolintlint,dupl
package country

import (
	"context"

	country1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/country"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	req := in.GetInfo()
	handler, err := country1.NewHandler(
		ctx,
		country1.WithID(req.ID),
		country1.WithCountry(req.Country),
		country1.WithFlag(req.Flag),
		country1.WithCode(req.Code),
		country1.WithShort(req.Short),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateCountryResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountries",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateCountries(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountries",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateCountriesResponse{
		Infos: infos,
	}, nil
}
