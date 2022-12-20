package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserRole struct {
	ent.Schema
}

func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Annotations(entproto.Field(11)),
		field.Int64("role_id").Annotations(entproto.Field(12)),
	}
}

func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required().Unique().Field("user_id").Annotations(entproto.Skip()),
		edge.To("role", Role.Type).Required().Unique().Field("role_id").Annotations(entproto.Skip()),
	}
}

func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
