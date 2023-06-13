package applang

import (
	"context"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/applang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLangOnly(ctx context.Context, in *npool.GetLangOnlyRequest) (*npool.GetLangOnlyResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLangOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetLangOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLangOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLangOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithConds(in.GetConds()),
		applang1.WithOffset(in.GetOffset()),
		applang1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLangs",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetLangs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLangs",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLangsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
