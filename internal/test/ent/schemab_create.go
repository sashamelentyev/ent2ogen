// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemab"
)

// SchemaBCreate is the builder for creating a SchemaB entity.
type SchemaBCreate struct {
	config
	mutation *SchemaBMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (sb *SchemaBCreate) SetID(i int64) *SchemaBCreate {
	sb.mutation.SetID(i)
	return sb
}

// Mutation returns the SchemaBMutation object of the builder.
func (sb *SchemaBCreate) Mutation() *SchemaBMutation {
	return sb.mutation
}

// Save creates the SchemaB in the database.
func (sb *SchemaBCreate) Save(ctx context.Context) (*SchemaB, error) {
	return withHooks[*SchemaB, SchemaBMutation](ctx, sb.sqlSave, sb.mutation, sb.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sb *SchemaBCreate) SaveX(ctx context.Context) *SchemaB {
	v, err := sb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sb *SchemaBCreate) Exec(ctx context.Context) error {
	_, err := sb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sb *SchemaBCreate) ExecX(ctx context.Context) {
	if err := sb.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sb *SchemaBCreate) check() error {
	return nil
}

func (sb *SchemaBCreate) sqlSave(ctx context.Context) (*SchemaB, error) {
	if err := sb.check(); err != nil {
		return nil, err
	}
	_node, _spec := sb.createSpec()
	if err := sqlgraph.CreateNode(ctx, sb.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sb.mutation.id = &_node.ID
	sb.mutation.done = true
	return _node, nil
}

func (sb *SchemaBCreate) createSpec() (*SchemaB, *sqlgraph.CreateSpec) {
	var (
		_node = &SchemaB{config: sb.config}
		_spec = sqlgraph.NewCreateSpec(schemab.Table, sqlgraph.NewFieldSpec(schemab.FieldID, field.TypeInt64))
	)
	if id, ok := sb.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	return _node, _spec
}

// SchemaBCreateBulk is the builder for creating many SchemaB entities in bulk.
type SchemaBCreateBulk struct {
	config
	builders []*SchemaBCreate
}

// Save creates the SchemaB entities in the database.
func (sbb *SchemaBCreateBulk) Save(ctx context.Context) ([]*SchemaB, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sbb.builders))
	nodes := make([]*SchemaB, len(sbb.builders))
	mutators := make([]Mutator, len(sbb.builders))
	for i := range sbb.builders {
		func(i int, root context.Context) {
			builder := sbb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SchemaBMutation)
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
					_, err = mutators[i+1].Mutate(root, sbb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sbb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sbb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sbb *SchemaBCreateBulk) SaveX(ctx context.Context) []*SchemaB {
	v, err := sbb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sbb *SchemaBCreateBulk) Exec(ctx context.Context) error {
	_, err := sbb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sbb *SchemaBCreateBulk) ExecX(ctx context.Context) {
	if err := sbb.Exec(ctx); err != nil {
		panic(err)
	}
}
