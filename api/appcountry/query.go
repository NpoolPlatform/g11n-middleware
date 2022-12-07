//nolint:nolintlint,dupl
package appcountry

import (
	"context"

	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/appcountry"
	commontracer "github.com/NpoolPlatform/g11n-middleware/pkg/tracer"

	constant1 "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-middleware/pkg/appcountry"

	"github.com/google/uuid"
)

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCountries")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "appcountry", "crud", "Rows")

	limit := constant1.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	conds := in.GetConds()
	if conds == nil {
		conds = &appcountrymgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCountries", "ID", conds.GetID().GetValue(), "error", err)
			return &npool.GetCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCountries", "AppID", conds.GetAppID().GetValue(), "error", err)
			return &npool.GetCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CountryID != nil {
		if _, err := uuid.Parse(conds.GetCountryID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCountries", "CountryID", conds.GetCountryID().GetValue(), "error", err)
			return &npool.GetCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	infos, total, err := appcountry1.GetCountries(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCountries", "error", err)
		return &npool.GetCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
