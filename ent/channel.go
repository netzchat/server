// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/netzchat/server/ent/channel"
)

// Channel is the model entity for the Channel schema.
type Channel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Channel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case channel.FieldID:
			values[i] = new(sql.NullInt64)
		case channel.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Channel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Channel fields.
func (c *Channel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case channel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case channel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Channel.
// Note that you need to call Channel.Unwrap() before calling this method if this Channel
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Channel) Update() *ChannelUpdateOne {
	return NewChannelClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Channel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Channel) Unwrap() *Channel {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Channel is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Channel) String() string {
	var builder strings.Builder
	builder.WriteString("Channel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Channels is a parsable slice of Channel.
type Channels []*Channel
