package applang

import (
	"context"

	applang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/applang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteLang(ctx context.Context, in *npool.DeleteLangRequest) (*npool.DeleteLangResponse, error) {
	req := in.GetInfo()
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithID(req.ID),
		applang1.WithAppID(req.AppID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLang",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLang",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLangResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteLangResponse{
		Info: info,
	}, nil
}
