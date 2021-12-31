package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Int("age").Positive(),
		field.String("name").Unique(),
		field.JSON("skills", []string{}).Optional(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("owner", User.Type). // default way
		// 	Ref("pets").
		// 	Unique(),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
