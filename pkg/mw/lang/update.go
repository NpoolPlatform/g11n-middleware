package lang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	"github.com/google/uuid"

	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

func (h *Handler) UpdateLang(ctx context.Context) (*npool.Lang, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetLang(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("lang not exist")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.Lang != nil {
			id := uuid.MustParse(info.EntID)
			h.EntID = &id
			h.Conds = &langcrud.Conds{
				EntID: &cruder.Cond{Op: cruder.NEQ, Val: *h.EntID},
				Lang:  &cruder.Cond{Op: cruder.EQ, Val: *h.Lang},
			}
			exist, err := h.ExistLangConds(ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("lang is exist")
			}
		}

		if _, err := langcrud.UpdateSet(
			tx.Lang.UpdateOneID(*h.ID),
			&langcrud.Req{
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
