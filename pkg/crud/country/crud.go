//nolint:dupl
package country

import (
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entcountry "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/country"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	Country   *string
	Flag      *string
	Code      *string
	Short     *string
	DeletedAt *uint32
}

func CreateSet(c *ent.CountryCreate, req *Req) *ent.CountryCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Country != nil {
		c.SetCountry(*req.Country)
	}
	if req.Flag != nil {
		c.SetFlag(*req.Flag)
	}
	if req.Code != nil {
		c.SetCode(*req.Code)
	}
	if req.Short != nil {
		c.SetShort(*req.Short)
	}
	return c
}

func UpdateSet(u *ent.CountryUpdateOne, req *Req) *ent.CountryUpdateOne {
	if req.Country != nil {
		u.SetCountry(*req.Country)
	}
	if req.Flag != nil {
		u.SetFlag(*req.Flag)
	}
	if req.Code != nil {
		u.SetCode(*req.Code)
	}
	if req.Short != nil {
		u.SetShort(*req.Short)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	Country   *cruder.Cond
	Code      *cruder.Cond
	Short     *cruder.Cond
	Countries *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CountryQuery, conds *Conds) (*ent.CountryQuery, error) {
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
				entcountry.ID(id),
				entcountry.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entcountry.IDNEQ(id),
				entcountry.DeletedAt(0),
			)
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
			q.Where(
				entcountry.IDIn(ids...),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid ids field")
		}
	}
	if conds.Country != nil {
		country, ok := conds.Country.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid country")
		}
		switch conds.Country.Op {
		case cruder.EQ:
			q.Where(
				entcountry.Country(country),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid country field")
		}
	}
	if conds.Code != nil {
		code, ok := conds.Code.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid code")
		}
		switch conds.Code.Op {
		case cruder.EQ:
			q.Where(
				entcountry.Code(code),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid code field")
		}
	}
	if conds.Short != nil {
		short, ok := conds.Short.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid short")
		}
		switch conds.Short.Op {
		case cruder.EQ:
			q.Where(
				entcountry.Short(short),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid short field")
		}
	}
	if conds.Countries != nil {
		countries, ok := conds.Countries.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid countries")
		}
		switch conds.Countries.Op {
		case cruder.IN:
			q.Where(
				entcountry.CountryIn(countries...),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid countries field")
		}
	}
	return q, nil
}
