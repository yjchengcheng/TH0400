package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUser holds the schema definition for the UserUser entity.
type UserUser struct {
	ent.Schema
}

// Fields of the UserUser.
func (UserUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("idol_id").Immutable(),
		field.Int("follower_id").Immutable(),
	}
}

// Edges of the UserUser.
func (UserUser) Edges() []ent.Edge {
	return nil
}
