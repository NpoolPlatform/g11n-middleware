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

func (s *Server) UpdateCountry(ctx context.Context, in *npool.UpdateCountryRequest) (*npool.UpdateCountryResponse, error) {
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
			"UpdateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateCountryResponse{
		Info: info,
	}, nil
}
