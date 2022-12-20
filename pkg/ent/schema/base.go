package schema

import (
	"entgo.io/contrib/entproto"
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
			Annotations(entsql.Annotation{Incremental: &falsePtr}, entproto.Field(1)).
			DefaultFunc(func() int64 {
				return tools.GenSnowflakeID()
			}),
		//field.Int64("id").Immutable().DefaultFunc(tools.GenSnowflakeID()),
		field.Int64("created_by").Default(0).StructTag(`json:"created_by"`).Annotations(entgql.Type("String"), entproto.Field(2)),
		field.Int64("updated_by").Default(0).StructTag(`json:"updated_by"`).Annotations(entgql.Type("String"), entproto.Field(3)),
		field.Time("created_at").Immutable().Default(time.Now).StructTag(`json:"created_at"`).Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(4)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).StructTag(`json:"updated_at"`).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(5)),
		field.Time("deleted_at").Default(tools.ZeroTime).StructTag(`json:"deleted_at"`).Annotations(entgql.OrderField("DELETED_AT"), entproto.Field(6)),
	}
}
