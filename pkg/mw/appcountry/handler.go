package appcountry

import (
	"context"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	appcountrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     uuid.UUID
	CountryID *uuid.UUID
	Reqs      []*appcountrycrud.Req
	Conds     *appcountrycrud.Conds
	Offset    int32
	Limit     int32
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

func WithAppID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppID = _id
		return nil
	}
}

func WithCountryID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CountryID = &_id
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appcountrycrud.Conds{}
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.CountryID != nil {
			id, err := uuid.Parse(conds.GetCountryID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CountryID = &cruder.Cond{Op: conds.GetCountryID().GetOp(), Val: id}
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
		_reqs := []*appcountrycrud.Req{}
		for _, req := range reqs {
			_req := &appcountrycrud.Req{}
			if req.ID != nil {
				id, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.CountryID != nil {
				id, err := uuid.Parse(*req.CountryID)
				if err != nil {
					return err
				}
				_req.CountryID = &id
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
