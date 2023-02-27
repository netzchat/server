package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return nil
}
