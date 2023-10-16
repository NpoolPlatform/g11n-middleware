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

func (s *Server) DeleteCountry(ctx context.Context, in *npool.DeleteCountryRequest) (*npool.DeleteCountryResponse, error) {
	req := in.GetInfo()
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCountry",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCountry",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCountryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteCountryResponse{
		Info: info,
	}, nil
}
