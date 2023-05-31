package country

import (
	"context"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	"github.com/google/uuid"
)

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := countrycrud.CreateSet(
			cli.Country.Create(),
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

func (h *Handler) CreateCountries(ctx context.Context) ([]*npool.Country, error) {
	ids := []uuid.UUID{}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = *req.ID
			}
			if _, err := countrycrud.CreateSet(
				cli.Country.Create(),
				&countrycrud.Req{
					ID:      &id,
					Country: req.Country,
					Flag:    req.Flag,
					Code:    req.Code,
					Short:   req.Short,
				},
			).Save(ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &countrycrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	infos, _, err := h.GetCountries(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
