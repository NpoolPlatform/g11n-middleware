//nolint:nolintlint,dupl
package applang

import (
	"context"

	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/applang"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant1 "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/applang"

	"github.com/google/uuid"
)

func ValidateConds(ctx context.Context, conds *applangmgrpb.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetLangs", "ID", conds.GetID().GetValue(), "error", err)
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetLangs", "AppID", conds.GetAppID().GetValue(), "error", err)
			return err
		}
	}
	if conds.LangID != nil {
		if _, err := uuid.Parse(conds.GetLangID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetLangs", "LangID", conds.GetLangID().GetValue(), "error", err)
			return err
		}
	}

	return nil
}

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetLangs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "applang", "crud", "Rows")

	limit := constant1.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	conds := in.GetConds()
	if conds == nil {
		conds = &applangmgrpb.Conds{}
	}

	if err := ValidateConds(ctx, conds); err != nil {
		return &npool.GetLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := applang1.GetLangs(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetLangs", "error", err)
		return &npool.GetLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLangsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetLangOnly(ctx context.Context, in *npool.GetLangOnlyRequest) (*npool.GetLangOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetLangOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "applang", "crud", "Rows")

	conds := in.GetConds()
	if conds == nil {
		conds = &applangmgrpb.Conds{}
	}

	if err := ValidateConds(ctx, conds); err != nil {
		return &npool.GetLangOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := applang1.GetLangOnly(ctx, conds)
	if err != nil {
		logger.Sugar().Errorw("GetLangOnly", "error", err)
		return &npool.GetLangOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLangOnlyResponse{
		Info: info,
	}, nil
}
