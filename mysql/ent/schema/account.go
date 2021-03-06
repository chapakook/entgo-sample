package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("money").Default(0),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("account").Unique().Required(),
	}
}
