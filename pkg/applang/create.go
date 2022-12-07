package applang

import (
	"context"

	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	applangmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/applang"
)

func CreateLang(ctx context.Context, in *applangmgrpb.LangReq) (*npool.Lang, error) {
	info, err := applangmgrcli.CreateLang(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetLang(ctx, info.ID)
}

func CreateLangs(ctx context.Context, in []*applangmgrpb.LangReq) ([]*npool.Lang, error) {
	infos, err := applangmgrcli.CreateLangs(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, info := range infos {
		ids = append(ids, info.ID)
	}

	return GetManyLangs(ctx, ids)
}
