package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// TagTopic holds the schema definition for the TagTopic entity.
type TagTopic struct {
	ent.Schema
}

// Fields of the TagTopic.
func (TagTopic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("tag_id").Immutable(),
		field.Int("topic_id").Immutable(),
	}
}

// Edges of the TagTopic.
func (TagTopic) Edges() []ent.Edge {
	return nil
}
