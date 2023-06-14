//nolint:nolintlint,dupl
package lang

import (
	"context"

	lang1 "github.com/NpoolPlatform/g11n-middleware/pkg/mw/lang"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	req := in.GetInfo()
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithID(req.ID),
		lang1.WithLang(req.Lang),
		lang1.WithLogo(req.Logo),
		lang1.WithName(req.Name),
		lang1.WithShort(req.Short),
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
