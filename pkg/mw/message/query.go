package message

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	entmessage "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

type queryHandler struct {
	*Handler
	stm   *ent.MessageSelect
	infos []*npool.Message
	total uint32
}

func (h *queryHandler) selectMessage(stm *ent.MessageQuery) {
	h.stm = stm.Select(
		entmessage.FieldID,
		entmessage.FieldAppID,
		entmessage.FieldLangID,
		entmessage.FieldMessageID,
		entmessage.FieldMessage,
		entmessage.FieldGetIndex,
		entmessage.FieldDisabled,
	)
}

func (h *queryHandler) queryMessage(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid messageid")
	}
	h.selectMessage(
		cli.Message.
			Query().
			Where(
				entmessage.ID(*h.ID),
				entmessage.DeletedAt(0),
			),
	)

	return nil
}

func (h *queryHandler) queryMessages(ctx context.Context, cli *ent.Client) error {
	stm, err := messagecrud.SetQueryConds(cli.Message.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectMessage(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetMessage(ctx context.Context) (*npool.Message, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryMessage(cli); err != nil {
			return err
		}
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetMessages(ctx context.Context) ([]*npool.Message, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryMessages(_ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}
