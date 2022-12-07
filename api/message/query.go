//nolint:nolintlint,dupl
package message

import (
	"context"

	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/message"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant1 "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	message1 "github.com/NpoolPlatform/g11n-middleware/pkg/message1"

	"github.com/google/uuid"
)

func (s *Server) GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetMessages")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "message", "crud", "Rows")

	limit := constant1.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	conds := in.GetConds()
	if conds == nil {
		conds = &messagemgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetMessages", "ID", conds.GetID().GetValue(), "error", err)
			return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetMessages", "AppID", conds.GetAppID().GetValue(), "error", err)
			return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.MessageID != nil {
		if _, err := uuid.Parse(conds.GetMessageID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetMessages", "MessageID", conds.GetMessageID().GetValue(), "error", err)
			return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	infos, total, err := message1.GetMessages(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetMessages", "error", err)
		return &npool.GetMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMessagesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
