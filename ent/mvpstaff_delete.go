// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/mvpstaff"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// MvpStaffDelete is the builder for deleting a MvpStaff entity.
type MvpStaffDelete struct {
	config
	hooks    []Hook
	mutation *MvpStaffMutation
}

// Where appends a list predicates to the MvpStaffDelete builder.
func (msd *MvpStaffDelete) Where(ps ...predicate.MvpStaff) *MvpStaffDelete {
	msd.mutation.Where(ps...)
	return msd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (msd *MvpStaffDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, msd.sqlExec, msd.mutation, msd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (msd *MvpStaffDelete) ExecX(ctx context.Context) int {
	n, err := msd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (msd *MvpStaffDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(mvpstaff.Table, sqlgraph.NewFieldSpec(mvpstaff.FieldID, field.TypeInt))
	if ps := msd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, msd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	msd.mutation.done = true
	return affected, err
}

// MvpStaffDeleteOne is the builder for deleting a single MvpStaff entity.
type MvpStaffDeleteOne struct {
	msd *MvpStaffDelete
}

// Where appends a list predicates to the MvpStaffDelete builder.
func (msdo *MvpStaffDeleteOne) Where(ps ...predicate.MvpStaff) *MvpStaffDeleteOne {
	msdo.msd.mutation.Where(ps...)
	return msdo
}

// Exec executes the deletion query.
func (msdo *MvpStaffDeleteOne) Exec(ctx context.Context) error {
	n, err := msdo.msd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mvpstaff.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (msdo *MvpStaffDeleteOne) ExecX(ctx context.Context) {
	if err := msdo.Exec(ctx); err != nil {
		panic(err)
	}
}
