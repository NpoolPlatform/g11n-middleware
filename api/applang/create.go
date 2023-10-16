//nolint:nolintlint,dupl
package applang

import (
	"context"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/applang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	req := in.GetInfo()
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithEntID(req.EntID, false),
		applang1.WithAppID(req.AppID, true),
		applang1.WithLangID(req.LangID, true),
		applang1.WithMain(req.Main, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLang",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLang",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateLangs(ctx context.Context, in *npool.CreateLangsRequest) (*npool.CreateLangsResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLangs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateLangs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLangs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateLangsResponse{
		Infos: infos,
	}, nil
}
