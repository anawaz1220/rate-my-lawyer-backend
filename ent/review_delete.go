// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/review"
)

// ReviewDelete is the builder for deleting a Review entity.
type ReviewDelete struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewDelete builder.
func (rd *ReviewDelete) Where(ps ...predicate.Review) *ReviewDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ReviewDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ReviewDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ReviewDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(review.Table, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// ReviewDeleteOne is the builder for deleting a single Review entity.
type ReviewDeleteOne struct {
	rd *ReviewDelete
}

// Where appends a list predicates to the ReviewDelete builder.
func (rdo *ReviewDeleteOne) Where(ps ...predicate.Review) *ReviewDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *ReviewDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{review.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ReviewDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
