// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GavukaAlexandr/ent-sandbox/ent/pet"
	"github.com/GavukaAlexandr/ent-sandbox/ent/user"
	"github.com/google/uuid"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAge sets the "age" field.
func (pc *PetCreate) SetAge(i int) *PetCreate {
	pc.mutation.SetAge(i)
	return pc
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetSkills sets the "skills" field.
func (pc *PetCreate) SetSkills(s []string) *PetCreate {
	pc.mutation.SetSkills(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PetCreate) SetID(u uuid.UUID) *PetCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PetCreate) SetOwnerID(id uuid.UUID) *PetCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(id *uuid.UUID) *PetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	var (
		err  error
		node *Pet
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PetCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PetCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PetCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := pet.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Pet.age"`)}
	}
	if v, ok := pc.mutation.Age(); ok {
		if err := pet.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "Pet.age": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	return nil
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pet.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pet.FieldAge,
		})
		_node.Age = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pet.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Skills(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: pet.FieldSkills,
		})
		_node.Skills = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pet.Create().
//		SetAge(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PetUpsert) {
//			SetAge(v+v).
//		}).
//		Exec(ctx)
//
func (pc *PetCreate) OnConflict(opts ...sql.ConflictOption) *PetUpsertOne {
	pc.conflict = opts
	return &PetUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pc *PetCreate) OnConflictColumns(columns ...string) *PetUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PetUpsertOne{
		create: pc,
	}
}

type (
	// PetUpsertOne is the builder for "upsert"-ing
	//  one Pet node.
	PetUpsertOne struct {
		create *PetCreate
	}

	// PetUpsert is the "OnConflict" setter.
	PetUpsert struct {
		*sql.UpdateSet
	}
)

// SetAge sets the "age" field.
func (u *PetUpsert) SetAge(v int) *PetUpsert {
	u.Set(pet.FieldAge, v)
	return u
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *PetUpsert) UpdateAge() *PetUpsert {
	u.SetExcluded(pet.FieldAge)
	return u
}

// AddAge adds v to the "age" field.
func (u *PetUpsert) AddAge(v int) *PetUpsert {
	u.Add(pet.FieldAge, v)
	return u
}

// SetName sets the "name" field.
func (u *PetUpsert) SetName(v string) *PetUpsert {
	u.Set(pet.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PetUpsert) UpdateName() *PetUpsert {
	u.SetExcluded(pet.FieldName)
	return u
}

// SetSkills sets the "skills" field.
func (u *PetUpsert) SetSkills(v []string) *PetUpsert {
	u.Set(pet.FieldSkills, v)
	return u
}

// UpdateSkills sets the "skills" field to the value that was provided on create.
func (u *PetUpsert) UpdateSkills() *PetUpsert {
	u.SetExcluded(pet.FieldSkills)
	return u
}

// ClearSkills clears the value of the "skills" field.
func (u *PetUpsert) ClearSkills() *PetUpsert {
	u.SetNull(pet.FieldSkills)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(pet.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PetUpsertOne) UpdateNewValues() *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(pet.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Pet.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PetUpsertOne) Ignore() *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PetUpsertOne) DoNothing() *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PetCreate.OnConflict
// documentation for more info.
func (u *PetUpsertOne) Update(set func(*PetUpsert)) *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PetUpsert{UpdateSet: update})
	}))
	return u
}

// SetAge sets the "age" field.
func (u *PetUpsertOne) SetAge(v int) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *PetUpsertOne) AddAge(v int) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateAge() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateAge()
	})
}

// SetName sets the "name" field.
func (u *PetUpsertOne) SetName(v string) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateName() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateName()
	})
}

// SetSkills sets the "skills" field.
func (u *PetUpsertOne) SetSkills(v []string) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetSkills(v)
	})
}

// UpdateSkills sets the "skills" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateSkills() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateSkills()
	})
}

// ClearSkills clears the value of the "skills" field.
func (u *PetUpsertOne) ClearSkills() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.ClearSkills()
	})
}

// Exec executes the query.
func (u *PetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PetUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PetUpsertOne.ID is not supported by MySQL driver. Use PetUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PetUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	builders []*PetCreate
	conflict []sql.ConflictOption
}

// Save creates the Pet entities in the database.
func (pcb *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PetCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pet.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PetUpsert) {
//			SetAge(v+v).
//		}).
//		Exec(ctx)
//
func (pcb *PetCreateBulk) OnConflict(opts ...sql.ConflictOption) *PetUpsertBulk {
	pcb.conflict = opts
	return &PetUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pcb *PetCreateBulk) OnConflictColumns(columns ...string) *PetUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PetUpsertBulk{
		create: pcb,
	}
}

// PetUpsertBulk is the builder for "upsert"-ing
// a bulk of Pet nodes.
type PetUpsertBulk struct {
	create *PetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(pet.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PetUpsertBulk) UpdateNewValues() *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(pet.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PetUpsertBulk) Ignore() *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PetUpsertBulk) DoNothing() *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PetCreateBulk.OnConflict
// documentation for more info.
func (u *PetUpsertBulk) Update(set func(*PetUpsert)) *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PetUpsert{UpdateSet: update})
	}))
	return u
}

// SetAge sets the "age" field.
func (u *PetUpsertBulk) SetAge(v int) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *PetUpsertBulk) AddAge(v int) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateAge() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateAge()
	})
}

// SetName sets the "name" field.
func (u *PetUpsertBulk) SetName(v string) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateName() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateName()
	})
}

// SetSkills sets the "skills" field.
func (u *PetUpsertBulk) SetSkills(v []string) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetSkills(v)
	})
}

// UpdateSkills sets the "skills" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateSkills() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateSkills()
	})
}

// ClearSkills clears the value of the "skills" field.
func (u *PetUpsertBulk) ClearSkills() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.ClearSkills()
	})
}

// Exec executes the query.
func (u *PetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
