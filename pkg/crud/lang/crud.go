//nolint:nolintlint,dupl
package lang

import (
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"
	entlang "github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/lang"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	Lang      *string
	Logo      *string
	Name      *string
	Short     *string
	DeletedAt *uint32
}

func CreateSet(c *ent.LangCreate, req *Req) *ent.LangCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Lang != nil {
		c.SetLang(*req.Lang)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Short != nil {
		c.SetShort(*req.Short)
	}
	return c
}

func UpdateSet(u *ent.LangUpdateOne, req *Req) *ent.LangUpdateOne {
	if req.Lang != nil {
		u.SetLang(*req.Lang)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
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
	ID    *cruder.Cond
	IDs   *cruder.Cond
	Lang  *cruder.Cond
	Langs *cruder.Cond
	Logo  *cruder.Cond
	Name  *cruder.Cond
	Short *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.LangQuery, conds *Conds) (*ent.LangQuery, error) {
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
			q.Where(entlang.ID(id))
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
			q.Where(entlang.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid ids field")
		}
	}
	if conds.Lang != nil {
		lang, ok := conds.Lang.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid lang")
		}
		switch conds.Lang.Op {
		case cruder.EQ:
			q.Where(entlang.Lang(lang))
		default:
			return nil, fmt.Errorf("invalid lang field")
		}
	}
	if conds.Langs != nil {
		langs, ok := conds.Langs.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid langs")
		}
		switch conds.Langs.Op {
		case cruder.IN:
			q.Where(entlang.LangIn(langs...))
		default:
			return nil, fmt.Errorf("invalid langs field")
		}
	}
	if conds.Logo != nil {
		logo, ok := conds.Logo.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid logo")
		}
		switch conds.Logo.Op {
		case cruder.EQ:
			q.Where(entlang.Logo(logo))
		default:
			return nil, fmt.Errorf("invalid logo field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entlang.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.Short != nil {
		short, ok := conds.Short.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid short")
		}
		switch conds.Short.Op {
		case cruder.EQ:
			q.Where(entlang.Short(short))
		default:
			return nil, fmt.Errorf("invalid short field")
		}
	}
	return q, nil
}
