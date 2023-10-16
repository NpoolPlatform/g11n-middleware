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

func (h *createHandler) validate(req *langcrud.Req) error {
	if req.Lang == nil || *req.Lang == "" {
		return fmt.Errorf("invalid lang")
	}
	if req.Logo == nil || *req.Logo == "" {
		return fmt.Errorf("invalid logo")
	}
	if req.Name == nil || *req.Name == "" {
		return fmt.Errorf("invalid name")
	}
	if req.Short == nil || *req.Short == "" {
		return fmt.Errorf("invalid short")
	}
	return nil
}

func (h *createHandler) createLang(ctx context.Context, cli *ent.Client, req *langcrud.Req) (*npool.Lang, error) {
	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixCreateLang,
		*req.Lang,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &langcrud.Conds{
		Lang: &cruder.Cond{Op: cruder.EQ, Val: *req.Lang},
	}
	h.Limit = 2
	infos, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}
	if infos != nil {
		return infos[0], nil
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := langcrud.CreateSet(
		cli.Lang.Create(),
		&langcrud.Req{
			EntID: req.EntID,
			Lang:  req.Lang,
			Logo:  req.Logo,
			Name:  req.Name,
			Short: req.Short,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil, nil
}

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		req := &langcrud.Req{
			EntID: h.EntID,
			Lang:  h.Lang,
			Logo:  h.Logo,
			Name:  h.Name,
			Short: h.Short,
		}
		if err := handler.validate(req); err != nil {
			return err
		}
		h.Conds = &langcrud.Conds{
			Lang: &cruder.Cond{Op: cruder.EQ, Val: *h.Lang},
		}
		exist, err := h.ExistLangConds(ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("lang exist")
		}
		info, err := handler.createLang(ctx, cli, req)
		if err != nil {
			return err
		}
		if info != nil {
			id, err := uuid.Parse(info.GetEntID())
			if err != nil {
				return err
			}
			h.EntID = &id
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
			if err := handler.validate(req); err != nil {
				return err
			}
			if req.EntID != nil {
				handler.EntID = req.EntID
			}
			info, err := handler.createLang(ctx, cli, req)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.GetEntID())
				if err != nil {
					return err
				}
				h.EntID = &id
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &langcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetLangs(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
