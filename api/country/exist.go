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

func (s *Server) ExistCountryConds(
	ctx context.Context,
	in *npool.ExistCountryCondsRequest,
) (
	*npool.ExistCountryCondsResponse,
	error,
) {
	handler, err := country1.NewHandler(ctx,
		country1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCountryConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistCountryCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCountryConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCountryConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistCountryCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCountryCondsResponse{
		Info: info,
	}, nil
}
