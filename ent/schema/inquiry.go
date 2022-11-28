package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Inquiry holds the schema definition for the Inquiry entity.
type Inquiry struct {
	ent.Schema
}

// Fields of the Inquiry.
func (Inquiry) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("名前").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}),
	}
}

// Edges of the Inquiry.
func (Inquiry) Edges() []ent.Edge {
	return nil
}
