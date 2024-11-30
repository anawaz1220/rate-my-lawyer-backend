// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/review"
	"github.com/mikestefanello/pagoda/ent/rmluser"
)

// RMLUserUpdate is the builder for updating RMLUser entities.
type RMLUserUpdate struct {
	config
	hooks    []Hook
	mutation *RMLUserMutation
}

// Where appends a list predicates to the RMLUserUpdate builder.
func (ruu *RMLUserUpdate) Where(ps ...predicate.RMLUser) *RMLUserUpdate {
	ruu.mutation.Where(ps...)
	return ruu
}

// SetName sets the "name" field.
func (ruu *RMLUserUpdate) SetName(s string) *RMLUserUpdate {
	ruu.mutation.SetName(s)
	return ruu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruu *RMLUserUpdate) SetNillableName(s *string) *RMLUserUpdate {
	if s != nil {
		ruu.SetName(*s)
	}
	return ruu
}

// SetEmail sets the "email" field.
func (ruu *RMLUserUpdate) SetEmail(s string) *RMLUserUpdate {
	ruu.mutation.SetEmail(s)
	return ruu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ruu *RMLUserUpdate) SetNillableEmail(s *string) *RMLUserUpdate {
	if s != nil {
		ruu.SetEmail(*s)
	}
	return ruu
}

// SetPassword sets the "password" field.
func (ruu *RMLUserUpdate) SetPassword(s string) *RMLUserUpdate {
	ruu.mutation.SetPassword(s)
	return ruu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ruu *RMLUserUpdate) SetNillablePassword(s *string) *RMLUserUpdate {
	if s != nil {
		ruu.SetPassword(*s)
	}
	return ruu
}

// SetRole sets the "role" field.
func (ruu *RMLUserUpdate) SetRole(s string) *RMLUserUpdate {
	ruu.mutation.SetRole(s)
	return ruu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ruu *RMLUserUpdate) SetNillableRole(s *string) *RMLUserUpdate {
	if s != nil {
		ruu.SetRole(*s)
	}
	return ruu
}

// SetReviewCount sets the "review_count" field.
func (ruu *RMLUserUpdate) SetReviewCount(i int) *RMLUserUpdate {
	ruu.mutation.ResetReviewCount()
	ruu.mutation.SetReviewCount(i)
	return ruu
}

// SetNillableReviewCount sets the "review_count" field if the given value is not nil.
func (ruu *RMLUserUpdate) SetNillableReviewCount(i *int) *RMLUserUpdate {
	if i != nil {
		ruu.SetReviewCount(*i)
	}
	return ruu
}

// AddReviewCount adds i to the "review_count" field.
func (ruu *RMLUserUpdate) AddReviewCount(i int) *RMLUserUpdate {
	ruu.mutation.AddReviewCount(i)
	return ruu
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ruu *RMLUserUpdate) AddReviewIDs(ids ...int) *RMLUserUpdate {
	ruu.mutation.AddReviewIDs(ids...)
	return ruu
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ruu *RMLUserUpdate) AddReviews(r ...*Review) *RMLUserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruu.AddReviewIDs(ids...)
}

// Mutation returns the RMLUserMutation object of the builder.
func (ruu *RMLUserUpdate) Mutation() *RMLUserMutation {
	return ruu.mutation
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (ruu *RMLUserUpdate) ClearReviews() *RMLUserUpdate {
	ruu.mutation.ClearReviews()
	return ruu
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (ruu *RMLUserUpdate) RemoveReviewIDs(ids ...int) *RMLUserUpdate {
	ruu.mutation.RemoveReviewIDs(ids...)
	return ruu
}

// RemoveReviews removes "reviews" edges to Review entities.
func (ruu *RMLUserUpdate) RemoveReviews(r ...*Review) *RMLUserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruu.RemoveReviewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ruu *RMLUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ruu.sqlSave, ruu.mutation, ruu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruu *RMLUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ruu *RMLUserUpdate) Exec(ctx context.Context) error {
	_, err := ruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruu *RMLUserUpdate) ExecX(ctx context.Context) {
	if err := ruu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruu *RMLUserUpdate) check() error {
	if v, ok := ruu.mutation.Name(); ok {
		if err := rmluser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RMLUser.name": %w`, err)}
		}
	}
	if v, ok := ruu.mutation.Email(); ok {
		if err := rmluser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "RMLUser.email": %w`, err)}
		}
	}
	if v, ok := ruu.mutation.Password(); ok {
		if err := rmluser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "RMLUser.password": %w`, err)}
		}
	}
	return nil
}

func (ruu *RMLUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ruu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(rmluser.Table, rmluser.Columns, sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt))
	if ps := ruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruu.mutation.Name(); ok {
		_spec.SetField(rmluser.FieldName, field.TypeString, value)
	}
	if value, ok := ruu.mutation.Email(); ok {
		_spec.SetField(rmluser.FieldEmail, field.TypeString, value)
	}
	if value, ok := ruu.mutation.Password(); ok {
		_spec.SetField(rmluser.FieldPassword, field.TypeString, value)
	}
	if value, ok := ruu.mutation.Role(); ok {
		_spec.SetField(rmluser.FieldRole, field.TypeString, value)
	}
	if value, ok := ruu.mutation.ReviewCount(); ok {
		_spec.SetField(rmluser.FieldReviewCount, field.TypeInt, value)
	}
	if value, ok := ruu.mutation.AddedReviewCount(); ok {
		_spec.AddField(rmluser.FieldReviewCount, field.TypeInt, value)
	}
	if ruu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rmluser.ReviewsTable,
			Columns: []string{rmluser.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruu.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ruu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rmluser.ReviewsTable,
			Columns: []string{rmluser.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruu.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rmluser.ReviewsTable,
			Columns: []string{rmluser.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rmluser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ruu.mutation.done = true
	return n, nil
}

// RMLUserUpdateOne is the builder for updating a single RMLUser entity.
type RMLUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RMLUserMutation
}

// SetName sets the "name" field.
func (ruuo *RMLUserUpdateOne) SetName(s string) *RMLUserUpdateOne {
	ruuo.mutation.SetName(s)
	return ruuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruuo *RMLUserUpdateOne) SetNillableName(s *string) *RMLUserUpdateOne {
	if s != nil {
		ruuo.SetName(*s)
	}
	return ruuo
}

// SetEmail sets the "email" field.
func (ruuo *RMLUserUpdateOne) SetEmail(s string) *RMLUserUpdateOne {
	ruuo.mutation.SetEmail(s)
	return ruuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ruuo *RMLUserUpdateOne) SetNillableEmail(s *string) *RMLUserUpdateOne {
	if s != nil {
		ruuo.SetEmail(*s)
	}
	return ruuo
}

// SetPassword sets the "password" field.
func (ruuo *RMLUserUpdateOne) SetPassword(s string) *RMLUserUpdateOne {
	ruuo.mutation.SetPassword(s)
	return ruuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ruuo *RMLUserUpdateOne) SetNillablePassword(s *string) *RMLUserUpdateOne {
	if s != nil {
		ruuo.SetPassword(*s)
	}
	return ruuo
}

// SetRole sets the "role" field.
func (ruuo *RMLUserUpdateOne) SetRole(s string) *RMLUserUpdateOne {
	ruuo.mutation.SetRole(s)
	return ruuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ruuo *RMLUserUpdateOne) SetNillableRole(s *string) *RMLUserUpdateOne {
	if s != nil {
		ruuo.SetRole(*s)
	}
	return ruuo
}

// SetReviewCount sets the "review_count" field.
func (ruuo *RMLUserUpdateOne) SetReviewCount(i int) *RMLUserUpdateOne {
	ruuo.mutation.ResetReviewCount()
	ruuo.mutation.SetReviewCount(i)
	return ruuo
}

// SetNillableReviewCount sets the "review_count" field if the given value is not nil.
func (ruuo *RMLUserUpdateOne) SetNillableReviewCount(i *int) *RMLUserUpdateOne {
	if i != nil {
		ruuo.SetReviewCount(*i)
	}
	return ruuo
}

// AddReviewCount adds i to the "review_count" field.
func (ruuo *RMLUserUpdateOne) AddReviewCount(i int) *RMLUserUpdateOne {
	ruuo.mutation.AddReviewCount(i)
	return ruuo
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ruuo *RMLUserUpdateOne) AddReviewIDs(ids ...int) *RMLUserUpdateOne {
	ruuo.mutation.AddReviewIDs(ids...)
	return ruuo
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ruuo *RMLUserUpdateOne) AddReviews(r ...*Review) *RMLUserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruuo.AddReviewIDs(ids...)
}

// Mutation returns the RMLUserMutation object of the builder.
func (ruuo *RMLUserUpdateOne) Mutation() *RMLUserMutation {
	return ruuo.mutation
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (ruuo *RMLUserUpdateOne) ClearReviews() *RMLUserUpdateOne {
	ruuo.mutation.ClearReviews()
	return ruuo
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (ruuo *RMLUserUpdateOne) RemoveReviewIDs(ids ...int) *RMLUserUpdateOne {
	ruuo.mutation.RemoveReviewIDs(ids...)
	return ruuo
}

// RemoveReviews removes "reviews" edges to Review entities.
func (ruuo *RMLUserUpdateOne) RemoveReviews(r ...*Review) *RMLUserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruuo.RemoveReviewIDs(ids...)
}

// Where appends a list predicates to the RMLUserUpdate builder.
func (ruuo *RMLUserUpdateOne) Where(ps ...predicate.RMLUser) *RMLUserUpdateOne {
	ruuo.mutation.Where(ps...)
	return ruuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruuo *RMLUserUpdateOne) Select(field string, fields ...string) *RMLUserUpdateOne {
	ruuo.fields = append([]string{field}, fields...)
	return ruuo
}

// Save executes the query and returns the updated RMLUser entity.
func (ruuo *RMLUserUpdateOne) Save(ctx context.Context) (*RMLUser, error) {
	return withHooks(ctx, ruuo.sqlSave, ruuo.mutation, ruuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruuo *RMLUserUpdateOne) SaveX(ctx context.Context) *RMLUser {
	node, err := ruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruuo *RMLUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruuo *RMLUserUpdateOne) ExecX(ctx context.Context) {
	if err := ruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruuo *RMLUserUpdateOne) check() error {
	if v, ok := ruuo.mutation.Name(); ok {
		if err := rmluser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RMLUser.name": %w`, err)}
		}
	}
	if v, ok := ruuo.mutation.Email(); ok {
		if err := rmluser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "RMLUser.email": %w`, err)}
		}
	}
	if v, ok := ruuo.mutation.Password(); ok {
		if err := rmluser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "RMLUser.password": %w`, err)}
		}
	}
	return nil
}

func (ruuo *RMLUserUpdateOne) sqlSave(ctx context.Context) (_node *RMLUser, err error) {
	if err := ruuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(rmluser.Table, rmluser.Columns, sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt))
	id, ok := ruuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RMLUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rmluser.FieldID)
		for _, f := range fields {
			if !rmluser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rmluser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruuo.mutation.Name(); ok {
		_spec.SetField(rmluser.FieldName, field.TypeString, value)
	}
	if value, ok := ruuo.mutation.Email(); ok {
		_spec.SetField(rmluser.FieldEmail, field.TypeString, value)
	}
	if value, ok := ruuo.mutation.Password(); ok {
		_spec.SetField(rmluser.FieldPassword, field.TypeString, value)
	}
	if value, ok := ruuo.mutation.Role(); ok {
		_spec.SetField(rmluser.FieldRole, field.TypeString, value)
	}
	if value, ok := ruuo.mutation.ReviewCount(); ok {
		_spec.SetField(rmluser.FieldReviewCount, field.TypeInt, value)
	}
	if value, ok := ruuo.mutation.AddedReviewCount(); ok {
		_spec.AddField(rmluser.FieldReviewCount, field.TypeInt, value)
	}
	if ruuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rmluser.ReviewsTable,
			Columns: []string{rmluser.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruuo.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ruuo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rmluser.ReviewsTable,
			Columns: []string{rmluser.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruuo.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   rmluser.ReviewsTable,
			Columns: []string{rmluser.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RMLUser{config: ruuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rmluser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruuo.mutation.done = true
	return _node, nil
}
