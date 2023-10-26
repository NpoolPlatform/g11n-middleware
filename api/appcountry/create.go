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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCountry",
			"In", in,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithEntID(req.EntID, false),
		appcountry1.WithAppID(req.AppID, true),
		appcountry1.WithCountryID(req.CountryID, true),
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
		appcountry1.WithReqs(in.GetInfos(), true),
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
