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

func (s *Server) DeleteMessage(ctx context.Context, in *npool.DeleteMessageRequest) (*npool.DeleteMessageResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
		)
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := message1.NewHandler(
		ctx,
		message1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteMessageResponse{
		Info: info,
	}, nil
}
