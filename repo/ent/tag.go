// Code generated by entc, DO NOT EDIT.

package ent

import (
	"TH0400/repo/ent/tag"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TagName holds the value of the "tag_name" field.
	TagName string `json:"tag_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Baned holds the value of the "baned" field.
	Baned bool `json:"baned,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // tag_name
		&sql.NullTime{},   // created_at
		&sql.NullBool{},   // baned
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(values ...interface{}) error {
	if m, n := len(values), len(tag.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tag_name", values[0])
	} else if value.Valid {
		t.TagName = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[1])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field baned", values[2])
	} else if value.Valid {
		t.Baned = value.Bool
	}
	return nil
}

// Update returns a builder for updating this Tag.
// Note that, you need to call Tag.Unwrap() before calling this method, if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return (&TagClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", tag_name=")
	builder.WriteString(t.TagName)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", baned=")
	builder.WriteString(fmt.Sprintf("%v", t.Baned))
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag

func (t Tags) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
