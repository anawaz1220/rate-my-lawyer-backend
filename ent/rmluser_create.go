// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/review"
	"github.com/mikestefanello/pagoda/ent/rmluser"
)

// RMLUserCreate is the builder for creating a RMLUser entity.
type RMLUserCreate struct {
	config
	mutation *RMLUserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ruc *RMLUserCreate) SetName(s string) *RMLUserCreate {
	ruc.mutation.SetName(s)
	return ruc
}

// SetEmail sets the "email" field.
func (ruc *RMLUserCreate) SetEmail(s string) *RMLUserCreate {
	ruc.mutation.SetEmail(s)
	return ruc
}

// SetPassword sets the "password" field.
func (ruc *RMLUserCreate) SetPassword(s string) *RMLUserCreate {
	ruc.mutation.SetPassword(s)
	return ruc
}

// SetRole sets the "role" field.
func (ruc *RMLUserCreate) SetRole(s string) *RMLUserCreate {
	ruc.mutation.SetRole(s)
	return ruc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ruc *RMLUserCreate) SetNillableRole(s *string) *RMLUserCreate {
	if s != nil {
		ruc.SetRole(*s)
	}
	return ruc
}

// SetReviewCount sets the "review_count" field.
func (ruc *RMLUserCreate) SetReviewCount(i int) *RMLUserCreate {
	ruc.mutation.SetReviewCount(i)
	return ruc
}

// SetNillableReviewCount sets the "review_count" field if the given value is not nil.
func (ruc *RMLUserCreate) SetNillableReviewCount(i *int) *RMLUserCreate {
	if i != nil {
		ruc.SetReviewCount(*i)
	}
	return ruc
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (ruc *RMLUserCreate) AddReviewIDs(ids ...int) *RMLUserCreate {
	ruc.mutation.AddReviewIDs(ids...)
	return ruc
}

// AddReviews adds the "reviews" edges to the Review entity.
func (ruc *RMLUserCreate) AddReviews(r ...*Review) *RMLUserCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruc.AddReviewIDs(ids...)
}

// Mutation returns the RMLUserMutation object of the builder.
func (ruc *RMLUserCreate) Mutation() *RMLUserMutation {
	return ruc.mutation
}

// Save creates the RMLUser in the database.
func (ruc *RMLUserCreate) Save(ctx context.Context) (*RMLUser, error) {
	ruc.defaults()
	return withHooks(ctx, ruc.sqlSave, ruc.mutation, ruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ruc *RMLUserCreate) SaveX(ctx context.Context) *RMLUser {
	v, err := ruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ruc *RMLUserCreate) Exec(ctx context.Context) error {
	_, err := ruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruc *RMLUserCreate) ExecX(ctx context.Context) {
	if err := ruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruc *RMLUserCreate) defaults() {
	if _, ok := ruc.mutation.Role(); !ok {
		v := rmluser.DefaultRole
		ruc.mutation.SetRole(v)
	}
	if _, ok := ruc.mutation.ReviewCount(); !ok {
		v := rmluser.DefaultReviewCount
		ruc.mutation.SetReviewCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruc *RMLUserCreate) check() error {
	if _, ok := ruc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "RMLUser.name"`)}
	}
	if v, ok := ruc.mutation.Name(); ok {
		if err := rmluser.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RMLUser.name": %w`, err)}
		}
	}
	if _, ok := ruc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "RMLUser.email"`)}
	}
	if v, ok := ruc.mutation.Email(); ok {
		if err := rmluser.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "RMLUser.email": %w`, err)}
		}
	}
	if _, ok := ruc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "RMLUser.password"`)}
	}
	if v, ok := ruc.mutation.Password(); ok {
		if err := rmluser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "RMLUser.password": %w`, err)}
		}
	}
	if _, ok := ruc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "RMLUser.role"`)}
	}
	if _, ok := ruc.mutation.ReviewCount(); !ok {
		return &ValidationError{Name: "review_count", err: errors.New(`ent: missing required field "RMLUser.review_count"`)}
	}
	return nil
}

func (ruc *RMLUserCreate) sqlSave(ctx context.Context) (*RMLUser, error) {
	if err := ruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ruc.mutation.id = &_node.ID
	ruc.mutation.done = true
	return _node, nil
}

func (ruc *RMLUserCreate) createSpec() (*RMLUser, *sqlgraph.CreateSpec) {
	var (
		_node = &RMLUser{config: ruc.config}
		_spec = sqlgraph.NewCreateSpec(rmluser.Table, sqlgraph.NewFieldSpec(rmluser.FieldID, field.TypeInt))
	)
	if value, ok := ruc.mutation.Name(); ok {
		_spec.SetField(rmluser.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ruc.mutation.Email(); ok {
		_spec.SetField(rmluser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ruc.mutation.Password(); ok {
		_spec.SetField(rmluser.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ruc.mutation.Role(); ok {
		_spec.SetField(rmluser.FieldRole, field.TypeString, value)
		_node.Role = value
	}
	if value, ok := ruc.mutation.ReviewCount(); ok {
		_spec.SetField(rmluser.FieldReviewCount, field.TypeInt, value)
		_node.ReviewCount = value
	}
	if nodes := ruc.mutation.ReviewsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RMLUserCreateBulk is the builder for creating many RMLUser entities in bulk.
type RMLUserCreateBulk struct {
	config
	err      error
	builders []*RMLUserCreate
}

// Save creates the RMLUser entities in the database.
func (rucb *RMLUserCreateBulk) Save(ctx context.Context) ([]*RMLUser, error) {
	if rucb.err != nil {
		return nil, rucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rucb.builders))
	nodes := make([]*RMLUser, len(rucb.builders))
	mutators := make([]Mutator, len(rucb.builders))
	for i := range rucb.builders {
		func(i int, root context.Context) {
			builder := rucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RMLUserMutation)
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
					_, err = mutators[i+1].Mutate(root, rucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rucb *RMLUserCreateBulk) SaveX(ctx context.Context) []*RMLUser {
	v, err := rucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rucb *RMLUserCreateBulk) Exec(ctx context.Context) error {
	_, err := rucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rucb *RMLUserCreateBulk) ExecX(ctx context.Context) {
	if err := rucb.Exec(ctx); err != nil {
		panic(err)
	}
}