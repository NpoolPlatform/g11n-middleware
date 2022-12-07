package applang

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	applangmgrcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
)

func DeleteLang(ctx context.Context, id string) (*npool.Lang, error) {
	info, err := GetLang(ctx, id)
	if err != nil {
		return nil, err
	}

	_, err = applangmgrcli.DeleteLang(ctx, id)
	if err != nil {
		return nil, err
	}

	return info, nil
}
