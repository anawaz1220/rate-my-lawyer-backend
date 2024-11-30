// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/review"
	"github.com/mikestefanello/pagoda/ent/rmluser"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetRating sets the "rating" field.
func (ru *ReviewUpdate) SetRating(i int) *ReviewUpdate {
	ru.mutation.ResetRating()
	ru.mutation.SetRating(i)
	return ru
}

// SetNillableRating sets the "rating" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableRating(i *int) *ReviewUpdate {
	if i != nil {
		ru.SetRating(*i)
	}
	return ru
}

// AddRating adds i to the "rating" field.
func (ru *ReviewUpdate) AddRating(i int) *ReviewUpdate {
	ru.mutation.AddRating(i)
	return ru
}

// SetComment sets the "comment" field.
func (ru *ReviewUpdate) SetComment(s string) *ReviewUpdate {
	ru.mutation.SetComment(s)
	return ru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableComment(s *string) *ReviewUpdate {
	if s != nil {
		ru.SetComment(*s)
	}
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ReviewUpdate) SetCreatedAt(t time.Time) *ReviewUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableCreatedAt(t *time.Time) *ReviewUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetLawyerReviews sets the "lawyer_reviews" field.
func (ru *ReviewUpdate) SetLawyerReviews(i int) *ReviewUpdate {
	ru.mutation.SetLawyerReviews(i)
	return ru
}

// SetNillableLawyerReviews sets the "lawyer_reviews" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableLawyerReviews(i *int) *ReviewUpdate {
	if i != nil {
		ru.SetLawyerReviews(*i)
	}
	return ru
}

// ClearLawyerReviews clears the value of the "lawyer_reviews" field.
func (ru *ReviewUpdate) ClearLawyerReviews() *ReviewUpdate {
	ru.mutation.ClearLawyerReviews()
	return ru
}

// SetRmlUserReviews sets the "rml_user_reviews" field.
func (ru *ReviewUpdate) SetRmlUserReviews(i int) *ReviewUpdate {
	ru.mutation.SetRmlUserReviews(i)
	return ru
}

// SetNillableRmlUserReviews sets the "rml_user_reviews" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableRmlUserReviews(i *int) *ReviewUpdate {
	if i != nil {
		ru.SetRmlUserReviews(*i)
	}
	return ru
}

// ClearRmlUserReviews clears the value of the "rml_user_reviews" field.
func (ru *ReviewUpdate) ClearRmlUserReviews() *ReviewUpdate {
	ru.mutation.ClearRmlUserReviews()
	return ru
}

// SetLawyerID sets the "lawyer" edge to the Lawyer entity by ID.
func (ru *ReviewUpdate) SetLawyerID(id int) *ReviewUpdate {
	ru.mutation.SetLawyerID(id)
	return ru
}

// SetNillableLawyerID sets the "lawyer" edge to the Lawyer entity by ID if the given value is not nil.
func (ru *ReviewUpdate) SetNillableLawyerID(id *int) *ReviewUpdate {
	if id != nil {
		ru = ru.SetLawyerID(*id)
	}
	return ru
}

// SetLawyer sets the "lawyer" edge to the Lawyer entity.
func (ru *ReviewUpdate) SetLawyer(l *Lawyer) *ReviewUpdate {
	return ru.SetLawyerID(l.ID)
}

// SetUserID sets the "user" edge to the RMLUser entity by ID.
func (ru *ReviewUpdate) SetUserID(id int) *ReviewUpdate {
	ru.mutation.SetUserID(id)
	return ru
}

// SetNillableUserID sets the "user" edge to the RMLUser entity by ID if the given value is not nil.
func (ru *ReviewUpdate) SetNillableUserID(id *int) *ReviewUpdate {
	if id != nil {
		ru = ru.SetUserID(*id)
	}
	return ru
}

// SetUser sets the "user" edge to the RMLUser entity.
func (ru *ReviewUpdate) SetUser(r *RMLUser) *ReviewUpdate {
	return ru.SetUserID(r.ID)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearLawyer clears the "lawyer" edge to the Lawyer entity.
func (ru *ReviewUpdate) ClearLawyer() *ReviewUpdate {
	ru.mutation.ClearLawyer()
	return ru
}

// ClearUser clears the "user" edge to the RMLUser entity.
func (ru *ReviewUpdate) ClearUser() *ReviewUpdate {
	ru.mutation.ClearUser()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewUpdate) check() error {
	if v, ok := ru.mutation.Rating(); ok {
		if err := review.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Review.rating": %w`, err)}
		}
	}
	return nil
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(review.Table, review.Columns, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Rating(); ok {
		_spec.SetField(review.FieldRating, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedRating(); ok {
		_spec.AddField(review.FieldRating, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Comment(); ok {
		_spec.SetField(review.FieldComment, field.TypeString, value)
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(review.FieldCreatedAt, field.TypeTime, value)
	}
	if ru.mutation.LawyerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.LawyerTable,
			Columns: []string{review.LawyerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LawyerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.LawyerTable,
			Columns: []string{review.LawyerColumn},
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
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetRating sets the "rating" field.
func (ruo *ReviewUpdateOne) SetRating(i int) *ReviewUpdateOne {
	ruo.mutation.ResetRating()
	ruo.mutation.SetRating(i)
	return ruo
}

// SetNillableRating sets the "rating" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableRating(i *int) *ReviewUpdateOne {
	if i != nil {
		ruo.SetRating(*i)
	}
	return ruo
}

// AddRating adds i to the "rating" field.
func (ruo *ReviewUpdateOne) AddRating(i int) *ReviewUpdateOne {
	ruo.mutation.AddRating(i)
	return ruo
}

// SetComment sets the "comment" field.
func (ruo *ReviewUpdateOne) SetComment(s string) *ReviewUpdateOne {
	ruo.mutation.SetComment(s)
	return ruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableComment(s *string) *ReviewUpdateOne {
	if s != nil {
		ruo.SetComment(*s)
	}
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ReviewUpdateOne) SetCreatedAt(t time.Time) *ReviewUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableCreatedAt(t *time.Time) *ReviewUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetLawyerReviews sets the "lawyer_reviews" field.
func (ruo *ReviewUpdateOne) SetLawyerReviews(i int) *ReviewUpdateOne {
	ruo.mutation.SetLawyerReviews(i)
	return ruo
}

// SetNillableLawyerReviews sets the "lawyer_reviews" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableLawyerReviews(i *int) *ReviewUpdateOne {
	if i != nil {
		ruo.SetLawyerReviews(*i)
	}
	return ruo
}

// ClearLawyerReviews clears the value of the "lawyer_reviews" field.
func (ruo *ReviewUpdateOne) ClearLawyerReviews() *ReviewUpdateOne {
	ruo.mutation.ClearLawyerReviews()
	return ruo
}

// SetRmlUserReviews sets the "rml_user_reviews" field.
func (ruo *ReviewUpdateOne) SetRmlUserReviews(i int) *ReviewUpdateOne {
	ruo.mutation.SetRmlUserReviews(i)
	return ruo
}

// SetNillableRmlUserReviews sets the "rml_user_reviews" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableRmlUserReviews(i *int) *ReviewUpdateOne {
	if i != nil {
		ruo.SetRmlUserReviews(*i)
	}
	return ruo
}

// ClearRmlUserReviews clears the value of the "rml_user_reviews" field.
func (ruo *ReviewUpdateOne) ClearRmlUserReviews() *ReviewUpdateOne {
	ruo.mutation.ClearRmlUserReviews()
	return ruo
}

// SetLawyerID sets the "lawyer" edge to the Lawyer entity by ID.
func (ruo *ReviewUpdateOne) SetLawyerID(id int) *ReviewUpdateOne {
	ruo.mutation.SetLawyerID(id)
	return ruo
}

// SetNillableLawyerID sets the "lawyer" edge to the Lawyer entity by ID if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableLawyerID(id *int) *ReviewUpdateOne {
	if id != nil {
		ruo = ruo.SetLawyerID(*id)
	}
	return ruo
}

// SetLawyer sets the "lawyer" edge to the Lawyer entity.
func (ruo *ReviewUpdateOne) SetLawyer(l *Lawyer) *ReviewUpdateOne {
	return ruo.SetLawyerID(l.ID)
}

// SetUserID sets the "user" edge to the RMLUser entity by ID.
func (ruo *ReviewUpdateOne) SetUserID(id int) *ReviewUpdateOne {
	ruo.mutation.SetUserID(id)
	return ruo
}

// SetNillableUserID sets the "user" edge to the RMLUser entity by ID if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableUserID(id *int) *ReviewUpdateOne {
	if id != nil {
		ruo = ruo.SetUserID(*id)
	}
	return ruo
}

// SetUser sets the "user" edge to the RMLUser entity.
func (ruo *ReviewUpdateOne) SetUser(r *RMLUser) *ReviewUpdateOne {
	return ruo.SetUserID(r.ID)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearLawyer clears the "lawyer" edge to the Lawyer entity.
func (ruo *ReviewUpdateOne) ClearLawyer() *ReviewUpdateOne {
	ruo.mutation.ClearLawyer()
	return ruo
}

// ClearUser clears the "user" edge to the RMLUser entity.
func (ruo *ReviewUpdateOne) ClearUser() *ReviewUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ruo *ReviewUpdateOne) Where(ps ...predicate.Review) *ReviewUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewUpdateOne) check() error {
	if v, ok := ruo.mutation.Rating(); ok {
		if err := review.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Review.rating": %w`, err)}
		}
	}
	return nil
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(review.Table, review.Columns, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Review.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Rating(); ok {
		_spec.SetField(review.FieldRating, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedRating(); ok {
		_spec.AddField(review.FieldRating, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Comment(); ok {
		_spec.SetField(review.FieldComment, field.TypeString, value)
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(review.FieldCreatedAt, field.TypeTime, value)
	}
	if ruo.mutation.LawyerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.LawyerTable,
			Columns: []string{review.LawyerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LawyerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.LawyerTable,
			Columns: []string{review.LawyerColumn},
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
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
