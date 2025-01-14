// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
)

// KeycapModelCreate is the builder for creating a KeycapModel entity.
type KeycapModelCreate struct {
	config
	mutation *KeycapModelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (kmc *KeycapModelCreate) SetName(s string) *KeycapModelCreate {
	kmc.mutation.SetName(s)
	return kmc
}

// SetProfile sets the "profile" field.
func (kmc *KeycapModelCreate) SetProfile(s string) *KeycapModelCreate {
	kmc.mutation.SetProfile(s)
	return kmc
}

// SetMaterial sets the "material" field.
func (kmc *KeycapModelCreate) SetMaterial(k keycapmodel.Material) *KeycapModelCreate {
	kmc.mutation.SetMaterial(k)
	return kmc
}

// SetID sets the "id" field.
func (kmc *KeycapModelCreate) SetID(i int64) *KeycapModelCreate {
	kmc.mutation.SetID(i)
	return kmc
}

// Mutation returns the KeycapModelMutation object of the builder.
func (kmc *KeycapModelCreate) Mutation() *KeycapModelMutation {
	return kmc.mutation
}

// Save creates the KeycapModel in the database.
func (kmc *KeycapModelCreate) Save(ctx context.Context) (*KeycapModel, error) {
	return withHooks[*KeycapModel, KeycapModelMutation](ctx, kmc.sqlSave, kmc.mutation, kmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kmc *KeycapModelCreate) SaveX(ctx context.Context) *KeycapModel {
	v, err := kmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kmc *KeycapModelCreate) Exec(ctx context.Context) error {
	_, err := kmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kmc *KeycapModelCreate) ExecX(ctx context.Context) {
	if err := kmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kmc *KeycapModelCreate) check() error {
	if _, ok := kmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "KeycapModel.name"`)}
	}
	if v, ok := kmc.mutation.Name(); ok {
		if err := keycapmodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KeycapModel.name": %w`, err)}
		}
	}
	if _, ok := kmc.mutation.Profile(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required field "KeycapModel.profile"`)}
	}
	if _, ok := kmc.mutation.Material(); !ok {
		return &ValidationError{Name: "material", err: errors.New(`ent: missing required field "KeycapModel.material"`)}
	}
	if v, ok := kmc.mutation.Material(); ok {
		if err := keycapmodel.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "KeycapModel.material": %w`, err)}
		}
	}
	return nil
}

func (kmc *KeycapModelCreate) sqlSave(ctx context.Context) (*KeycapModel, error) {
	if err := kmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	kmc.mutation.id = &_node.ID
	kmc.mutation.done = true
	return _node, nil
}

func (kmc *KeycapModelCreate) createSpec() (*KeycapModel, *sqlgraph.CreateSpec) {
	var (
		_node = &KeycapModel{config: kmc.config}
		_spec = sqlgraph.NewCreateSpec(keycapmodel.Table, sqlgraph.NewFieldSpec(keycapmodel.FieldID, field.TypeInt64))
	)
	if id, ok := kmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := kmc.mutation.Name(); ok {
		_spec.SetField(keycapmodel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := kmc.mutation.Profile(); ok {
		_spec.SetField(keycapmodel.FieldProfile, field.TypeString, value)
		_node.Profile = value
	}
	if value, ok := kmc.mutation.Material(); ok {
		_spec.SetField(keycapmodel.FieldMaterial, field.TypeEnum, value)
		_node.Material = value
	}
	return _node, _spec
}

// KeycapModelCreateBulk is the builder for creating many KeycapModel entities in bulk.
type KeycapModelCreateBulk struct {
	config
	builders []*KeycapModelCreate
}

// Save creates the KeycapModel entities in the database.
func (kmcb *KeycapModelCreateBulk) Save(ctx context.Context) ([]*KeycapModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kmcb.builders))
	nodes := make([]*KeycapModel, len(kmcb.builders))
	mutators := make([]Mutator, len(kmcb.builders))
	for i := range kmcb.builders {
		func(i int, root context.Context) {
			builder := kmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeycapModelMutation)
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
					_, err = mutators[i+1].Mutate(root, kmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kmcb *KeycapModelCreateBulk) SaveX(ctx context.Context) []*KeycapModel {
	v, err := kmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kmcb *KeycapModelCreateBulk) Exec(ctx context.Context) error {
	_, err := kmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kmcb *KeycapModelCreateBulk) ExecX(ctx context.Context) {
	if err := kmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
