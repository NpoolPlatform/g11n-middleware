package applang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	applangcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/applang"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"github.com/google/uuid"
)

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}
	if h.LangID == nil {
		return nil, fmt.Errorf("invalid langid")
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := applangcrud.SetQueryConds(
			cli.AppLang.Query(),
			&applangcrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
				LangID: &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			h.ID = &info.ID
			return nil
		}
		if _, err := applangcrud.CreateSet(
			cli.AppLang.Create(),
			&applangcrud.Req{
				ID:     h.ID,
				AppID:  &h.AppID,
				LangID: h.LangID,
				Main:   h.Main,
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
			if _, err := applangcrud.CreateSet(
				cli.AppLang.Create(),
				&applangcrud.Req{
					ID:     &id,
					AppID:  req.AppID,
					LangID: req.LangID,
					Main:   req.Main,
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

	h.Conds = &applangcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	infos, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
