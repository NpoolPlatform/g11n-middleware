//nolint:nolintlint,dupl
package message

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/message"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	req := in.GetInfo()
	handler, err := message1.NewHandler(
		ctx,
		message1.WithEntID(req.EntID, false),
		message1.WithAppID(req.AppID, true),
		message1.WithLangID(req.LangID, true),
		message1.WithMessageID(req.MessageID, true),
		message1.WithMessage(req.Message, true),
		message1.WithGetIndex(req.GetIndex, false),
		message1.WithDisabled(req.Disabled, false),
	)

	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessages",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateMessages(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessages",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateMessagesResponse{
		Infos: infos,
	}, nil
}
