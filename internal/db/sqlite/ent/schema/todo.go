package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

// This schema is used to create test data.
type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.Int("listId").Default(0),
		field.String("text").Default(""),
		field.Bool("completed").Default(false),
	}
}

func (Todo) Annotations() []schema.Annotation {
	return nil
}
