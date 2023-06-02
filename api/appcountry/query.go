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

func (s *Server) GetCountryOnly(ctx context.Context, in *npool.GetCountryOnlyRequest) (*npool.GetCountryOnlyResponse, error) {
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountry",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountryOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountry",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountryOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCountryOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithConds(in.GetConds()),
		appcountry1.WithOffset(in.GetOffset()),
		appcountry1.WithLimit(in.GetLimit()),
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
