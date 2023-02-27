// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/netzchat/server/ent/channel"
)

// ChannelCreate is the builder for creating a Channel entity.
type ChannelCreate struct {
	config
	mutation *ChannelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ChannelCreate) SetName(s string) *ChannelCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetID sets the "id" field.
func (cc *ChannelCreate) SetID(s string) *ChannelCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableID(s *string) *ChannelCreate {
	if s != nil {
		cc.SetID(*s)
	}
	return cc
}

// Mutation returns the ChannelMutation object of the builder.
func (cc *ChannelCreate) Mutation() *ChannelMutation {
	return cc.mutation
}

// Save creates the Channel in the database.
func (cc *ChannelCreate) Save(ctx context.Context) (*Channel, error) {
	cc.defaults()
	return withHooks[*Channel, ChannelMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChannelCreate) SaveX(ctx context.Context) *Channel {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChannelCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChannelCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChannelCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := channel.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChannelCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Channel.name"`)}
	}
	return nil
}

func (cc *ChannelCreate) sqlSave(ctx context.Context) (*Channel, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Channel.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChannelCreate) createSpec() (*Channel, *sqlgraph.CreateSpec) {
	var (
		_node = &Channel{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(channel.Table, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeString))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// ChannelCreateBulk is the builder for creating many Channel entities in bulk.
type ChannelCreateBulk struct {
	config
	builders []*ChannelCreate
}

// Save creates the Channel entities in the database.
func (ccb *ChannelCreateBulk) Save(ctx context.Context) ([]*Channel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Channel, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChannelCreateBulk) SaveX(ctx context.Context) []*Channel {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChannelCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
