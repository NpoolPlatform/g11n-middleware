package message

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (h *Handler) DeleteMessage(ctx context.Context) (*npool.Message, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := h.GetMessage(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		h.Conds = &messagecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			ID:    &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		}
		exist, err := h.ExistMessageConds(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("message not exist")
		}
		now := uint32(time.Now().Unix())
		if _, err := messagecrud.UpdateSet(
			cli.Message.UpdateOneID(*h.ID),
			&messagecrud.Req{
				ID:        h.ID,
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
