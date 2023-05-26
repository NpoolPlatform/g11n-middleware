package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// AppCountry holds the schema definition for the AppCountry entity.
type AppCountry struct {
	ent.Schema
}

func (AppCountry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppCountry.
func (AppCountry) Fields() []ent.Field {
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
			UUID("country_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
	}
}

// Edges of the AppCountry.
func (AppCountry) Edges() []ent.Edge {
	return nil
}
