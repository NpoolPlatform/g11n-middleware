package applang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	applangcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/applang"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createLang(ctx context.Context, cli *ent.Client, req *applangcrud.Req) error {
	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppLang,
		*req.AppID,
		*req.LangID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &applangcrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID: &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
	}
	exist, err := h.ExistAppLangConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("applang exist")
	}
	if h.Main != nil {
		if *h.Main {
			h.Conds = &applangcrud.Conds{
				AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
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

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	info, err := applangcrud.CreateSet(
		cli.AppLang.Create(),
		&applangcrud.Req{
			EntID:  h.EntID,
			AppID:  h.AppID,
			LangID: h.LangID,
			Main:   h.Main,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil
}

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		req := &applangcrud.Req{
			EntID:  h.EntID,
			AppID:  h.AppID,
			LangID: h.LangID,
			Main:   h.Main,
		}
		if err := handler.createLang(ctx, cli, req); err != nil {
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
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			if req.EntID != nil {
				handler.EntID = req.EntID
			}
			if err := handler.createLang(ctx, cli, req); err != nil {
				return err
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &applangcrud.Conds{
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
