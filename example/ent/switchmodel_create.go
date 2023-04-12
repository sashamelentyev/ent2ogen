// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/switchmodel"
)

// SwitchModelCreate is the builder for creating a SwitchModel entity.
type SwitchModelCreate struct {
	config
	mutation *SwitchModelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (smc *SwitchModelCreate) SetName(s string) *SwitchModelCreate {
	smc.mutation.SetName(s)
	return smc
}

// SetSwitchType sets the "switch_type" field.
func (smc *SwitchModelCreate) SetSwitchType(st switchmodel.SwitchType) *SwitchModelCreate {
	smc.mutation.SetSwitchType(st)
	return smc
}

// SetID sets the "id" field.
func (smc *SwitchModelCreate) SetID(i int64) *SwitchModelCreate {
	smc.mutation.SetID(i)
	return smc
}

// Mutation returns the SwitchModelMutation object of the builder.
func (smc *SwitchModelCreate) Mutation() *SwitchModelMutation {
	return smc.mutation
}

// Save creates the SwitchModel in the database.
func (smc *SwitchModelCreate) Save(ctx context.Context) (*SwitchModel, error) {
	return withHooks[*SwitchModel, SwitchModelMutation](ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SwitchModelCreate) SaveX(ctx context.Context) *SwitchModel {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SwitchModelCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SwitchModelCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SwitchModelCreate) check() error {
	if _, ok := smc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SwitchModel.name"`)}
	}
	if v, ok := smc.mutation.Name(); ok {
		if err := switchmodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SwitchModel.name": %w`, err)}
		}
	}
	if _, ok := smc.mutation.SwitchType(); !ok {
		return &ValidationError{Name: "switch_type", err: errors.New(`ent: missing required field "SwitchModel.switch_type"`)}
	}
	if v, ok := smc.mutation.SwitchType(); ok {
		if err := switchmodel.SwitchTypeValidator(v); err != nil {
			return &ValidationError{Name: "switch_type", err: fmt.Errorf(`ent: validator failed for field "SwitchModel.switch_type": %w`, err)}
		}
	}
	return nil
}

func (smc *SwitchModelCreate) sqlSave(ctx context.Context) (*SwitchModel, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *SwitchModelCreate) createSpec() (*SwitchModel, *sqlgraph.CreateSpec) {
	var (
		_node = &SwitchModel{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(switchmodel.Table, sqlgraph.NewFieldSpec(switchmodel.FieldID, field.TypeInt64))
	)
	if id, ok := smc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smc.mutation.Name(); ok {
		_spec.SetField(switchmodel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := smc.mutation.SwitchType(); ok {
		_spec.SetField(switchmodel.FieldSwitchType, field.TypeEnum, value)
		_node.SwitchType = value
	}
	return _node, _spec
}

// SwitchModelCreateBulk is the builder for creating many SwitchModel entities in bulk.
type SwitchModelCreateBulk struct {
	config
	builders []*SwitchModelCreate
}

// Save creates the SwitchModel entities in the database.
func (smcb *SwitchModelCreateBulk) Save(ctx context.Context) ([]*SwitchModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SwitchModel, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SwitchModelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SwitchModelCreateBulk) SaveX(ctx context.Context) []*SwitchModel {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SwitchModelCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SwitchModelCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
