package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("given_name").
			Comment("名").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}),
		field.String("family_name").
			Comment("姓").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}),
		field.String("email").
			Comment("メールアドレス").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}),
		field.Int("age").
			Comment("年齢").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}),
		field.Int("company_id").Comment("企業ID"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Unique().
			Required().
			Field("company_id"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("given_name", "family_name"),
	}
}
