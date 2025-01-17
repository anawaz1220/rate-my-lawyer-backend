// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/mvptruck"
)

// MvpTruckCreate is the builder for creating a MvpTruck entity.
type MvpTruckCreate struct {
	config
	mutation *MvpTruckMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mtc *MvpTruckCreate) SetName(s string) *MvpTruckCreate {
	mtc.mutation.SetName(s)
	return mtc
}

// Mutation returns the MvpTruckMutation object of the builder.
func (mtc *MvpTruckCreate) Mutation() *MvpTruckMutation {
	return mtc.mutation
}

// Save creates the MvpTruck in the database.
func (mtc *MvpTruckCreate) Save(ctx context.Context) (*MvpTruck, error) {
	return withHooks(ctx, mtc.sqlSave, mtc.mutation, mtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MvpTruckCreate) SaveX(ctx context.Context) *MvpTruck {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtc *MvpTruckCreate) Exec(ctx context.Context) error {
	_, err := mtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtc *MvpTruckCreate) ExecX(ctx context.Context) {
	if err := mtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MvpTruckCreate) check() error {
	if _, ok := mtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MvpTruck.name"`)}
	}
	return nil
}

func (mtc *MvpTruckCreate) sqlSave(ctx context.Context) (*MvpTruck, error) {
	if err := mtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mtc.mutation.id = &_node.ID
	mtc.mutation.done = true
	return _node, nil
}

func (mtc *MvpTruckCreate) createSpec() (*MvpTruck, *sqlgraph.CreateSpec) {
	var (
		_node = &MvpTruck{config: mtc.config}
		_spec = sqlgraph.NewCreateSpec(mvptruck.Table, sqlgraph.NewFieldSpec(mvptruck.FieldID, field.TypeInt))
	)
	if value, ok := mtc.mutation.Name(); ok {
		_spec.SetField(mvptruck.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// MvpTruckCreateBulk is the builder for creating many MvpTruck entities in bulk.
type MvpTruckCreateBulk struct {
	config
	err      error
	builders []*MvpTruckCreate
}

// Save creates the MvpTruck entities in the database.
func (mtcb *MvpTruckCreateBulk) Save(ctx context.Context) ([]*MvpTruck, error) {
	if mtcb.err != nil {
		return nil, mtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*MvpTruck, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MvpTruckMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MvpTruckCreateBulk) SaveX(ctx context.Context) []*MvpTruck {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtcb *MvpTruckCreateBulk) Exec(ctx context.Context) error {
	_, err := mtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtcb *MvpTruckCreateBulk) ExecX(ctx context.Context) {
	if err := mtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
