package message1

import (
	"context"
	"fmt"

	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"github.com/NpoolPlatform/g11n-manager/pkg/db"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"

	entlang "github.com/NpoolPlatform/g11n-manager/pkg/db/ent/lang"
	entmessage "github.com/NpoolPlatform/g11n-manager/pkg/db/ent/message"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

func GetMessage(ctx context.Context, id string) (*npool.Message, error) {
	infos := []*npool.Message{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Message.
			Query().
			Where(
				entmessage.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func GetMessages(ctx context.Context, conds *messagemgrpb.Conds, offset, limit int32) ([]*npool.Message, uint32, error) {
	infos := []*npool.Message{}
	total := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Message.
			Query()

		if conds.ID != nil {
			stm.
				Where(
					entmessage.ID(uuid.MustParse(conds.GetID().GetValue())),
				)
		}
		if conds.AppID != nil {
			stm.
				Where(
					entmessage.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
		}
		if conds.LangID != nil {
			stm.
				Where(
					entmessage.LangID(uuid.MustParse(conds.GetLangID().GetValue())),
				)
		}
		if conds.MessageID != nil {
			stm.
				Where(
					entmessage.MessageID(conds.GetMessageID().GetValue()),
				)
		}
		if conds.Disabled != nil {
			stm.
				Where(
					entmessage.Disabled(conds.GetDisabled().GetValue()),
				)
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, fmt.Errorf("no record")
	}
	if len(infos) > 1 {
		return nil, 0, fmt.Errorf("too many record")
	}

	return infos, total, nil
}

func GetManyMessages(ctx context.Context, ids []string) ([]*npool.Message, error) {
	infos := []*npool.Message{}

	messageIDs := []uuid.UUID{}
	for _, id := range ids {
		messageIDs = append(messageIDs, uuid.MustParse(id))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Message.
			Query().
			Where(
				entmessage.IDIn(messageIDs...),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func join(stm *ent.MessageQuery) *ent.MessageSelect {
	return stm.
		Select(
			entmessage.FieldID,
			entmessage.FieldAppID,
			entmessage.FieldLangID,
			entmessage.FieldMessageID,
			entmessage.FieldMessage,
			entmessage.FieldGetIndex,
			entmessage.FieldDisabled,
			entmessage.FieldCreatedAt,
			entmessage.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entlang.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entmessage.FieldLangID),
					t1.C(entlang.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entlang.FieldLang), "lang"),
				)
		})
}
