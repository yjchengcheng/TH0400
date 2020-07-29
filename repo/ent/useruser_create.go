// Code generated by entc, DO NOT EDIT.

package ent

import (
	"TH0400/repo/ent/useruser"
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUserCreate is the builder for creating a UserUser entity.
type UserUserCreate struct {
	config
	mutation *UserUserMutation
	hooks    []Hook
}

// SetIdolID sets the idol_id field.
func (uuc *UserUserCreate) SetIdolID(i int) *UserUserCreate {
	uuc.mutation.SetIdolID(i)
	return uuc
}

// SetFollowerID sets the follower_id field.
func (uuc *UserUserCreate) SetFollowerID(i int) *UserUserCreate {
	uuc.mutation.SetFollowerID(i)
	return uuc
}

// SetID sets the id field.
func (uuc *UserUserCreate) SetID(i int) *UserUserCreate {
	uuc.mutation.SetID(i)
	return uuc
}

// Mutation returns the UserUserMutation object of the builder.
func (uuc *UserUserCreate) Mutation() *UserUserMutation {
	return uuc.mutation
}

// Save creates the UserUser in the database.
func (uuc *UserUserCreate) Save(ctx context.Context) (*UserUser, error) {
	if _, ok := uuc.mutation.IdolID(); !ok {
		return nil, &ValidationError{Name: "idol_id", err: errors.New("ent: missing required field \"idol_id\"")}
	}
	if _, ok := uuc.mutation.FollowerID(); !ok {
		return nil, &ValidationError{Name: "follower_id", err: errors.New("ent: missing required field \"follower_id\"")}
	}
	var (
		err  error
		node *UserUser
	)
	if len(uuc.hooks) == 0 {
		node, err = uuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuc.mutation = mutation
			node, err = uuc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuc.hooks) - 1; i >= 0; i-- {
			mut = uuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uuc *UserUserCreate) SaveX(ctx context.Context) *UserUser {
	v, err := uuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uuc *UserUserCreate) sqlSave(ctx context.Context) (*UserUser, error) {
	uu, _spec := uuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uuc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if uu.ID == 0 {
		id := _spec.ID.Value.(int64)
		uu.ID = int(id)
	}
	return uu, nil
}

func (uuc *UserUserCreate) createSpec() (*UserUser, *sqlgraph.CreateSpec) {
	var (
		uu    = &UserUser{config: uuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: useruser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useruser.FieldID,
			},
		}
	)
	if id, ok := uuc.mutation.ID(); ok {
		uu.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uuc.mutation.IdolID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: useruser.FieldIdolID,
		})
		uu.IdolID = value
	}
	if value, ok := uuc.mutation.FollowerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: useruser.FieldFollowerID,
		})
		uu.FollowerID = value
	}
	return uu, _spec
}