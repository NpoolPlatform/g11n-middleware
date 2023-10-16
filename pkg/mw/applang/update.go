package applang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	applangcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/applang"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
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
		return nil, fmt.Errorf("applang not exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if h.Main != nil {
			if *h.Main {
				h.Conds = &applangcrud.Conds{
					EntID: &cruder.Cond{Op: cruder.NEQ, Val: *h.EntID},
					Main:  &cruder.Cond{Op: cruder.EQ, Val: true},
				}
				exist, err := h.ExistAppLangConds(ctx)
				if err != nil {
					return err
				}
				if exist {
					return fmt.Errorf("applang main exist")
				}
			}
		}
		if _, err := applangcrud.UpdateSet(
			cli.AppLang.UpdateOneID(*h.ID),
			&applangcrud.Req{
				Main: h.Main,
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
