package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Text("content").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.Bool("is_deleted"),
		field.Int("creater_id").Immutable(),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return nil
}
