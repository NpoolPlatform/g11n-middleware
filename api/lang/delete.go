package lang

import (
	"context"

	lang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/lang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteLang(ctx context.Context, in *npool.DeleteLangRequest) (*npool.DeleteLangResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteLang",
			"In", in,
		)
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithID(req.ID, true),
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
