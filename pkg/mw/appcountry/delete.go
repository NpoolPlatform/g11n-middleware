package appcountry

import (
	"context"
	"time"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	appcountrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
)

func (h *Handler) DeleteCountry(ctx context.Context) (*npool.Country, error) {
	info, err := h.GetCountry(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := appcountrycrud.UpdateSet(
			cli.AppCountry.UpdateOneID(*h.ID),
			&appcountrycrud.Req{
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
