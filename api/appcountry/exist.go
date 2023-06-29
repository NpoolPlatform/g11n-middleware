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

func (s *Server) ExistCountryConds(
	ctx context.Context,
	in *npool.ExistCountryCondsRequest,
) (
	*npool.ExistCountryCondsResponse,
	error,
) {
	handler, err := appcountry1.NewHandler(ctx,
		appcountry1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCountryConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistCountryCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistAppCountryConds(ctx)
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
