package message

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/const"
	messagecrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	MessageID *string
	Message   *string
	GetIndex  *uint32
	Disabled  *bool
	Reqs      []*messagecrud.Req
	Conds     *messagecrud.Conds
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

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithLangID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.LangID = &_id
		return nil
	}
}

func WithMessageID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		if *id == "" {
			return fmt.Errorf("invalid messageid")
		}
		h.MessageID = id
		return nil
	}
}

func WithMessage(message *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Message = message
		return nil
	}
}

func WithGetIndex(getindex *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GetIndex = getindex
		return nil
	}
}

func WithDisabled(disabled *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Disabled = disabled
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &messagecrud.Conds{}
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
		if conds.LangID != nil {
			id, err := uuid.Parse(conds.GetLangID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.LangID = &cruder.Cond{Op: conds.GetLangID().GetOp(), Val: id}
		}
		if conds.MessageID != nil {
			h.Conds.MessageID = &cruder.Cond{Op: conds.GetMessageID().GetOp(), Val: conds.GetMessageID().GetValue()}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{Op: conds.GetDisabled().GetOp(), Val: conds.GetDisabled().GetValue()}
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

func WithReqs(reqs []*npool.MessageReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*messagecrud.Req{}
		for _, req := range reqs {
			_req := &messagecrud.Req{}
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
			if req.LangID != nil {
				id, err := uuid.Parse(*req.LangID)
				if err != nil {
					return err
				}
				_req.LangID = &id
			}
			if req.MessageID != nil {
				_req.MessageID = req.MessageID
			}
			if req.Message != nil {
				_req.Message = req.Message
			}
			if req.GetIndex != nil {
				_req.GetIndex = req.GetIndex
			}
			if req.Disabled != nil {
				_req.Disabled = req.Disabled
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
