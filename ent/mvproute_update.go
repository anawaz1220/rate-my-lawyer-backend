// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/mvproute"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// MvpRouteUpdate is the builder for updating MvpRoute entities.
type MvpRouteUpdate struct {
	config
	hooks    []Hook
	mutation *MvpRouteMutation
}

// Where appends a list predicates to the MvpRouteUpdate builder.
func (mru *MvpRouteUpdate) Where(ps ...predicate.MvpRoute) *MvpRouteUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// SetName sets the "name" field.
func (mru *MvpRouteUpdate) SetName(s string) *MvpRouteUpdate {
	mru.mutation.SetName(s)
	return mru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mru *MvpRouteUpdate) SetNillableName(s *string) *MvpRouteUpdate {
	if s != nil {
		mru.SetName(*s)
	}
	return mru
}

// SetDayOfWeek sets the "day_of_week" field.
func (mru *MvpRouteUpdate) SetDayOfWeek(s string) *MvpRouteUpdate {
	mru.mutation.SetDayOfWeek(s)
	return mru
}

// SetNillableDayOfWeek sets the "day_of_week" field if the given value is not nil.
func (mru *MvpRouteUpdate) SetNillableDayOfWeek(s *string) *MvpRouteUpdate {
	if s != nil {
		mru.SetDayOfWeek(*s)
	}
	return mru
}

// Mutation returns the MvpRouteMutation object of the builder.
func (mru *MvpRouteUpdate) Mutation() *MvpRouteMutation {
	return mru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MvpRouteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mru.sqlSave, mru.mutation, mru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MvpRouteUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MvpRouteUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MvpRouteUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mru *MvpRouteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mvproute.Table, mvproute.Columns, sqlgraph.NewFieldSpec(mvproute.FieldID, field.TypeInt))
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mru.mutation.Name(); ok {
		_spec.SetField(mvproute.FieldName, field.TypeString, value)
	}
	if value, ok := mru.mutation.DayOfWeek(); ok {
		_spec.SetField(mvproute.FieldDayOfWeek, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mvproute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mru.mutation.done = true
	return n, nil
}

// MvpRouteUpdateOne is the builder for updating a single MvpRoute entity.
type MvpRouteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MvpRouteMutation
}

// SetName sets the "name" field.
func (mruo *MvpRouteUpdateOne) SetName(s string) *MvpRouteUpdateOne {
	mruo.mutation.SetName(s)
	return mruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mruo *MvpRouteUpdateOne) SetNillableName(s *string) *MvpRouteUpdateOne {
	if s != nil {
		mruo.SetName(*s)
	}
	return mruo
}

// SetDayOfWeek sets the "day_of_week" field.
func (mruo *MvpRouteUpdateOne) SetDayOfWeek(s string) *MvpRouteUpdateOne {
	mruo.mutation.SetDayOfWeek(s)
	return mruo
}

// SetNillableDayOfWeek sets the "day_of_week" field if the given value is not nil.
func (mruo *MvpRouteUpdateOne) SetNillableDayOfWeek(s *string) *MvpRouteUpdateOne {
	if s != nil {
		mruo.SetDayOfWeek(*s)
	}
	return mruo
}

// Mutation returns the MvpRouteMutation object of the builder.
func (mruo *MvpRouteUpdateOne) Mutation() *MvpRouteMutation {
	return mruo.mutation
}

// Where appends a list predicates to the MvpRouteUpdate builder.
func (mruo *MvpRouteUpdateOne) Where(ps ...predicate.MvpRoute) *MvpRouteUpdateOne {
	mruo.mutation.Where(ps...)
	return mruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MvpRouteUpdateOne) Select(field string, fields ...string) *MvpRouteUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MvpRoute entity.
func (mruo *MvpRouteUpdateOne) Save(ctx context.Context) (*MvpRoute, error) {
	return withHooks(ctx, mruo.sqlSave, mruo.mutation, mruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MvpRouteUpdateOne) SaveX(ctx context.Context) *MvpRoute {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MvpRouteUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MvpRouteUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mruo *MvpRouteUpdateOne) sqlSave(ctx context.Context) (_node *MvpRoute, err error) {
	_spec := sqlgraph.NewUpdateSpec(mvproute.Table, mvproute.Columns, sqlgraph.NewFieldSpec(mvproute.FieldID, field.TypeInt))
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MvpRoute.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mvproute.FieldID)
		for _, f := range fields {
			if !mvproute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mvproute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mruo.mutation.Name(); ok {
		_spec.SetField(mvproute.FieldName, field.TypeString, value)
	}
	if value, ok := mruo.mutation.DayOfWeek(); ok {
		_spec.SetField(mvproute.FieldDayOfWeek, field.TypeString, value)
	}
	_node = &MvpRoute{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mvproute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mruo.mutation.done = true
	return _node, nil
}