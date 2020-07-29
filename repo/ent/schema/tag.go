package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("tag_name").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Bool("baned").Default(false),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
