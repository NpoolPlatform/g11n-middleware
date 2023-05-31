package country

import (
	"context"
	"time"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

func (h *Handler) DeleteCountry(ctx context.Context) (*npool.Country, error) {
	info, err := h.GetCountry(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := countrycrud.UpdateSet(
			cli.Country.UpdateOneID(*h.ID),
			&countrycrud.Req{
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
