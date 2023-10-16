package applang

import (
	"context"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/applang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	req := in.GetInfo()
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithID(req.ID, true),
		applang1.WithAppID(req.AppID, true),
		applang1.WithMain(req.Main, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLang",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLang",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateLangResponse{
		Info: info,
	}, nil
}
