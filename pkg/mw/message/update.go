package message

import (
	"context"
	"fmt"

	applangcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/applang"
	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entmessage "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/message"
	applangmw "github.com/NpoolPlatform/g11n-middleware/pkg/mw/applang"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
	"github.com/google/uuid"
)

func (h *Handler) UpdateMessage(ctx context.Context) (*npool.Message, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetMessage(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("message not exist")
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.MessageID != nil {
			if h.LangID == nil {
				return fmt.Errorf("invalid langid")
			}
			applanghandler, err := applangmw.NewHandler(
				ctx,
			)
			if err != nil {
				return err
			}
			applanghandler.Conds = &applangcrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				LangID: &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
			}
			applangexist, err := applanghandler.ExistAppLangConds(ctx)
			if err != nil {
				return err
			}
			if !applangexist {
				return fmt.Errorf("applang not exist")
			}

			id := uuid.MustParse(info.EntID)
			h.EntID = &id
			h.Conds = &messagecrud.Conds{
				AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				LangID:    &cruder.Cond{Op: cruder.EQ, Val: *h.LangID},
				EntID:     &cruder.Cond{Op: cruder.NEQ, Val: *h.EntID},
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
