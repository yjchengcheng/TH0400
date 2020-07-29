// Code generated by entc, DO NOT EDIT.

package ent

import (
	"TH0400/repo/ent/answer"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Answer is the model entity for the Answer schema.
type Answer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// CreaterID holds the value of the "creater_id" field.
	CreaterID int `json:"creater_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Answer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // content
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullBool{},   // is_deleted
		&sql.NullInt64{},  // creater_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Answer fields.
func (a *Answer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(answer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field content", values[0])
	} else if value.Valid {
		a.Content = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[1])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[2])
	} else if value.Valid {
		a.UpdatedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field is_deleted", values[3])
	} else if value.Valid {
		a.IsDeleted = value.Bool
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field creater_id", values[4])
	} else if value.Valid {
		a.CreaterID = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this Answer.
// Note that, you need to call Answer.Unwrap() before calling this method, if this Answer
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Answer) Update() *AnswerUpdateOne {
	return (&AnswerClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Answer) Unwrap() *Answer {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Answer is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Answer) String() string {
	var builder strings.Builder
	builder.WriteString("Answer(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", content=")
	builder.WriteString(a.Content)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", a.IsDeleted))
	builder.WriteString(", creater_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CreaterID))
	builder.WriteByte(')')
	return builder.String()
}

// Answers is a parsable slice of Answer.
type Answers []*Answer

func (a Answers) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}