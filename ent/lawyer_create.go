// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/decision"
	"github.com/mikestefanello/pagoda/ent/jurisdiction"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/lawyerjurisdiction"
	"github.com/mikestefanello/pagoda/ent/review"
)

// LawyerCreate is the builder for creating a Lawyer entity.
type LawyerCreate struct {
	config
	mutation *LawyerMutation
	hooks    []Hook
}

// SetFullName sets the "full_name" field.
func (lc *LawyerCreate) SetFullName(s string) *LawyerCreate {
	lc.mutation.SetFullName(s)
	return lc
}

// SetFirstName sets the "first_name" field.
func (lc *LawyerCreate) SetFirstName(s string) *LawyerCreate {
	lc.mutation.SetFirstName(s)
	return lc
}

// SetMiddleName sets the "middle_name" field.
func (lc *LawyerCreate) SetMiddleName(s string) *LawyerCreate {
	lc.mutation.SetMiddleName(s)
	return lc
}

// SetLastName sets the "last_name" field.
func (lc *LawyerCreate) SetLastName(s string) *LawyerCreate {
	lc.mutation.SetLastName(s)
	return lc
}

// SetGender sets the "gender" field.
func (lc *LawyerCreate) SetGender(s string) *LawyerCreate {
	lc.mutation.SetGender(s)
	return lc
}

// SetAddress sets the "address" field.
func (lc *LawyerCreate) SetAddress(s string) *LawyerCreate {
	lc.mutation.SetAddress(s)
	return lc
}

// SetPhone sets the "phone" field.
func (lc *LawyerCreate) SetPhone(s string) *LawyerCreate {
	lc.mutation.SetPhone(s)
	return lc
}

// SetPracticingStatus sets the "practicing_status" field.
func (lc *LawyerCreate) SetPracticingStatus(s string) *LawyerCreate {
	lc.mutation.SetPracticingStatus(s)
	return lc
}

// SetProfileURL sets the "profile_url" field.
func (lc *LawyerCreate) SetProfileURL(s string) *LawyerCreate {
	lc.mutation.SetProfileURL(s)
	return lc
}

// SetAverageRating sets the "average_rating" field.
func (lc *LawyerCreate) SetAverageRating(f float64) *LawyerCreate {
	lc.mutation.SetAverageRating(f)
	return lc
}

// SetNillableAverageRating sets the "average_rating" field if the given value is not nil.
func (lc *LawyerCreate) SetNillableAverageRating(f *float64) *LawyerCreate {
	if f != nil {
		lc.SetAverageRating(*f)
	}
	return lc
}

// SetReviewCount sets the "review_count" field.
func (lc *LawyerCreate) SetReviewCount(i int) *LawyerCreate {
	lc.mutation.SetReviewCount(i)
	return lc
}

// SetNillableReviewCount sets the "review_count" field if the given value is not nil.
func (lc *LawyerCreate) SetNillableReviewCount(i *int) *LawyerCreate {
	if i != nil {
		lc.SetReviewCount(*i)
	}
	return lc
}

// AddJurisdictionIDs adds the "jurisdictions" edge to the Jurisdiction entity by IDs.
func (lc *LawyerCreate) AddJurisdictionIDs(ids ...int) *LawyerCreate {
	lc.mutation.AddJurisdictionIDs(ids...)
	return lc
}

// AddJurisdictions adds the "jurisdictions" edges to the Jurisdiction entity.
func (lc *LawyerCreate) AddJurisdictions(j ...*Jurisdiction) *LawyerCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return lc.AddJurisdictionIDs(ids...)
}

// AddDecisionIDs adds the "decisions" edge to the Decision entity by IDs.
func (lc *LawyerCreate) AddDecisionIDs(ids ...int) *LawyerCreate {
	lc.mutation.AddDecisionIDs(ids...)
	return lc
}

// AddDecisions adds the "decisions" edges to the Decision entity.
func (lc *LawyerCreate) AddDecisions(d ...*Decision) *LawyerCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return lc.AddDecisionIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (lc *LawyerCreate) AddReviewIDs(ids ...int) *LawyerCreate {
	lc.mutation.AddReviewIDs(ids...)
	return lc
}

// AddReviews adds the "reviews" edges to the Review entity.
func (lc *LawyerCreate) AddReviews(r ...*Review) *LawyerCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return lc.AddReviewIDs(ids...)
}

// AddLawyerJurisdictionIDs adds the "lawyer_jurisdictions" edge to the LawyerJurisdiction entity by IDs.
func (lc *LawyerCreate) AddLawyerJurisdictionIDs(ids ...int) *LawyerCreate {
	lc.mutation.AddLawyerJurisdictionIDs(ids...)
	return lc
}

// AddLawyerJurisdictions adds the "lawyer_jurisdictions" edges to the LawyerJurisdiction entity.
func (lc *LawyerCreate) AddLawyerJurisdictions(l ...*LawyerJurisdiction) *LawyerCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lc.AddLawyerJurisdictionIDs(ids...)
}

// Mutation returns the LawyerMutation object of the builder.
func (lc *LawyerCreate) Mutation() *LawyerMutation {
	return lc.mutation
}

// Save creates the Lawyer in the database.
func (lc *LawyerCreate) Save(ctx context.Context) (*Lawyer, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LawyerCreate) SaveX(ctx context.Context) *Lawyer {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LawyerCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LawyerCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LawyerCreate) defaults() {
	if _, ok := lc.mutation.AverageRating(); !ok {
		v := lawyer.DefaultAverageRating
		lc.mutation.SetAverageRating(v)
	}
	if _, ok := lc.mutation.ReviewCount(); !ok {
		v := lawyer.DefaultReviewCount
		lc.mutation.SetReviewCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LawyerCreate) check() error {
	if _, ok := lc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "Lawyer.full_name"`)}
	}
	if v, ok := lc.mutation.FullName(); ok {
		if err := lawyer.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Lawyer.full_name": %w`, err)}
		}
	}
	if _, ok := lc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Lawyer.first_name"`)}
	}
	if _, ok := lc.mutation.MiddleName(); !ok {
		return &ValidationError{Name: "middle_name", err: errors.New(`ent: missing required field "Lawyer.middle_name"`)}
	}
	if _, ok := lc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Lawyer.last_name"`)}
	}
	if _, ok := lc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Lawyer.gender"`)}
	}
	if _, ok := lc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Lawyer.address"`)}
	}
	if _, ok := lc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Lawyer.phone"`)}
	}
	if _, ok := lc.mutation.PracticingStatus(); !ok {
		return &ValidationError{Name: "practicing_status", err: errors.New(`ent: missing required field "Lawyer.practicing_status"`)}
	}
	if _, ok := lc.mutation.ProfileURL(); !ok {
		return &ValidationError{Name: "profile_url", err: errors.New(`ent: missing required field "Lawyer.profile_url"`)}
	}
	if _, ok := lc.mutation.AverageRating(); !ok {
		return &ValidationError{Name: "average_rating", err: errors.New(`ent: missing required field "Lawyer.average_rating"`)}
	}
	if _, ok := lc.mutation.ReviewCount(); !ok {
		return &ValidationError{Name: "review_count", err: errors.New(`ent: missing required field "Lawyer.review_count"`)}
	}
	return nil
}

func (lc *LawyerCreate) sqlSave(ctx context.Context) (*Lawyer, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LawyerCreate) createSpec() (*Lawyer, *sqlgraph.CreateSpec) {
	var (
		_node = &Lawyer{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(lawyer.Table, sqlgraph.NewFieldSpec(lawyer.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.FullName(); ok {
		_spec.SetField(lawyer.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := lc.mutation.FirstName(); ok {
		_spec.SetField(lawyer.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := lc.mutation.MiddleName(); ok {
		_spec.SetField(lawyer.FieldMiddleName, field.TypeString, value)
		_node.MiddleName = value
	}
	if value, ok := lc.mutation.LastName(); ok {
		_spec.SetField(lawyer.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := lc.mutation.Gender(); ok {
		_spec.SetField(lawyer.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	if value, ok := lc.mutation.Address(); ok {
		_spec.SetField(lawyer.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := lc.mutation.Phone(); ok {
		_spec.SetField(lawyer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := lc.mutation.PracticingStatus(); ok {
		_spec.SetField(lawyer.FieldPracticingStatus, field.TypeString, value)
		_node.PracticingStatus = value
	}
	if value, ok := lc.mutation.ProfileURL(); ok {
		_spec.SetField(lawyer.FieldProfileURL, field.TypeString, value)
		_node.ProfileURL = value
	}
	if value, ok := lc.mutation.AverageRating(); ok {
		_spec.SetField(lawyer.FieldAverageRating, field.TypeFloat64, value)
		_node.AverageRating = value
	}
	if value, ok := lc.mutation.ReviewCount(); ok {
		_spec.SetField(lawyer.FieldReviewCount, field.TypeInt, value)
		_node.ReviewCount = value
	}
	if nodes := lc.mutation.JurisdictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lawyer.JurisdictionsTable,
			Columns: lawyer.JurisdictionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.DecisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lawyer.DecisionsTable,
			Columns: []string{lawyer.DecisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decision.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lawyer.ReviewsTable,
			Columns: []string{lawyer.ReviewsColumn},
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
	if nodes := lc.mutation.LawyerJurisdictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   lawyer.LawyerJurisdictionsTable,
			Columns: []string{lawyer.LawyerJurisdictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lawyerjurisdiction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LawyerCreateBulk is the builder for creating many Lawyer entities in bulk.
type LawyerCreateBulk struct {
	config
	err      error
	builders []*LawyerCreate
}

// Save creates the Lawyer entities in the database.
func (lcb *LawyerCreateBulk) Save(ctx context.Context) ([]*Lawyer, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lawyer, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LawyerMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LawyerCreateBulk) SaveX(ctx context.Context) []*Lawyer {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LawyerCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LawyerCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}