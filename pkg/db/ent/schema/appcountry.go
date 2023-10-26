package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// AppCountry holds the schema definition for the AppCountry entity.
type AppCountry struct {
	ent.Schema
}

func (AppCountry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppCountry.
func (AppCountry) Fields() []ent.Field {
	return []ent.Field{
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
