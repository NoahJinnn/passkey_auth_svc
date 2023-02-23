// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type BankAccount struct {
	ent.Schema
}

func (BankAccount) Fields() []ent.Field {
	return []ent.Field{field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}), field.Uint("user_id").Optional().SchemaType(map[string]string{"postgres": "serial"}), field.Uint("asset_info_id").SchemaType(map[string]string{"postgres": "serial"}), field.Time("created_at"), field.Time("updated_at")}

}
func (BankAccount) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("bank_accounts").Unique().Field("user_id")}
}
func (BankAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "bank_account"}}
}
