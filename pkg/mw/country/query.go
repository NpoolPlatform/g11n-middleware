package country

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	entcountry "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

type queryHandler struct {
	*Handler
	stm   *ent.CountrySelect
	infos []*npool.Country
	total uint32
}

func (h *queryHandler) selectCountry(stm *ent.CountryQuery) {
	h.stm = stm.Select(
		entcountry.FieldID,
		entcountry.FieldCountry,
		entcountry.FieldFlag,
		entcountry.FieldCode,
		entcountry.FieldShort,
	)
}

func (h *queryHandler) queryCountry(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}

	h.selectCountry(
		cli.Country.
			Query().
			Where(
				entcountry.ID(*h.ID),
				entcountry.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryCountries(ctx context.Context, cli *ent.Client) error {
	stm, err := countrycrud.SetQueryConds(cli.Country.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectCountry(stm)
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
		if err := handler.queryCountry(cli); err != nil {
			return err
		}
		const limit = 2
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(limit).
			Modify(func(s *sql.Selector) {})
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

func (h *Handler) GetCountries(ctx context.Context) ([]*npool.Country, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCountries(ctx, cli); err != nil {
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

func (h *Handler) GetCountryOnly(ctx context.Context) (*npool.Country, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCountries(_ctx, cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}
