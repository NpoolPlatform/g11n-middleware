package lang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createLang(ctx context.Context, cli *ent.Client) (*npool.Lang, error) {
	if h.Lang == nil || *h.Lang == "" {
		return nil, fmt.Errorf("invalid lang")
	}
	if h.Logo == nil || *h.Logo == "" {
		return nil, fmt.Errorf("invalid logo")
	}
	if h.Name == nil || *h.Name == "" {
		return nil, fmt.Errorf("invalid name")
	}
	if h.Short == nil || *h.Short == "" {
		return nil, fmt.Errorf("invalid short")
	}
	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixCreateLang,
		*h.Lang,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &langcrud.Conds{
		Lang: &cruder.Cond{Op: cruder.EQ, Val: *h.Lang},
	}
	h.Limit = 2
	exist, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return exist[0], nil
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	info, err := langcrud.CreateSet(
		cli.Lang.Create(),
		&langcrud.Req{
			ID:    h.ID,
			Lang:  h.Lang,
			Logo:  h.Logo,
			Name:  h.Name,
			Short: h.Short,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID

	return nil, nil
}

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := handler.createLang(ctx, cli)
		if err != nil {
			return err
		}
		if info != nil {
			id, err := uuid.Parse(info.GetID())
			if err != nil {
				return err
			}
			h.ID = &id
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLang(ctx)
}

func (h *Handler) CreateLangs(ctx context.Context) ([]*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.Lang = req.Lang
			handler.Logo = req.Logo
			handler.Name = req.Name
			handler.Short = req.Short
			info, err := handler.createLang(ctx, cli)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.GetID())
				if err != nil {
					return err
				}
				h.ID = &id
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &langcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
