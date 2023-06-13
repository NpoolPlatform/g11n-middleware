//nolint:nolintlint,dupl
package message

import (
	"context"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (s *Server) GetMessage(ctx context.Context, in *npool.GetMessageRequest) (*npool.GetMessageResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessage",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessage",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessageResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithConds(in.GetConds()),
		message1.WithOffset(in.GetOffset()),
		message1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessages",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessagesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetMessages(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessages",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessagesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMessagesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMessageOnly(ctx context.Context, in *npool.GetMessageOnlyRequest) (*npool.GetMessageOnlyResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessageOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessageOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetMessageOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessageOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessageOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMessageOnlyResponse{
		Info: info,
	}, nil
}
