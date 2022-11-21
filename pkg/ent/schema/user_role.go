package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserRole struct {
	ent.Schema
}

func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("role_id"),
	}
}

func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required().Unique().Field("user_id"),
		edge.To("role", Role.Type).Required().Unique().Field("role_id"),
	}
}

func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
