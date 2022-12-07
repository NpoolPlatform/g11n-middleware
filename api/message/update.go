//nolint:nolintlint,dupl
package message

import (
	"context"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	messagemgrtracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/message1"

	"github.com/google/uuid"
)

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateMessage")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = messagemgrtracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateMessage", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Update")

	info, err := message1.UpdateMessage(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateMessage", "error", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateMessageResponse{
		Info: info,
	}, nil
}
