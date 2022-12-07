package applang

import (
	"context"

	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	applangmgrcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
)

func UpdateLang(ctx context.Context, in *applangmgrpb.LangReq) (*npool.Lang, error) {
	info, err := applangmgrcli.UpdateLang(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetLang(ctx, info.ID)
}
