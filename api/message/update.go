//nolint:nolintlint,dupl
package message

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/message"
)

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	req := in.GetInfo()
	handler, err := message1.NewHandler(
		ctx,
		message1.WithID(req.ID),
		message1.WithAppID(req.AppID),
		message1.WithLangID(req.LangID),
		message1.WithMessageID(req.MessageID),
		message1.WithMessage(req.Message),
		message1.WithGetIndex(req.GetIndex),
		message1.WithDisabled(req.Disabled),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateMessageResponse{
		Info: info,
	}, nil
}
