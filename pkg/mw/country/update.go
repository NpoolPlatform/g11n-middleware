package country

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

func (h *Handler) UpdateCountry(ctx context.Context) (*npool.Country, error) {
	info, err := h.GetCountry(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("country not exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.Country != nil {
			h.Conds = &countrycrud.Conds{
				ID:      &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
				Country: &cruder.Cond{Op: cruder.EQ, Val: *h.Country},
			}
			exist, err := h.ExistCountryConds(ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("country is exist")
			}
		}

		if _, err := countrycrud.UpdateSet(
			tx.Country.UpdateOneID(*h.ID),
			&countrycrud.Req{
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
