//nolint:nolintlint,dupl
package message

import (
	"context"

	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/message1"

	"github.com/google/uuid"
)

func (s *Server) DeleteMessage(ctx context.Context, in *npool.DeleteMessageRequest) (*npool.DeleteMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteMessage")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteMessage", "ID", in.GetID(), "error", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "message", "mw", "Delete")

	info, err := message1.DeleteMessage(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteMessage", "error", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteMessageResponse{
		Info: info,
	}, nil
}
