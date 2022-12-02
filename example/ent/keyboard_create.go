// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/keyboard"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
	"github.com/ogen-go/ent2ogen/example/ent/switchmodel"
)

// KeyboardCreate is the builder for creating a Keyboard entity.
type KeyboardCreate struct {
	config
	mutation *KeyboardMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (kc *KeyboardCreate) SetName(s string) *KeyboardCreate {
	kc.mutation.SetName(s)
	return kc
}

// SetPrice sets the "price" field.
func (kc *KeyboardCreate) SetPrice(i int64) *KeyboardCreate {
	kc.mutation.SetPrice(i)
	return kc
}

// SetDiscount sets the "discount" field.
func (kc *KeyboardCreate) SetDiscount(i int64) *KeyboardCreate {
	kc.mutation.SetDiscount(i)
	return kc
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (kc *KeyboardCreate) SetNillableDiscount(i *int64) *KeyboardCreate {
	if i != nil {
		kc.SetDiscount(*i)
	}
	return kc
}

// SetID sets the "id" field.
func (kc *KeyboardCreate) SetID(i int64) *KeyboardCreate {
	kc.mutation.SetID(i)
	return kc
}

// SetSwitchesID sets the "switches" edge to the SwitchModel entity by ID.
func (kc *KeyboardCreate) SetSwitchesID(id int64) *KeyboardCreate {
	kc.mutation.SetSwitchesID(id)
	return kc
}

// SetSwitches sets the "switches" edge to the SwitchModel entity.
func (kc *KeyboardCreate) SetSwitches(s *SwitchModel) *KeyboardCreate {
	return kc.SetSwitchesID(s.ID)
}

// SetKeycapsID sets the "keycaps" edge to the KeycapModel entity by ID.
func (kc *KeyboardCreate) SetKeycapsID(id int64) *KeyboardCreate {
	kc.mutation.SetKeycapsID(id)
	return kc
}

// SetKeycaps sets the "keycaps" edge to the KeycapModel entity.
func (kc *KeyboardCreate) SetKeycaps(k *KeycapModel) *KeyboardCreate {
	return kc.SetKeycapsID(k.ID)
}

// Mutation returns the KeyboardMutation object of the builder.
func (kc *KeyboardCreate) Mutation() *KeyboardMutation {
	return kc.mutation
}

// Save creates the Keyboard in the database.
func (kc *KeyboardCreate) Save(ctx context.Context) (*Keyboard, error) {
	var (
		err  error
		node *Keyboard
	)
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeyboardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			if node, err = kc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			if kc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Keyboard)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KeyboardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KeyboardCreate) SaveX(ctx context.Context) *Keyboard {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KeyboardCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KeyboardCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KeyboardCreate) check() error {
	if _, ok := kc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Keyboard.name"`)}
	}
	if v, ok := kc.mutation.Name(); ok {
		if err := keyboard.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Keyboard.name": %w`, err)}
		}
	}
	if _, ok := kc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Keyboard.price"`)}
	}
	if _, ok := kc.mutation.SwitchesID(); !ok {
		return &ValidationError{Name: "switches", err: errors.New(`ent: missing required edge "Keyboard.switches"`)}
	}
	if _, ok := kc.mutation.KeycapsID(); !ok {
		return &ValidationError{Name: "keycaps", err: errors.New(`ent: missing required edge "Keyboard.keycaps"`)}
	}
	return nil
}

func (kc *KeyboardCreate) sqlSave(ctx context.Context) (*Keyboard, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (kc *KeyboardCreate) createSpec() (*Keyboard, *sqlgraph.CreateSpec) {
	var (
		_node = &Keyboard{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: keyboard.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: keyboard.FieldID,
			},
		}
	)
	if id, ok := kc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := kc.mutation.Name(); ok {
		_spec.SetField(keyboard.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := kc.mutation.Price(); ok {
		_spec.SetField(keyboard.FieldPrice, field.TypeInt64, value)
		_node.Price = value
	}
	if value, ok := kc.mutation.Discount(); ok {
		_spec.SetField(keyboard.FieldDiscount, field.TypeInt64, value)
		_node.Discount = &value
	}
	if nodes := kc.mutation.SwitchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   keyboard.SwitchesTable,
			Columns: []string{keyboard.SwitchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: switchmodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.keyboard_switches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kc.mutation.KeycapsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   keyboard.KeycapsTable,
			Columns: []string{keyboard.KeycapsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: keycapmodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.keyboard_keycaps = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KeyboardCreateBulk is the builder for creating many Keyboard entities in bulk.
type KeyboardCreateBulk struct {
	config
	builders []*KeyboardCreate
}

// Save creates the Keyboard entities in the database.
func (kcb *KeyboardCreateBulk) Save(ctx context.Context) ([]*Keyboard, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Keyboard, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeyboardMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KeyboardCreateBulk) SaveX(ctx context.Context) []*Keyboard {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KeyboardCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KeyboardCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}
