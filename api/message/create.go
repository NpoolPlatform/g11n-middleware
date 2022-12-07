//nolint:nolintlint,dupl
package message

import (
	"context"

	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/message"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	messagemgrapi "github.com/NpoolPlatform/g11n-manager/api/message"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/message1"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateMessage")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = messagemgrapi.Validate(in.GetInfo())
	if err != nil {
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Create")

	info, err := message1.CreateMessage(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateMessage", "error", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	for _, info := range in.GetInfos() {
		if err := messagemgrapi.Validate(info); err != nil {
			return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	infos, err := message1.CreateMessages(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessagesResponse{
		Infos: infos,
	}, nil
}
