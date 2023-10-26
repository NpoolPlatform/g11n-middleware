package lang

import (
	"context"
	"time"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

func (h *Handler) DeleteLang(ctx context.Context) (*npool.Lang, error) {
	info, err := h.GetLang(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := langcrud.UpdateSet(
			cli.Lang.UpdateOneID(*h.ID),
			&langcrud.Req{
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
