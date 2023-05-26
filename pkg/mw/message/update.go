package message

import (
	"context"
	"fmt"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entmessage "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (h *Handler) UpdateMessage(ctx context.Context) (*npool.Message, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Message.
			Query().
			Where(
				entmessage.ID(*h.ID),
				entmessage.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := messagecrud.UpdateSet(
			info.Update(),
			&messagecrud.Req{
				ID:        h.ID,
				MessageID: h.MessageID,
				Message:   h.Message,
				GetIndex:  &h.GetIndex,
				Disabled:  &h.Disabled,
				Short:     h.Short,
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
