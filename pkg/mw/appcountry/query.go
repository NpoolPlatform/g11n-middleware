package appcountry

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	entappcountry "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/appcountry"
	entcountry "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	appcountrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/appcountry"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppCountrySelect
	infos []*npool.Country
	total uint32
}

func (h *queryHandler) selectAppCountry(stm *ent.AppCountryQuery) {
	h.stm = stm.Select(
		entappcountry.FieldID,
	)
}

func (h *queryHandler) queryAppCountry(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid appcountryid")
	}

	h.selectAppCountry(
		cli.AppCountry.
			Query().
			Where(
				entappcountry.ID(*h.ID),
				entappcountry.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryAppCountries(ctx context.Context, cli *ent.Client) error {
	stm, err := appcountrycrud.SetQueryConds(cli.AppCountry.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAppCountry(stm)
	return nil
}

func (h *queryHandler) queryJoinCountry(s *sql.Selector) {
	t := sql.Table(entcountry.Table)
	stm := s.LeftJoin(t).
		On(
			s.C(entappcountry.FieldCountryID),
			t.C(entcountry.FieldID),
		)

	stm.AppendSelect(
		s.C(entappcountry.FieldAppID),
		s.C(entappcountry.FieldCountryID),
		s.C(entappcountry.FieldCreatedAt),
		s.C(entappcountry.FieldUpdatedAt),
		sql.As(t.C(entcountry.FieldCountry), "country"),
		sql.As(t.C(entcountry.FieldFlag), "flag"),
		sql.As(t.C(entcountry.FieldCode), "code"),
		sql.As(t.C(entcountry.FieldShort), "short"),
	)
}

func (h *queryHandler) queryJoinSelect() {
	h.stm.Select(
		entappcountry.FieldID,
	)
}

func (h *queryHandler) queryJoin(ctx context.Context) error {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinSelect()
		h.queryJoinCountry(s)
	})
	total, err := h.stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetCountry(ctx context.Context) (*npool.Country, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCountry(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(ctx); err != nil {
			return err
		}
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

func (h *Handler) GetCountries(ctx context.Context) ([]*npool.Country, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppCountries(ctx, cli); err != nil {
			return err
		}
		if err := handler.queryJoin(ctx); err != nil {
			return err
		}
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
