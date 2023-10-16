package message

import (
	"context"
	"time"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (h *Handler) DeleteMessage(ctx context.Context) (*npool.Message, error) {
	info, err := h.GetMessage(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := messagecrud.UpdateSet(
			cli.Message.UpdateOneID(*h.ID),
			&messagecrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
