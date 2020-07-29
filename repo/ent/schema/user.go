package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("user_name").NotEmpty().MaxLen(255).Unique(),
		field.String("password").NotEmpty(),
		field.String("school").Optional(),
		field.Int("level").Default(0).Positive(),
		field.Int("likes").Default(0).Positive(),
		field.Int("views").Default(0).Positive(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
