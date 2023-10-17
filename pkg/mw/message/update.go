package message

import (
	"context"
	"fmt"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entmessage "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/message"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
	"github.com/google/uuid"
)

func (h *Handler) UpdateMessage(ctx context.Context) (*npool.Message, error) {
	info, err := h.GetMessage(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("message not exist")
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.MessageID != nil {
			appid := uuid.MustParse(info.AppID)
			h.AppID = &appid
			langid := uuid.MustParse(info.LangID)
			h.LangID = &langid
			h.Conds = &messagecrud.Conds{
				AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				LangID:    &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
				ID:        &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
				MessageID: &cruder.Cond{Op: cruder.EQ, Val: *h.MessageID},
			}
			exist, err := h.ExistMessageConds(ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("messageid exist")
			}
		}

		if _, err := tx.
			Message.
			Query().
			Where(
				entmessage.ID(*h.ID),
				entmessage.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx); err != nil {
			return err
		}

		if _, err := messagecrud.UpdateSet(
			tx.Message.UpdateOneID(*h.ID),
			&messagecrud.Req{
				MessageID: h.MessageID,
				Message:   h.Message,
				GetIndex:  h.GetIndex,
				Disabled:  h.Disabled,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetMessage(ctx)
}
