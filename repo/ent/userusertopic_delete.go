// Code generated by entc, DO NOT EDIT.

package ent

import (
	"TH0400/repo/ent/predicate"
	"TH0400/repo/ent/userusertopic"
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUserTopicDelete is the builder for deleting a UserUserTopic entity.
type UserUserTopicDelete struct {
	config
	hooks      []Hook
	mutation   *UserUserTopicMutation
	predicates []predicate.UserUserTopic
}

// Where adds a new predicate to the delete builder.
func (uutd *UserUserTopicDelete) Where(ps ...predicate.UserUserTopic) *UserUserTopicDelete {
	uutd.predicates = append(uutd.predicates, ps...)
	return uutd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uutd *UserUserTopicDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uutd.hooks) == 0 {
		affected, err = uutd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserUserTopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uutd.mutation = mutation
			affected, err = uutd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uutd.hooks) - 1; i >= 0; i-- {
			mut = uutd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uutd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uutd *UserUserTopicDelete) ExecX(ctx context.Context) int {
	n, err := uutd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uutd *UserUserTopicDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userusertopic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userusertopic.FieldID,
			},
		},
	}
	if ps := uutd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uutd.driver, _spec)
}

// UserUserTopicDeleteOne is the builder for deleting a single UserUserTopic entity.
type UserUserTopicDeleteOne struct {
	uutd *UserUserTopicDelete
}

// Exec executes the deletion query.
func (uutdo *UserUserTopicDeleteOne) Exec(ctx context.Context) error {
	n, err := uutdo.uutd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userusertopic.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uutdo *UserUserTopicDeleteOne) ExecX(ctx context.Context) {
	uutdo.uutd.ExecX(ctx)
}
