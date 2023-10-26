//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// Lang holds the schema definition for the Lang entity.
type Lang struct {
	ent.Schema
}

func (Lang) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Lang.
func (Lang) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("lang").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("short").
			Optional().
			Default(""),
	}
}

// Edges of the Lang.
func (Lang) Edges() []ent.Edge {
	return nil
}
