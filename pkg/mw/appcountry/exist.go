package appcountry

import (
	"context"

	appcountrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/appcountry"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
)

func (h *Handler) ExistAppCountryConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appcountrycrud.SetQueryConds(cli.AppCountry.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
