package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// AppLang holds the schema definition for the AppLang entity.
type AppLang struct {
	ent.Schema
}

func (AppLang) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppLang.
func (AppLang) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("lang_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			Bool("main").
			Optional().
			Default(false),
	}
}

// Edges of the AppLang.
func (AppLang) Edges() []ent.Edge {
	return nil
}
