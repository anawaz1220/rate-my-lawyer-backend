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

// JurisdictionUpdate is the builder for updating Jurisdiction entities.
type JurisdictionUpdate struct {
	config
	hooks    []Hook
	mutation *JurisdictionMutation
}

// Where appends a list predicates to the JurisdictionUpdate builder.
func (ju *JurisdictionUpdate) Where(ps ...predicate.Jurisdiction) *JurisdictionUpdate {
	ju.mutation.Where(ps...)
	return ju
}

// SetName sets the "name" field.
func (ju *JurisdictionUpdate) SetName(s string) *JurisdictionUpdate {
	ju.mutation.SetName(s)
	return ju
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ju *JurisdictionUpdate) SetNillableName(s *string) *JurisdictionUpdate {
	if s != nil {
		ju.SetName(*s)
	}
	return ju
}

// AddLawyerIDs adds the "lawyers" edge to the Lawyer entity by IDs.
func (ju *JurisdictionUpdate) AddLawyerIDs(ids ...int) *JurisdictionUpdate {
	ju.mutation.AddLawyerIDs(ids...)
	return ju
}

// AddLawyers adds the "lawyers" edges to the Lawyer entity.
func (ju *JurisdictionUpdate) AddLawyers(l ...*Lawyer) *JurisdictionUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ju.AddLawyerIDs(ids...)
}

// AddLawyerJurisdictionIDs adds the "lawyer_jurisdictions" edge to the LawyerJurisdiction entity by IDs.
func (ju *JurisdictionUpdate) AddLawyerJurisdictionIDs(ids ...int) *JurisdictionUpdate {
	ju.mutation.AddLawyerJurisdictionIDs(ids...)
	return ju
}

// AddLawyerJurisdictions adds the "lawyer_jurisdictions" edges to the LawyerJurisdiction entity.
func (ju *JurisdictionUpdate) AddLawyerJurisdictions(l ...*LawyerJurisdiction) *JurisdictionUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ju.AddLawyerJurisdictionIDs(ids...)
}

// Mutation returns the JurisdictionMutation object of the builder.
func (ju *JurisdictionUpdate) Mutation() *JurisdictionMutation {
	return ju.mutation
}

// ClearLawyers clears all "lawyers" edges to the Lawyer entity.
func (ju *JurisdictionUpdate) ClearLawyers() *JurisdictionUpdate {
	ju.mutation.ClearLawyers()
	return ju
}

// RemoveLawyerIDs removes the "lawyers" edge to Lawyer entities by IDs.
func (ju *JurisdictionUpdate) RemoveLawyerIDs(ids ...int) *JurisdictionUpdate {
	ju.mutation.RemoveLawyerIDs(ids...)
	return ju
}

// RemoveLawyers removes "lawyers" edges to Lawyer entities.
func (ju *JurisdictionUpdate) RemoveLawyers(l ...*Lawyer) *JurisdictionUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ju.RemoveLawyerIDs(ids...)
}

// ClearLawyerJurisdictions clears all "lawyer_jurisdictions" edges to the LawyerJurisdiction entity.
func (ju *JurisdictionUpdate) ClearLawyerJurisdictions() *JurisdictionUpdate {
	ju.mutation.ClearLawyerJurisdictions()
	return ju
}

// RemoveLawyerJurisdictionIDs removes the "lawyer_jurisdictions" edge to LawyerJurisdiction entities by IDs.
func (ju *JurisdictionUpdate) RemoveLawyerJurisdictionIDs(ids ...int) *JurisdictionUpdate {
	ju.mutation.RemoveLawyerJurisdictionIDs(ids...)
	return ju
}

// RemoveLawyerJurisdictions removes "lawyer_jurisdictions" edges to LawyerJurisdiction entities.
func (ju *JurisdictionUpdate) RemoveLawyerJurisdictions(l ...*LawyerJurisdiction) *JurisdictionUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ju.RemoveLawyerJurisdictionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ju *JurisdictionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ju.sqlSave, ju.mutation, ju.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JurisdictionUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JurisdictionUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JurisdictionUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ju *JurisdictionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jurisdiction.Table, jurisdiction.Columns, sqlgraph.NewFieldSpec(jurisdiction.FieldID, field.TypeInt))
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.Name(); ok {
		_spec.SetField(jurisdiction.FieldName, field.TypeString, value)
	}
	if ju.mutation.LawyersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jurisdiction.LawyersTable,
			Columns: jurisdiction.LawyersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.RemovedLawyersIDs(); len(nodes) > 0 && !ju.mutation.LawyersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jurisdiction.LawyersTable,
			Columns: jurisdiction.LawyersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.LawyersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jurisdiction.LawyersTable,
			Columns: jurisdiction.LawyersPrimaryKey,
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
	if ju.mutation.LawyerJurisdictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   jurisdiction.LawyerJurisdictionsTable,
			Columns: []string{jurisdiction.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.RemovedLawyerJurisdictionsIDs(); len(nodes) > 0 && !ju.mutation.LawyerJurisdictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   jurisdiction.LawyerJurisdictionsTable,
			Columns: []string{jurisdiction.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.LawyerJurisdictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   jurisdiction.LawyerJurisdictionsTable,
			Columns: []string{jurisdiction.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jurisdiction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ju.mutation.done = true
	return n, nil
}

// JurisdictionUpdateOne is the builder for updating a single Jurisdiction entity.
type JurisdictionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JurisdictionMutation
}

// SetName sets the "name" field.
func (juo *JurisdictionUpdateOne) SetName(s string) *JurisdictionUpdateOne {
	juo.mutation.SetName(s)
	return juo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (juo *JurisdictionUpdateOne) SetNillableName(s *string) *JurisdictionUpdateOne {
	if s != nil {
		juo.SetName(*s)
	}
	return juo
}

// AddLawyerIDs adds the "lawyers" edge to the Lawyer entity by IDs.
func (juo *JurisdictionUpdateOne) AddLawyerIDs(ids ...int) *JurisdictionUpdateOne {
	juo.mutation.AddLawyerIDs(ids...)
	return juo
}

// AddLawyers adds the "lawyers" edges to the Lawyer entity.
func (juo *JurisdictionUpdateOne) AddLawyers(l ...*Lawyer) *JurisdictionUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return juo.AddLawyerIDs(ids...)
}

// AddLawyerJurisdictionIDs adds the "lawyer_jurisdictions" edge to the LawyerJurisdiction entity by IDs.
func (juo *JurisdictionUpdateOne) AddLawyerJurisdictionIDs(ids ...int) *JurisdictionUpdateOne {
	juo.mutation.AddLawyerJurisdictionIDs(ids...)
	return juo
}

// AddLawyerJurisdictions adds the "lawyer_jurisdictions" edges to the LawyerJurisdiction entity.
func (juo *JurisdictionUpdateOne) AddLawyerJurisdictions(l ...*LawyerJurisdiction) *JurisdictionUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return juo.AddLawyerJurisdictionIDs(ids...)
}

// Mutation returns the JurisdictionMutation object of the builder.
func (juo *JurisdictionUpdateOne) Mutation() *JurisdictionMutation {
	return juo.mutation
}

// ClearLawyers clears all "lawyers" edges to the Lawyer entity.
func (juo *JurisdictionUpdateOne) ClearLawyers() *JurisdictionUpdateOne {
	juo.mutation.ClearLawyers()
	return juo
}

// RemoveLawyerIDs removes the "lawyers" edge to Lawyer entities by IDs.
func (juo *JurisdictionUpdateOne) RemoveLawyerIDs(ids ...int) *JurisdictionUpdateOne {
	juo.mutation.RemoveLawyerIDs(ids...)
	return juo
}

// RemoveLawyers removes "lawyers" edges to Lawyer entities.
func (juo *JurisdictionUpdateOne) RemoveLawyers(l ...*Lawyer) *JurisdictionUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return juo.RemoveLawyerIDs(ids...)
}

// ClearLawyerJurisdictions clears all "lawyer_jurisdictions" edges to the LawyerJurisdiction entity.
func (juo *JurisdictionUpdateOne) ClearLawyerJurisdictions() *JurisdictionUpdateOne {
	juo.mutation.ClearLawyerJurisdictions()
	return juo
}

// RemoveLawyerJurisdictionIDs removes the "lawyer_jurisdictions" edge to LawyerJurisdiction entities by IDs.
func (juo *JurisdictionUpdateOne) RemoveLawyerJurisdictionIDs(ids ...int) *JurisdictionUpdateOne {
	juo.mutation.RemoveLawyerJurisdictionIDs(ids...)
	return juo
}

// RemoveLawyerJurisdictions removes "lawyer_jurisdictions" edges to LawyerJurisdiction entities.
func (juo *JurisdictionUpdateOne) RemoveLawyerJurisdictions(l ...*LawyerJurisdiction) *JurisdictionUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return juo.RemoveLawyerJurisdictionIDs(ids...)
}

// Where appends a list predicates to the JurisdictionUpdate builder.
func (juo *JurisdictionUpdateOne) Where(ps ...predicate.Jurisdiction) *JurisdictionUpdateOne {
	juo.mutation.Where(ps...)
	return juo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (juo *JurisdictionUpdateOne) Select(field string, fields ...string) *JurisdictionUpdateOne {
	juo.fields = append([]string{field}, fields...)
	return juo
}

// Save executes the query and returns the updated Jurisdiction entity.
func (juo *JurisdictionUpdateOne) Save(ctx context.Context) (*Jurisdiction, error) {
	return withHooks(ctx, juo.sqlSave, juo.mutation, juo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JurisdictionUpdateOne) SaveX(ctx context.Context) *Jurisdiction {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JurisdictionUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JurisdictionUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (juo *JurisdictionUpdateOne) sqlSave(ctx context.Context) (_node *Jurisdiction, err error) {
	_spec := sqlgraph.NewUpdateSpec(jurisdiction.Table, jurisdiction.Columns, sqlgraph.NewFieldSpec(jurisdiction.FieldID, field.TypeInt))
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Jurisdiction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := juo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jurisdiction.FieldID)
		for _, f := range fields {
			if !jurisdiction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jurisdiction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := juo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := juo.mutation.Name(); ok {
		_spec.SetField(jurisdiction.FieldName, field.TypeString, value)
	}
	if juo.mutation.LawyersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jurisdiction.LawyersTable,
			Columns: jurisdiction.LawyersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.RemovedLawyersIDs(); len(nodes) > 0 && !juo.mutation.LawyersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jurisdiction.LawyersTable,
			Columns: jurisdiction.LawyersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.LawyersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jurisdiction.LawyersTable,
			Columns: jurisdiction.LawyersPrimaryKey,
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
	if juo.mutation.LawyerJurisdictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   jurisdiction.LawyerJurisdictionsTable,
			Columns: []string{jurisdiction.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.RemovedLawyerJurisdictionsIDs(); len(nodes) > 0 && !juo.mutation.LawyerJurisdictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   jurisdiction.LawyerJurisdictionsTable,
			Columns: []string{jurisdiction.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.LawyerJurisdictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   jurisdiction.LawyerJurisdictionsTable,
			Columns: []string{jurisdiction.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Jurisdiction{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jurisdiction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	juo.mutation.done = true
	return _node, nil
}
