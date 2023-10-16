package message

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate(req *messagecrud.Req) error {
	if req.LangID == nil {
		return fmt.Errorf("invalid langid")
	}
	if req.MessageID == nil || *req.MessageID == "" {
		return fmt.Errorf("invalid messageid")
	}
	if req.Message == nil || *req.Message == "" {
		return fmt.Errorf("invalid message")
	}
	return nil
}

func (h *createHandler) createMessage(ctx context.Context, cli *ent.Client, req *messagecrud.Req) (*npool.Message, error) {
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCountry,
		*req.AppID,
		*req.LangID,
		*req.MessageID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &messagecrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
		LangID:    &cruder.Cond{Op: cruder.EQ, Val: *req.LangID},
		MessageID: &cruder.Cond{Op: cruder.EQ, Val: *req.MessageID},
	}
	h.Limit = 2
	exist, _, err := h.GetMessages(ctx)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return exist[0], nil
	}

	id := uuid.New()
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := messagecrud.CreateSet(
		cli.Message.Create(),
		&messagecrud.Req{
			EntID:     req.EntID,
			AppID:     req.AppID,
			LangID:    req.LangID,
			MessageID: req.MessageID,
			Message:   req.Message,
			GetIndex:  req.GetIndex,
			Disabled:  req.Disabled,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil, nil
}

func (h *Handler) CreateMessage(ctx context.Context) (*npool.Message, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		req := &messagecrud.Req{
			EntID:     h.EntID,
			AppID:     h.AppID,
			LangID:    h.LangID,
			MessageID: h.MessageID,
			Message:   h.Message,
			GetIndex:  h.GetIndex,
			Disabled:  h.Disabled,
		}
		if err := handler.validate(req); err != nil {
			return err
		}
		h.Conds = &messagecrud.Conds{
			AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			LangID:    &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
			MessageID: &cruder.Cond{Op: cruder.EQ, Val: *h.MessageID},
		}
		exist, err := h.ExistMessageConds(ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("message exist")
		}
		info, err := handler.createMessage(ctx, cli, req)
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

	return h.GetMessage(ctx)
}

func (h *Handler) CreateMessages(ctx context.Context) ([]*npool.Message, error) {
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
			info, err := handler.createMessage(ctx, cli, req)
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

	h.Conds = &messagecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetMessages(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
