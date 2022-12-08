package schema

import (
	"entgo.io/ent/dialect/entsql"
	"github.com/stark-sim/cas/tools"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	falsePtr := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			Annotations(entsql.Annotation{Incremental: &falsePtr}).
			DefaultFunc(func() int64 {
				return tools.GenSnowflakeID()
			}),
		//field.Int64("id").Immutable().DefaultFunc(tools.GenSnowflakeID()),
		field.Int64("created_by").Default(0).StructTag(`json:"created_by"`).Annotations(entgql.Type("String")),
		field.Int64("updated_by").Default(0).StructTag(`json:"updated_by"`).Annotations(entgql.Type("String")),
		field.Time("created_at").Immutable().Default(time.Now).StructTag(`json:"created_at"`).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).StructTag(`json:"updated_at"`).Annotations(entgql.OrderField("UPDATED_AT")),
		field.Time("deleted_at").Default(tools.ZeroTime).StructTag(`json:"deleted_at"`).Annotations(entgql.OrderField("DELETED_AT")),
	}
}
