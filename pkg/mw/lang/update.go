package lang

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entlang "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/lang"

	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

func (h *Handler) UpdateLang(ctx context.Context) (*npool.Lang, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Lang.
			Query().
			Where(
				entlang.ID(*h.ID),
				entlang.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := langcrud.UpdateSet(
			info.Update(),
			&langcrud.Req{
				ID:    h.ID,
				Lang:  h.Lang,
				Logo:  h.Logo,
				Name:  h.Name,
				Short: h.Short,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLang(ctx)
}
