// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/jurisdiction"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/lawyerjurisdiction"
)

// LawyerJurisdictionCreate is the builder for creating a LawyerJurisdiction entity.
type LawyerJurisdictionCreate struct {
	config
	mutation *LawyerJurisdictionMutation
	hooks    []Hook
}

// SetLawyerID sets the "lawyer_id" field.
func (ljc *LawyerJurisdictionCreate) SetLawyerID(i int) *LawyerJurisdictionCreate {
	ljc.mutation.SetLawyerID(i)
	return ljc
}

// SetJurisdictionID sets the "jurisdiction_id" field.
func (ljc *LawyerJurisdictionCreate) SetJurisdictionID(i int) *LawyerJurisdictionCreate {
	ljc.mutation.SetJurisdictionID(i)
	return ljc
}

// SetLawyer sets the "lawyer" edge to the Lawyer entity.
func (ljc *LawyerJurisdictionCreate) SetLawyer(l *Lawyer) *LawyerJurisdictionCreate {
	return ljc.SetLawyerID(l.ID)
}

// SetJurisdiction sets the "jurisdiction" edge to the Jurisdiction entity.
func (ljc *LawyerJurisdictionCreate) SetJurisdiction(j *Jurisdiction) *LawyerJurisdictionCreate {
	return ljc.SetJurisdictionID(j.ID)
}

// Mutation returns the LawyerJurisdictionMutation object of the builder.
func (ljc *LawyerJurisdictionCreate) Mutation() *LawyerJurisdictionMutation {
	return ljc.mutation
}

// Save creates the LawyerJurisdiction in the database.
func (ljc *LawyerJurisdictionCreate) Save(ctx context.Context) (*LawyerJurisdiction, error) {
	return withHooks(ctx, ljc.sqlSave, ljc.mutation, ljc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ljc *LawyerJurisdictionCreate) SaveX(ctx context.Context) *LawyerJurisdiction {
	v, err := ljc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ljc *LawyerJurisdictionCreate) Exec(ctx context.Context) error {
	_, err := ljc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ljc *LawyerJurisdictionCreate) ExecX(ctx context.Context) {
	if err := ljc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ljc *LawyerJurisdictionCreate) check() error {
	if _, ok := ljc.mutation.LawyerID(); !ok {
		return &ValidationError{Name: "lawyer_id", err: errors.New(`ent: missing required field "LawyerJurisdiction.lawyer_id"`)}
	}
	if _, ok := ljc.mutation.JurisdictionID(); !ok {
		return &ValidationError{Name: "jurisdiction_id", err: errors.New(`ent: missing required field "LawyerJurisdiction.jurisdiction_id"`)}
	}
	if _, ok := ljc.mutation.LawyerID(); !ok {
		return &ValidationError{Name: "lawyer", err: errors.New(`ent: missing required edge "LawyerJurisdiction.lawyer"`)}
	}
	if _, ok := ljc.mutation.JurisdictionID(); !ok {
		return &ValidationError{Name: "jurisdiction", err: errors.New(`ent: missing required edge "LawyerJurisdiction.jurisdiction"`)}
	}
	return nil
}

func (ljc *LawyerJurisdictionCreate) sqlSave(ctx context.Context) (*LawyerJurisdiction, error) {
	if err := ljc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ljc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ljc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ljc.mutation.id = &_node.ID
	ljc.mutation.done = true
	return _node, nil
}

func (ljc *LawyerJurisdictionCreate) createSpec() (*LawyerJurisdiction, *sqlgraph.CreateSpec) {
	var (
		_node = &LawyerJurisdiction{config: ljc.config}
		_spec = sqlgraph.NewCreateSpec(lawyerjurisdiction.Table, sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt))
	)
	if nodes := ljc.mutation.LawyerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lawyerjurisdiction.LawyerTable,
			Columns: []string{lawyerjurisdiction.LawyerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LawyerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ljc.mutation.JurisdictionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lawyerjurisdiction.JurisdictionTable,
			Columns: []string{lawyerjurisdiction.JurisdictionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.JurisdictionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LawyerJurisdictionCreateBulk is the builder for creating many LawyerJurisdiction entities in bulk.
type LawyerJurisdictionCreateBulk struct {
	config
	err      error
	builders []*LawyerJurisdictionCreate
}

// Save creates the LawyerJurisdiction entities in the database.
func (ljcb *LawyerJurisdictionCreateBulk) Save(ctx context.Context) ([]*LawyerJurisdiction, error) {
	if ljcb.err != nil {
		return nil, ljcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ljcb.builders))
	nodes := make([]*LawyerJurisdiction, len(ljcb.builders))
	mutators := make([]Mutator, len(ljcb.builders))
	for i := range ljcb.builders {
		func(i int, root context.Context) {
			builder := ljcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LawyerJurisdictionMutation)
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
					_, err = mutators[i+1].Mutate(root, ljcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ljcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ljcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ljcb *LawyerJurisdictionCreateBulk) SaveX(ctx context.Context) []*LawyerJurisdiction {
	v, err := ljcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ljcb *LawyerJurisdictionCreateBulk) Exec(ctx context.Context) error {
	_, err := ljcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ljcb *LawyerJurisdictionCreateBulk) ExecX(ctx context.Context) {
	if err := ljcb.Exec(ctx); err != nil {
		panic(err)
	}
}
