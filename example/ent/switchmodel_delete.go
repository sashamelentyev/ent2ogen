// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
	"github.com/ogen-go/ent2ogen/example/ent/switchmodel"
)

// SwitchModelDelete is the builder for deleting a SwitchModel entity.
type SwitchModelDelete struct {
	config
	hooks    []Hook
	mutation *SwitchModelMutation
}

// Where appends a list predicates to the SwitchModelDelete builder.
func (smd *SwitchModelDelete) Where(ps ...predicate.SwitchModel) *SwitchModelDelete {
	smd.mutation.Where(ps...)
	return smd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smd *SwitchModelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(smd.hooks) == 0 {
		affected, err = smd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SwitchModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smd.mutation = mutation
			affected, err = smd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smd.hooks) - 1; i >= 0; i-- {
			if smd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (smd *SwitchModelDelete) ExecX(ctx context.Context) int {
	n, err := smd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smd *SwitchModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: switchmodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: switchmodel.FieldID,
			},
		},
	}
	if ps := smd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, smd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SwitchModelDeleteOne is the builder for deleting a single SwitchModel entity.
type SwitchModelDeleteOne struct {
	smd *SwitchModelDelete
}

// Exec executes the deletion query.
func (smdo *SwitchModelDeleteOne) Exec(ctx context.Context) error {
	n, err := smdo.smd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{switchmodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smdo *SwitchModelDeleteOne) ExecX(ctx context.Context) {
	smdo.smd.ExecX(ctx)
}