// Code generated by entc, DO NOT EDIT.

package ent

import (
	"TH0400/repo/ent/predicate"
	"TH0400/repo/ent/topic"
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks      []Hook
	mutation   *TopicMutation
	predicates []predicate.Topic
}

// Where adds a new predicate for the builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetTitle sets the title field.
func (tu *TopicUpdate) SetTitle(s string) *TopicUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetContent sets the content field.
func (tu *TopicUpdate) SetContent(s string) *TopicUpdate {
	tu.mutation.SetContent(s)
	return tu
}

// SetNillableContent sets the content field if the given value is not nil.
func (tu *TopicUpdate) SetNillableContent(s *string) *TopicUpdate {
	if s != nil {
		tu.SetContent(*s)
	}
	return tu
}

// SetIsReleased sets the is_released field.
func (tu *TopicUpdate) SetIsReleased(b bool) *TopicUpdate {
	tu.mutation.SetIsReleased(b)
	return tu
}

// SetUpdatedAt sets the updated_at field.
func (tu *TopicUpdate) SetUpdatedAt(t time.Time) *TopicUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (tu *TopicUpdate) SetNillableUpdatedAt(t *time.Time) *TopicUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Title(); ok {
		if err := topic.TitleValidator(v); err != nil {
			return 0, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldContent,
		})
	}
	if value, ok := tu.mutation.IsReleased(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: topic.FieldIsReleased,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: topic.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// SetTitle sets the title field.
func (tuo *TopicUpdateOne) SetTitle(s string) *TopicUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetContent sets the content field.
func (tuo *TopicUpdateOne) SetContent(s string) *TopicUpdateOne {
	tuo.mutation.SetContent(s)
	return tuo
}

// SetNillableContent sets the content field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableContent(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetContent(*s)
	}
	return tuo
}

// SetIsReleased sets the is_released field.
func (tuo *TopicUpdateOne) SetIsReleased(b bool) *TopicUpdateOne {
	tuo.mutation.SetIsReleased(b)
	return tuo
}

// SetUpdatedAt sets the updated_at field.
func (tuo *TopicUpdateOne) SetUpdatedAt(t time.Time) *TopicUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableUpdatedAt(t *time.Time) *TopicUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// Save executes the query and returns the updated entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	if v, ok := tuo.mutation.Title(); ok {
		if err := topic.TitleValidator(v); err != nil {
			return nil, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	var (
		err  error
		node *Topic
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (t *Topic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Topic.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldContent,
		})
	}
	if value, ok := tuo.mutation.IsReleased(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: topic.FieldIsReleased,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: topic.FieldUpdatedAt,
		})
	}
	t = &Topic{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}