//nolint:nolintlint,dupl
package lang

import (
	"context"

	lang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/lang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistLangConds(
	ctx context.Context,
	in *npool.ExistLangCondsRequest,
) (
	*npool.ExistLangCondsResponse,
	error,
) {
	handler, err := lang1.NewHandler(ctx,
		lang1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLangConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistLangCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistLangConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLangConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistLangCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistLangCondsResponse{
		Info: info,
	}, nil
}
