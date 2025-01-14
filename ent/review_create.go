// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/review"
	"github.com/mikestefanello/pagoda/ent/rmluser"
)

// ReviewCreate is the builder for creating a Review entity.
type ReviewCreate struct {
	config
	mutation *ReviewMutation
	hooks    []Hook
}

// SetRating sets the "rating" field.
func (rc *ReviewCreate) SetRating(i int) *ReviewCreate {
	rc.mutation.SetRating(i)
	return rc
}

// SetComment sets the "comment" field.
func (rc *ReviewCreate) SetComment(s string) *ReviewCreate {
	rc.mutation.SetComment(s)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReviewCreate) SetCreatedAt(t time.Time) *ReviewCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableCreatedAt(t *time.Time) *ReviewCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetLawyerReviews sets the "lawyer_reviews" field.
func (rc *ReviewCreate) SetLawyerReviews(i int) *ReviewCreate {
	rc.mutation.SetLawyerReviews(i)
	return rc
}

// SetNillableLawyerReviews sets the "lawyer_reviews" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableLawyerReviews(i *int) *ReviewCreate {
	if i != nil {
		rc.SetLawyerReviews(*i)
	}
	return rc
}

// SetRmlUserReviews sets the "rml_user_reviews" field.
func (rc *ReviewCreate) SetRmlUserReviews(i int) *ReviewCreate {
	rc.mutation.SetRmlUserReviews(i)
	return rc
}

// SetNillableRmlUserReviews sets the "rml_user_reviews" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableRmlUserReviews(i *int) *ReviewCreate {
	if i != nil {
		rc.SetRmlUserReviews(*i)
	}
	return rc
}

// SetLawyerID sets the "lawyer" edge to the Lawyer entity by ID.
func (rc *ReviewCreate) SetLawyerID(id int) *ReviewCreate {
	rc.mutation.SetLawyerID(id)
	return rc
}

// SetNillableLawyerID sets the "lawyer" edge to the Lawyer entity by ID if the given value is not nil.
func (rc *ReviewCreate) SetNillableLawyerID(id *int) *ReviewCreate {
	if id != nil {
		rc = rc.SetLawyerID(*id)
	}
	return rc
}

// SetLawyer sets the "lawyer" edge to the Lawyer entity.
func (rc *ReviewCreate) SetLawyer(l *Lawyer) *ReviewCreate {
	return rc.SetLawyerID(l.ID)
}

// SetUserID sets the "user" edge to the RMLUser entity by ID.
func (rc *ReviewCreate) SetUserID(id int) *ReviewCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetNillableUserID sets the "user" edge to the RMLUser entity by ID if the given value is not nil.
func (rc *ReviewCreate) SetNillableUserID(id *int) *ReviewCreate {
	if id != nil {
		rc = rc.SetUserID(*id)
	}
	return rc
}

// SetUser sets the "user" edge to the RMLUser entity.
func (rc *ReviewCreate) SetUser(r *RMLUser) *ReviewCreate {
	return rc.SetUserID(r.ID)
}

// Mutation returns the ReviewMutation object of the builder.
func (rc *ReviewCreate) Mutation() *ReviewMutation {
	return rc.mutation
}

// Save creates the Review in the database.
func (rc *ReviewCreate) Save(ctx context.Context) (*Review, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReviewCreate) SaveX(ctx context.Context) *Review {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReviewCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReviewCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReviewCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := review.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReviewCreate) check() error {
	if _, ok := rc.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New(`ent: missing required field "Review.rating"`)}
	}
	if v, ok := rc.mutation.Rating(); ok {
		if err := review.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Review.rating": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Review.comment"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Review.created_at"`)}
	}
	return nil
}

func (rc *ReviewCreate) sqlSave(ctx context.Context) (*Review, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReviewCreate) createSpec() (*Review, *sqlgraph.CreateSpec) {
	var (
		_node = &Review{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(review.Table, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Rating(); ok {
		_spec.SetField(review.FieldRating, field.TypeInt, value)
		_node.Rating = value
	}
	if value, ok := rc.mutation.Comment(); ok {
		_spec.SetField(review.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(review.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := rc.mutation.LawyerIDs(); len(nodes) > 0 {
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
		_node.LawyerReviews = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.RmlUserReviews = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReviewCreateBulk is the builder for creating many Review entities in bulk.
type ReviewCreateBulk struct {
	config
	err      error
	builders []*ReviewCreate
}

// Save creates the Review entities in the database.
func (rcb *ReviewCreateBulk) Save(ctx context.Context) ([]*Review, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Review, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReviewMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReviewCreateBulk) SaveX(ctx context.Context) []*Review {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReviewCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReviewCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
