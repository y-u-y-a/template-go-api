package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).Immutable(),
		// -->
		field.String("name").Comment("会社名").Immutable(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}
