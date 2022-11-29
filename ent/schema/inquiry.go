package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Inquiry holds the schema definition for the Inquiry entity.
type Inquiry struct {
	ent.Schema
}

// Fields of the Inquiry.
func (Inquiry) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).Immutable(),
		// -->
		field.String("name").Comment("名前"),
		field.String("email").Comment("メールアドレス"),
		field.String("tel").Comment("電話番号"),
		field.String("content").Comment("お問い合わせ内容"),
		field.Bool("is_confirm").Comment("対応フラグ").Default(false),
	}
}

// Edges of the Inquiry.
func (Inquiry) Edges() []ent.Edge {
	return nil
}
