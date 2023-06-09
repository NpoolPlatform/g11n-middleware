package lang

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	entlang "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

type queryHandler struct {
	*Handler
	stm   *ent.LangSelect
	infos []*npool.Lang
	total uint32
}

func (h *queryHandler) selectLang(stm *ent.LangQuery) {
	h.stm = stm.Select(
		entlang.FieldID,
		entlang.FieldLang,
		entlang.FieldLogo,
		entlang.FieldName,
		entlang.FieldShort,
	)
}

func (h *queryHandler) queryLang(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid langid")
	}

	h.selectLang(
		cli.Lang.
			Query().
			Where(
				entlang.ID(*h.ID),
				entlang.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryLangs(ctx context.Context, cli *ent.Client) error {
	stm, err := langcrud.SetQueryConds(cli.Lang.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectLang(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetLang(ctx context.Context) (*npool.Lang, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLang(cli); err != nil {
			return err
		}
		const limit = 2
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(limit).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
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

func (h *Handler) GetLangs(ctx context.Context) ([]*npool.Lang, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLangs(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
