//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("country").
			Optional().
			Default(""),
		field.
			String("flag").
			Optional().
			Default(""),
		field.
			String("code").
			Optional().
			Default(""),
		field.
			String("short").
			Optional().
			Default(""),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return nil
}
