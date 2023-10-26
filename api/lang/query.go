package lang

import (
	"context"

	lang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/lang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLang(ctx context.Context, in *npool.GetLangRequest) (*npool.GetLangResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLang",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLang",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLangResponse{
		Info: info,
	}, nil
}

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithConds(in.GetConds()),
		lang1.WithOffset(in.GetOffset()),
		lang1.WithLimit(in.GetLimit()),
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
