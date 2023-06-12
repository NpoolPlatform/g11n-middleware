package message

import (
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entmessage "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/message"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	LangID    *uuid.UUID
	MessageID *string
	Message   *string
	GetIndex  *uint32
	Disabled  *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.MessageCreate, req *Req) *ent.MessageCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.LangID != nil {
		c.SetLangID(*req.LangID)
	}
	if req.MessageID != nil {
		c.SetMessageID(*req.MessageID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.GetIndex != nil {
		c.SetGetIndex(*req.GetIndex)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	if req.DeletedAt != nil {
		c.SetDeletedAt(*req.DeletedAt)
	}
	return c
}

func UpdateSet(u *ent.MessageUpdateOne, req *Req) *ent.MessageUpdateOne {
	if req.MessageID != nil {
		u.SetMessageID(*req.MessageID)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.GetIndex != nil {
		u.SetGetIndex(*req.GetIndex)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	AppID      *cruder.Cond
	LangID     *cruder.Cond
	MessageID  *cruder.Cond
	Disabled   *cruder.Cond
	MessageIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.MessageQuery, conds *Conds) (*ent.MessageQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.ID(id),
				entmessage.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entmessage.IDNEQ(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message id field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(
				entmessage.IDIn(ids...),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message ids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.AppID(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message appids field")
		}
	}
	if conds.LangID != nil {
		id, ok := conds.LangID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.LangID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.LangID(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message langid field")
		}
	}

	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(
				entmessage.Disabled(disabled),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message disabled field")
		}
	}
	if conds.MessageID != nil {
		id, ok := conds.MessageID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid messageid")
		}
		switch conds.MessageID.Op {
		case cruder.EQ:
			q.Where(
				entmessage.MessageID(id),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message messageid field")
		}
	}
	if conds.MessageIDs != nil {
		ids, ok := conds.MessageIDs.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid messageids")
		}
		switch conds.MessageIDs.Op {
		case cruder.IN:
			q.Where(
				entmessage.MessageIDIn(ids...),
				entmessage.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid message messageids field")
		}
	}
	return q, nil
}
