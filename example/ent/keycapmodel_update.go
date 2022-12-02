// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/example/ent/keycapmodel"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
)

// KeycapModelUpdate is the builder for updating KeycapModel entities.
type KeycapModelUpdate struct {
	config
	hooks    []Hook
	mutation *KeycapModelMutation
}

// Where appends a list predicates to the KeycapModelUpdate builder.
func (kmu *KeycapModelUpdate) Where(ps ...predicate.KeycapModel) *KeycapModelUpdate {
	kmu.mutation.Where(ps...)
	return kmu
}

// SetName sets the "name" field.
func (kmu *KeycapModelUpdate) SetName(s string) *KeycapModelUpdate {
	kmu.mutation.SetName(s)
	return kmu
}

// SetProfile sets the "profile" field.
func (kmu *KeycapModelUpdate) SetProfile(s string) *KeycapModelUpdate {
	kmu.mutation.SetProfile(s)
	return kmu
}

// SetMaterial sets the "material" field.
func (kmu *KeycapModelUpdate) SetMaterial(k keycapmodel.Material) *KeycapModelUpdate {
	kmu.mutation.SetMaterial(k)
	return kmu
}

// Mutation returns the KeycapModelMutation object of the builder.
func (kmu *KeycapModelUpdate) Mutation() *KeycapModelMutation {
	return kmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kmu *KeycapModelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kmu.hooks) == 0 {
		if err = kmu.check(); err != nil {
			return 0, err
		}
		affected, err = kmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeycapModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kmu.check(); err != nil {
				return 0, err
			}
			kmu.mutation = mutation
			affected, err = kmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kmu.hooks) - 1; i >= 0; i-- {
			if kmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kmu *KeycapModelUpdate) SaveX(ctx context.Context) int {
	affected, err := kmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kmu *KeycapModelUpdate) Exec(ctx context.Context) error {
	_, err := kmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kmu *KeycapModelUpdate) ExecX(ctx context.Context) {
	if err := kmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kmu *KeycapModelUpdate) check() error {
	if v, ok := kmu.mutation.Name(); ok {
		if err := keycapmodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KeycapModel.name": %w`, err)}
		}
	}
	if v, ok := kmu.mutation.Material(); ok {
		if err := keycapmodel.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "KeycapModel.material": %w`, err)}
		}
	}
	return nil
}

func (kmu *KeycapModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keycapmodel.Table,
			Columns: keycapmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: keycapmodel.FieldID,
			},
		},
	}
	if ps := kmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kmu.mutation.Name(); ok {
		_spec.SetField(keycapmodel.FieldName, field.TypeString, value)
	}
	if value, ok := kmu.mutation.Profile(); ok {
		_spec.SetField(keycapmodel.FieldProfile, field.TypeString, value)
	}
	if value, ok := kmu.mutation.Material(); ok {
		_spec.SetField(keycapmodel.FieldMaterial, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keycapmodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// KeycapModelUpdateOne is the builder for updating a single KeycapModel entity.
type KeycapModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KeycapModelMutation
}

// SetName sets the "name" field.
func (kmuo *KeycapModelUpdateOne) SetName(s string) *KeycapModelUpdateOne {
	kmuo.mutation.SetName(s)
	return kmuo
}

// SetProfile sets the "profile" field.
func (kmuo *KeycapModelUpdateOne) SetProfile(s string) *KeycapModelUpdateOne {
	kmuo.mutation.SetProfile(s)
	return kmuo
}

// SetMaterial sets the "material" field.
func (kmuo *KeycapModelUpdateOne) SetMaterial(k keycapmodel.Material) *KeycapModelUpdateOne {
	kmuo.mutation.SetMaterial(k)
	return kmuo
}

// Mutation returns the KeycapModelMutation object of the builder.
func (kmuo *KeycapModelUpdateOne) Mutation() *KeycapModelMutation {
	return kmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kmuo *KeycapModelUpdateOne) Select(field string, fields ...string) *KeycapModelUpdateOne {
	kmuo.fields = append([]string{field}, fields...)
	return kmuo
}

// Save executes the query and returns the updated KeycapModel entity.
func (kmuo *KeycapModelUpdateOne) Save(ctx context.Context) (*KeycapModel, error) {
	var (
		err  error
		node *KeycapModel
	)
	if len(kmuo.hooks) == 0 {
		if err = kmuo.check(); err != nil {
			return nil, err
		}
		node, err = kmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeycapModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kmuo.check(); err != nil {
				return nil, err
			}
			kmuo.mutation = mutation
			node, err = kmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kmuo.hooks) - 1; i >= 0; i-- {
			if kmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*KeycapModel)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KeycapModelMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kmuo *KeycapModelUpdateOne) SaveX(ctx context.Context) *KeycapModel {
	node, err := kmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kmuo *KeycapModelUpdateOne) Exec(ctx context.Context) error {
	_, err := kmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kmuo *KeycapModelUpdateOne) ExecX(ctx context.Context) {
	if err := kmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kmuo *KeycapModelUpdateOne) check() error {
	if v, ok := kmuo.mutation.Name(); ok {
		if err := keycapmodel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KeycapModel.name": %w`, err)}
		}
	}
	if v, ok := kmuo.mutation.Material(); ok {
		if err := keycapmodel.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "KeycapModel.material": %w`, err)}
		}
	}
	return nil
}

func (kmuo *KeycapModelUpdateOne) sqlSave(ctx context.Context) (_node *KeycapModel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keycapmodel.Table,
			Columns: keycapmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: keycapmodel.FieldID,
			},
		},
	}
	id, ok := kmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "KeycapModel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keycapmodel.FieldID)
		for _, f := range fields {
			if !keycapmodel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != keycapmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kmuo.mutation.Name(); ok {
		_spec.SetField(keycapmodel.FieldName, field.TypeString, value)
	}
	if value, ok := kmuo.mutation.Profile(); ok {
		_spec.SetField(keycapmodel.FieldProfile, field.TypeString, value)
	}
	if value, ok := kmuo.mutation.Material(); ok {
		_spec.SetField(keycapmodel.FieldMaterial, field.TypeEnum, value)
	}
	_node = &KeycapModel{config: kmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keycapmodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
