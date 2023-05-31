//nolint:nolintlint,dupl
package appcountry

import (
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entappcountry "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/appcountry"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	CountryID *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppCountryCreate, req *Req) *ent.AppCountryCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CountryID != nil {
		c.SetCountryID(*req.CountryID)
	}
	return c
}

func UpdateSet(u *ent.AppCountryUpdateOne, req *Req) *ent.AppCountryUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	AppID      *cruder.Cond
	CountryID  *cruder.Cond
	AppIDs     *cruder.Cond
	CountryIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppCountryQuery, conds *Conds) (*ent.AppCountryQuery, error) {
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
			q.Where(entappcountry.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappcountry.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid ids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappcountry.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entappcountry.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appids field")
		}
	}
	if conds.CountryID != nil {
		id, ok := conds.CountryID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.CountryID.Op {
		case cruder.EQ:
			q.Where(entappcountry.CountryID(id))
		default:
			return nil, fmt.Errorf("invalid langid field")
		}
	}
	if conds.CountryIDs != nil {
		ids, ok := conds.CountryIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.CountryIDs.Op {
		case cruder.IN:
			q.Where(entappcountry.CountryIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid langids field")
		}
	}
	return q, nil
}
