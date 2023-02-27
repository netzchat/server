package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/netzchat/server/ent/mixin"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

func (Channel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ObjectIDMixin{},
	}
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
