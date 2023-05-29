package message

import (
	"context"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"github.com/google/uuid"
)

func (h *Handler) CreateMessage(ctx context.Context) (*npool.Message, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := messagecrud.CreateSet(
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

func (h *Handler) CreateMessages(ctx context.Context) ([]*npool.Message, error) {
	ids := []uuid.UUID{}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = uuid.MustParse(*req.ID)
			}
			appID := uuid.MustParse(*req.AppID)
			langID := uuid.MustParse(*req.LangID)
			if _, err := messagecrud.CreateSet(
				cli.Message.Create(),
				&messagecrud.Req{
					ID:        &id,
					AppID:     &appID,
					LangID:    &langID,
					MessageID: req.Message,
					Message:   req.Message,
					GetIndex:  req.GetIndex,
					Disabled:  req.Disabled,
				},
			).Save(ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &messagecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	infos, _, err := h.GetMessages(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
