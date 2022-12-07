package applang

import (
	"context"

	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
)

func GetLangs(ctx context.Context, in *applangmgrpb.Conds, offset, limit int32) ([]*npool.Lang, uint32, error) {
	return nil, 0, nil
}
