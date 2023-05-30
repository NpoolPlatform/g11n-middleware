package lang

import (
	"context"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	"github.com/google/uuid"
)

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := langcrud.CreateSet(
			cli.Lang.Create(),
			&langcrud.Req{
				ID:    h.ID,
				Lang:  h.Lang,
				Logo:  h.Logo,
				Name:  h.Name,
				Short: h.Short,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLang(ctx)
}

func (h *Handler) CreateLangs(ctx context.Context) ([]*npool.Lang, error) {
	ids := []uuid.UUID{}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = *req.ID
			}
			if _, err := langcrud.CreateSet(
				cli.Lang.Create(),
				&langcrud.Req{
					ID:    &id,
					Lang:  req.Lang,
					Logo:  req.Logo,
					Name:  req.Name,
					Short: req.Short,
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

	h.Conds = &langcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	infos, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
