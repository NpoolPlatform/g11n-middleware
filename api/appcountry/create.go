//nolint:nolintlint,dupl
package appcountry

import (
	"context"

	appcountry1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/appcountry"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	req := in.GetInfo()
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithID(req.ID),
		appcountry1.WithAppID(req.AppID),
		appcountry1.WithCountryID(req.CountryID),
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
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithReqs(in.GetInfos()),
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
