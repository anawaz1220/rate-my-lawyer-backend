// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/jurisdiction"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/lawyerjurisdiction"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// LawyerJurisdictionUpdate is the builder for updating LawyerJurisdiction entities.
type LawyerJurisdictionUpdate struct {
	config
	hooks    []Hook
	mutation *LawyerJurisdictionMutation
}

// Where appends a list predicates to the LawyerJurisdictionUpdate builder.
func (lju *LawyerJurisdictionUpdate) Where(ps ...predicate.LawyerJurisdiction) *LawyerJurisdictionUpdate {
	lju.mutation.Where(ps...)
	return lju
}

// SetLawyerID sets the "lawyer_id" field.
func (lju *LawyerJurisdictionUpdate) SetLawyerID(i int) *LawyerJurisdictionUpdate {
	lju.mutation.SetLawyerID(i)
	return lju
}

// SetNillableLawyerID sets the "lawyer_id" field if the given value is not nil.
func (lju *LawyerJurisdictionUpdate) SetNillableLawyerID(i *int) *LawyerJurisdictionUpdate {
	if i != nil {
		lju.SetLawyerID(*i)
	}
	return lju
}

// SetJurisdictionID sets the "jurisdiction_id" field.
func (lju *LawyerJurisdictionUpdate) SetJurisdictionID(i int) *LawyerJurisdictionUpdate {
	lju.mutation.SetJurisdictionID(i)
	return lju
}

// SetNillableJurisdictionID sets the "jurisdiction_id" field if the given value is not nil.
func (lju *LawyerJurisdictionUpdate) SetNillableJurisdictionID(i *int) *LawyerJurisdictionUpdate {
	if i != nil {
		lju.SetJurisdictionID(*i)
	}
	return lju
}

// SetLawyer sets the "lawyer" edge to the Lawyer entity.
func (lju *LawyerJurisdictionUpdate) SetLawyer(l *Lawyer) *LawyerJurisdictionUpdate {
	return lju.SetLawyerID(l.ID)
}

// SetJurisdiction sets the "jurisdiction" edge to the Jurisdiction entity.
func (lju *LawyerJurisdictionUpdate) SetJurisdiction(j *Jurisdiction) *LawyerJurisdictionUpdate {
	return lju.SetJurisdictionID(j.ID)
}

// Mutation returns the LawyerJurisdictionMutation object of the builder.
func (lju *LawyerJurisdictionUpdate) Mutation() *LawyerJurisdictionMutation {
	return lju.mutation
}

// ClearLawyer clears the "lawyer" edge to the Lawyer entity.
func (lju *LawyerJurisdictionUpdate) ClearLawyer() *LawyerJurisdictionUpdate {
	lju.mutation.ClearLawyer()
	return lju
}

// ClearJurisdiction clears the "jurisdiction" edge to the Jurisdiction entity.
func (lju *LawyerJurisdictionUpdate) ClearJurisdiction() *LawyerJurisdictionUpdate {
	lju.mutation.ClearJurisdiction()
	return lju
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lju *LawyerJurisdictionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lju.sqlSave, lju.mutation, lju.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lju *LawyerJurisdictionUpdate) SaveX(ctx context.Context) int {
	affected, err := lju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lju *LawyerJurisdictionUpdate) Exec(ctx context.Context) error {
	_, err := lju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lju *LawyerJurisdictionUpdate) ExecX(ctx context.Context) {
	if err := lju.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lju *LawyerJurisdictionUpdate) check() error {
	if _, ok := lju.mutation.LawyerID(); lju.mutation.LawyerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LawyerJurisdiction.lawyer"`)
	}
	if _, ok := lju.mutation.JurisdictionID(); lju.mutation.JurisdictionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LawyerJurisdiction.jurisdiction"`)
	}
	return nil
}

func (lju *LawyerJurisdictionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lju.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lawyerjurisdiction.Table, lawyerjurisdiction.Columns, sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt))
	if ps := lju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if lju.mutation.LawyerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lju.mutation.LawyerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lju.mutation.JurisdictionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lju.mutation.JurisdictionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lawyerjurisdiction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lju.mutation.done = true
	return n, nil
}

// LawyerJurisdictionUpdateOne is the builder for updating a single LawyerJurisdiction entity.
type LawyerJurisdictionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LawyerJurisdictionMutation
}

// SetLawyerID sets the "lawyer_id" field.
func (ljuo *LawyerJurisdictionUpdateOne) SetLawyerID(i int) *LawyerJurisdictionUpdateOne {
	ljuo.mutation.SetLawyerID(i)
	return ljuo
}

// SetNillableLawyerID sets the "lawyer_id" field if the given value is not nil.
func (ljuo *LawyerJurisdictionUpdateOne) SetNillableLawyerID(i *int) *LawyerJurisdictionUpdateOne {
	if i != nil {
		ljuo.SetLawyerID(*i)
	}
	return ljuo
}

// SetJurisdictionID sets the "jurisdiction_id" field.
func (ljuo *LawyerJurisdictionUpdateOne) SetJurisdictionID(i int) *LawyerJurisdictionUpdateOne {
	ljuo.mutation.SetJurisdictionID(i)
	return ljuo
}

// SetNillableJurisdictionID sets the "jurisdiction_id" field if the given value is not nil.
func (ljuo *LawyerJurisdictionUpdateOne) SetNillableJurisdictionID(i *int) *LawyerJurisdictionUpdateOne {
	if i != nil {
		ljuo.SetJurisdictionID(*i)
	}
	return ljuo
}

// SetLawyer sets the "lawyer" edge to the Lawyer entity.
func (ljuo *LawyerJurisdictionUpdateOne) SetLawyer(l *Lawyer) *LawyerJurisdictionUpdateOne {
	return ljuo.SetLawyerID(l.ID)
}

// SetJurisdiction sets the "jurisdiction" edge to the Jurisdiction entity.
func (ljuo *LawyerJurisdictionUpdateOne) SetJurisdiction(j *Jurisdiction) *LawyerJurisdictionUpdateOne {
	return ljuo.SetJurisdictionID(j.ID)
}

// Mutation returns the LawyerJurisdictionMutation object of the builder.
func (ljuo *LawyerJurisdictionUpdateOne) Mutation() *LawyerJurisdictionMutation {
	return ljuo.mutation
}

// ClearLawyer clears the "lawyer" edge to the Lawyer entity.
func (ljuo *LawyerJurisdictionUpdateOne) ClearLawyer() *LawyerJurisdictionUpdateOne {
	ljuo.mutation.ClearLawyer()
	return ljuo
}

// ClearJurisdiction clears the "jurisdiction" edge to the Jurisdiction entity.
func (ljuo *LawyerJurisdictionUpdateOne) ClearJurisdiction() *LawyerJurisdictionUpdateOne {
	ljuo.mutation.ClearJurisdiction()
	return ljuo
}

// Where appends a list predicates to the LawyerJurisdictionUpdate builder.
func (ljuo *LawyerJurisdictionUpdateOne) Where(ps ...predicate.LawyerJurisdiction) *LawyerJurisdictionUpdateOne {
	ljuo.mutation.Where(ps...)
	return ljuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ljuo *LawyerJurisdictionUpdateOne) Select(field string, fields ...string) *LawyerJurisdictionUpdateOne {
	ljuo.fields = append([]string{field}, fields...)
	return ljuo
}

// Save executes the query and returns the updated LawyerJurisdiction entity.
func (ljuo *LawyerJurisdictionUpdateOne) Save(ctx context.Context) (*LawyerJurisdiction, error) {
	return withHooks(ctx, ljuo.sqlSave, ljuo.mutation, ljuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ljuo *LawyerJurisdictionUpdateOne) SaveX(ctx context.Context) *LawyerJurisdiction {
	node, err := ljuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ljuo *LawyerJurisdictionUpdateOne) Exec(ctx context.Context) error {
	_, err := ljuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ljuo *LawyerJurisdictionUpdateOne) ExecX(ctx context.Context) {
	if err := ljuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ljuo *LawyerJurisdictionUpdateOne) check() error {
	if _, ok := ljuo.mutation.LawyerID(); ljuo.mutation.LawyerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LawyerJurisdiction.lawyer"`)
	}
	if _, ok := ljuo.mutation.JurisdictionID(); ljuo.mutation.JurisdictionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LawyerJurisdiction.jurisdiction"`)
	}
	return nil
}

func (ljuo *LawyerJurisdictionUpdateOne) sqlSave(ctx context.Context) (_node *LawyerJurisdiction, err error) {
	if err := ljuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lawyerjurisdiction.Table, lawyerjurisdiction.Columns, sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt))
	id, ok := ljuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LawyerJurisdiction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ljuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lawyerjurisdiction.FieldID)
		for _, f := range fields {
			if !lawyerjurisdiction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lawyerjurisdiction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ljuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ljuo.mutation.LawyerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ljuo.mutation.LawyerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ljuo.mutation.JurisdictionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ljuo.mutation.JurisdictionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LawyerJurisdiction{config: ljuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ljuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lawyerjurisdiction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ljuo.mutation.done = true
	return _node, nil
}