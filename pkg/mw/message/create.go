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

func (h *createHandler) createMessage(ctx context.Context, cli *ent.Client) error {
	if h.LangID == nil {
		return fmt.Errorf("langid invalid")
	}
	if h.MessageID == nil {
		return fmt.Errorf("messageid invalid")
	}
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCountry,
		&h.AppID,
		*h.LangID,
		*h.MessageID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &messagecrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
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

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	info, err := messagecrud.CreateSet(
		cli.Message.Create(),
		&messagecrud.Req{
			ID:        h.ID,
			AppID:     &h.AppID,
			LangID:    h.LangID,
			MessageID: h.MessageID,
			Message:   h.Message,
			GetIndex:  h.GetIndex,
			Disabled:  h.Disabled,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *Handler) CreateMessage(ctx context.Context) (*npool.Message, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createMessage(ctx, cli); err != nil {
			return err
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
			handler.ID = nil
			handler.AppID = *req.AppID
			handler.LangID = req.LangID
			handler.MessageID = req.MessageID
			handler.Message = req.Message
			handler.GetIndex = req.GetIndex
			handler.Disabled = req.Disabled
			if err := handler.createMessage(ctx, cli); err != nil {
				return err
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &messagecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetMessages(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
