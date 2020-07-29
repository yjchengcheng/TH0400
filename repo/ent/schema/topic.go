package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title").NotEmpty(),
		field.Text("content").Default("caonima"),
		field.Bool("is_released"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.Int("creater_id").Immutable(),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return nil
}
