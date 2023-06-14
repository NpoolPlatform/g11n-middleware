package lang

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	langcrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID     *uuid.UUID
	Lang   *string
	Name   *string
	Logo   *string
	Short  *string
	Reqs   []*langcrud.Req
	Conds  *langcrud.Conds
	Offset int32
	Limit  int32
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

func WithLang(lang *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lang == nil {
			return nil
		}
		if *lang == "" {
			return fmt.Errorf("invalid lang")
		}
		h.Lang = lang
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if logo == nil {
			return nil
		}
		if *logo == "" {
			return fmt.Errorf("invalid logo")
		}
		h.Logo = logo
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
		h.Conds = &langcrud.Conds{}
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
		if conds.Lang != nil {
			h.Conds.Lang = &cruder.Cond{Op: conds.GetLang().GetOp(), Val: conds.GetLang().GetValue()}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{Op: conds.GetName().GetOp(), Val: conds.GetName().GetValue()}
		}
		if conds.Short != nil {
			h.Conds.Short = &cruder.Cond{Op: conds.GetShort().GetOp(), Val: conds.GetShort().GetValue()}
		}
		if len(conds.GetLangs().GetValue()) > 0 {
			h.Conds.Langs = &cruder.Cond{Op: conds.GetLangs().GetOp(), Val: conds.GetLangs().GetValue()}
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

func WithReqs(reqs []*npool.LangReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*langcrud.Req{}
		for _, req := range reqs {
			_req := &langcrud.Req{}
			if req.ID != nil {
				id, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.Lang != nil {
				_req.Lang = req.Lang
			}
			if req.Name != nil {
				_req.Name = req.Name
			}
			if req.Logo != nil {
				_req.Logo = req.Logo
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
