package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// UserTopic holds the schema definition for the UserTopic entity.
type UserTopic struct {
	ent.Schema
}

// Fields of the UserTopic.
func (UserTopic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("user_id").Immutable(),
		field.Int("topic_id").Immutable(),
	}
}

// Edges of the UserTopic.
func (UserTopic) Edges() []ent.Edge {
	return nil
}
