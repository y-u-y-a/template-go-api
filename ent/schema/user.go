package schema

import (
	"time"

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
		field.String("id").Immutable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).Immutable(),
		// -->
		field.String("given_name").Comment("名"),
		field.String("family_name").Comment("姓"),
		field.Int("gender").Comment("性別").Range(1, 3),
		field.String("email").Comment("メールアドレス").Unique(),
		field.Time("birthday").Comment("生年月日").SchemaType(map[string]string{dialect.Postgres: "date"}),
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
