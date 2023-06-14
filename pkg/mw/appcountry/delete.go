package appcountry

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	appcountrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/appcountry"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
)

func (h *Handler) DeleteCountry(ctx context.Context) (*npool.Country, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := h.GetCountry(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("appcountry not exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.Conds = &appcountrycrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			ID:    &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		}
		exist, err := h.ExistAppCountryConds(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("appcountry not exist")
		}
		now := uint32(time.Now().Unix())
		if _, err := appcountrycrud.UpdateSet(
			cli.AppCountry.UpdateOneID(*h.ID),
			&appcountrycrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
