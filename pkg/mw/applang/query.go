package applang

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	entapplang "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/applang"
	entlang "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	applangcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/applang"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppLangSelect
	infos []*npool.Lang
	total uint32
}

func (h *queryHandler) selectAppLang(stm *ent.AppLangQuery) {
	h.stm = stm.Select(
		entapplang.FieldID,
		entapplang.FieldAppID,
		entapplang.FieldLangID,
		entapplang.FieldMain,
		entapplang.FieldCreatedAt,
		entapplang.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAppLang(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid applangid")
	}

	h.selectAppLang(
		cli.AppLang.
			Query().
			Where(
				entapplang.ID(*h.ID),
				entapplang.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryAppLangs(ctx context.Context, cli *ent.Client) error {
	stm, err := applangcrud.SetQueryConds(cli.AppLang.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAppLang(stm)
	return nil
}

func (h *queryHandler) queryJoinLang(s *sql.Selector) {
	t := sql.Table(entlang.Table)
	s.LeftJoin(t).
		On(
			s.C(entapplang.FieldLangID),
			t.C(entlang.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entlang.FieldLang), "lang"),
			sql.As(t.C(entlang.FieldLogo), "logo"),
			sql.As(t.C(entlang.FieldName), "name"),
			sql.As(t.C(entlang.FieldShort), "short"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinLang(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetLang(ctx context.Context) (*npool.Lang, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppLang(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const limit = 2
		handler.stm = handler.stm.
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
		if err := handler.queryAppLangs(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm = handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
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
