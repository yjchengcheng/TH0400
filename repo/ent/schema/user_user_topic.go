package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUserTopic holds the schema definition for the UserUserTopic entity.
type UserUserTopic struct {
	ent.Schema
}

// Fields of the UserUserTopic.
func (UserUserTopic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("inviter_id").Immutable(),
		field.Int("invited_id").Immutable(),
		field.Int("topic_id").Immutable(),
	}
}

// Edges of the UserUserTopic.
func (UserUserTopic) Edges() []ent.Edge {
	return nil
}
