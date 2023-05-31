package country

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uuid.UUID
	Country *string
	Flag    *string
	Code    *string
	Short   *string
	Reqs    []*countrycrud.Req
	Conds   *countrycrud.Conds
	Offset  int32
	Limit   int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithCountry(country *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if country == nil {
			return nil
		}
		if *country == "" {
			return fmt.Errorf("invalid country")
		}
		h.Country = country
		return nil
	}
}

func WithFlag(flag *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if flag == nil {
			return nil
		}
		if *flag == "" {
			return fmt.Errorf("invalid flag")
		}
		h.Flag = flag
		return nil
	}
}

func WithCode(code *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if code == nil {
			return nil
		}
		if *code == "" {
			return fmt.Errorf("invalid code")
		}
		h.Code = code
		return nil
	}
}

func WithShort(short *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if short == nil {
			return nil
		}
		if *short == "" {
			return fmt.Errorf("invalid short")
		}
		h.Short = short
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &countrycrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.Country != nil {
			h.Conds.Country = &cruder.Cond{Op: conds.GetCountry().GetOp(), Val: conds.GetCountry().GetValue()}
		}
		if conds.Code != nil {
			h.Conds.Code = &cruder.Cond{Op: conds.GetCode().GetOp(), Val: conds.GetCode().GetValue()}
		}
		if conds.Short != nil {
			h.Conds.Short = &cruder.Cond{Op: conds.GetShort().GetOp(), Val: conds.GetShort().GetValue()}
		}
		if len(conds.GetCountries().GetValue()) > 0 {
			_ids := []uuid.UUID{}
			for _, id := range conds.GetCountries().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_ids = append(_ids, _id)
			}
			h.Conds.Countries = &cruder.Cond{Op: conds.GetCountries().GetOp(), Val: _ids}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}

func WithReqs(reqs []*npool.CountryReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*countrycrud.Req{}
		for _, req := range reqs {
			_req := &countrycrud.Req{}
			if req.ID != nil {
				id, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.Country != nil {
				_req.Country = req.Country
			}
			if req.Flag != nil {
				_req.Flag = req.Flag
			}
			if req.Code != nil {
				_req.Code = req.Code
			}
			if req.Short != nil {
				_req.Short = req.Short
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
