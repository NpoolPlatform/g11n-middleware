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
	EntID     *uuid.UUID
	Lang      *string
	Logo      *string
	Name      *string
	Short     *string
	DeletedAt *uint32
}

func CreateSet(c *ent.LangCreate, req *Req) *ent.LangCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	EntID  *cruder.Cond
	EntIDs *cruder.Cond
	Lang   *cruder.Cond
	Langs  *cruder.Cond
	Logo   *cruder.Cond
	Name   *cruder.Cond
	Short  *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.LangQuery, conds *Conds) (*ent.LangQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entlang.EntID(id),
				entlang.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entlang.EntIDNEQ(id),
				entlang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(
				entlang.EntIDIn(ids...),
				entlang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.Lang != nil {
		lang, ok := conds.Lang.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid lang")
		}
		switch conds.Lang.Op {
		case cruder.EQ:
			q.Where(
				entlang.Lang(lang),
				entlang.DeletedAt(0),
			)
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
			q.Where(
				entlang.LangIn(langs...),
				entlang.DeletedAt(0),
			)
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
			q.Where(
				entlang.Logo(logo),
				entlang.DeletedAt(0),
			)
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
			q.Where(
				entlang.Name(name),
				entlang.DeletedAt(0),
			)
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
			q.Where(
				entlang.Short(short),
				entlang.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid short field")
		}
	}
	return q, nil
}
