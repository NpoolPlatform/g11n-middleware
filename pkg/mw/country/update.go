package country

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entcountry "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/country"

	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

func (h *Handler) UpdateCountry(ctx context.Context) (*npool.Country, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Country.
			Query().
			Where(
				entcountry.ID(*h.ID),
				entcountry.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := countrycrud.UpdateSet(
			info.Update(),
			&countrycrud.Req{
				ID:      h.ID,
				Country: h.Country,
				Flag:    h.Flag,
				Code:    h.Code,
				Short:   h.Short,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCountry(ctx)
}
