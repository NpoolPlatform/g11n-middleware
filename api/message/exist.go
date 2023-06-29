//nolint:nolintlint,dupl
package message

import (
	"context"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/message"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistMessageConds(
	ctx context.Context,
	in *npool.ExistMessageCondsRequest,
) (
	*npool.ExistMessageCondsResponse,
	error,
) {
	handler, err := message1.NewHandler(ctx,
		message1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMessageConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistMessageCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistMessageConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMessageConds",
			"Req", in,
			"Error", err,
		)
		return &npool.ExistMessageCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistMessageCondsResponse{
		Info: info,
	}, nil
}
