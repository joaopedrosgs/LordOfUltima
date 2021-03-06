// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/joaopedrosgs/openlou/ent/city"
	"github.com/joaopedrosgs/openlou/ent/construction"
)

// Construction is the model entity for the Construction schema.
type Construction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// X holds the value of the "x" field.
	X int `json:"x,omitempty"`
	// Y holds the value of the "y" field.
	Y int `json:"y,omitempty"`
	// RawProduction holds the value of the "raw_production" field.
	RawProduction int `json:"raw_production,omitempty"`
	// Type holds the value of the "type" field.
	Type int `json:"type,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// Modifier holds the value of the "modifier" field.
	Modifier float64 `json:"modifier,omitempty"`
	// NeedRefresh holds the value of the "need_refresh" field.
	NeedRefresh bool `json:"need_refresh,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConstructionQuery when eager-loading is set.
	Edges              ConstructionEdges `json:"edges"`
	city_constructions *int
}

// ConstructionEdges holds the relations/edges for other nodes in the graph.
type ConstructionEdges struct {
	// City holds the value of the city edge.
	City *City
	// Queue holds the value of the queue edge.
	Queue []*Queue
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CityOrErr returns the City value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConstructionEdges) CityOrErr() (*City, error) {
	if e.loadedTypes[0] {
		if e.City == nil {
			// The edge city was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: city.Label}
		}
		return e.City, nil
	}
	return nil, &NotLoadedError{edge: "city"}
}

// QueueOrErr returns the Queue value or an error if the edge
// was not loaded in eager-loading.
func (e ConstructionEdges) QueueOrErr() ([]*Queue, error) {
	if e.loadedTypes[1] {
		return e.Queue, nil
	}
	return nil, &NotLoadedError{edge: "queue"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Construction) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullInt64{},   // x
		&sql.NullInt64{},   // y
		&sql.NullInt64{},   // raw_production
		&sql.NullInt64{},   // type
		&sql.NullInt64{},   // level
		&sql.NullFloat64{}, // modifier
		&sql.NullBool{},    // need_refresh
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Construction) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // city_constructions
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Construction fields.
func (c *Construction) assignValues(values ...interface{}) error {
	if m, n := len(values), len(construction.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field x", values[0])
	} else if value.Valid {
		c.X = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field y", values[1])
	} else if value.Valid {
		c.Y = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field raw_production", values[2])
	} else if value.Valid {
		c.RawProduction = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[3])
	} else if value.Valid {
		c.Type = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field level", values[4])
	} else if value.Valid {
		c.Level = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field modifier", values[5])
	} else if value.Valid {
		c.Modifier = value.Float64
	}
	if value, ok := values[6].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field need_refresh", values[6])
	} else if value.Valid {
		c.NeedRefresh = value.Bool
	}
	values = values[7:]
	if len(values) == len(construction.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field city_constructions", value)
		} else if value.Valid {
			c.city_constructions = new(int)
			*c.city_constructions = int(value.Int64)
		}
	}
	return nil
}

// QueryCity queries the city edge of the Construction.
func (c *Construction) QueryCity() *CityQuery {
	return (&ConstructionClient{config: c.config}).QueryCity(c)
}

// QueryQueue queries the queue edge of the Construction.
func (c *Construction) QueryQueue() *QueueQuery {
	return (&ConstructionClient{config: c.config}).QueryQueue(c)
}

// Update returns a builder for updating this Construction.
// Note that, you need to call Construction.Unwrap() before calling this method, if this Construction
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Construction) Update() *ConstructionUpdateOne {
	return (&ConstructionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Construction) Unwrap() *Construction {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Construction is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Construction) String() string {
	var builder strings.Builder
	builder.WriteString("Construction(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", x=")
	builder.WriteString(fmt.Sprintf("%v", c.X))
	builder.WriteString(", y=")
	builder.WriteString(fmt.Sprintf("%v", c.Y))
	builder.WriteString(", raw_production=")
	builder.WriteString(fmt.Sprintf("%v", c.RawProduction))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", level=")
	builder.WriteString(fmt.Sprintf("%v", c.Level))
	builder.WriteString(", modifier=")
	builder.WriteString(fmt.Sprintf("%v", c.Modifier))
	builder.WriteString(", need_refresh=")
	builder.WriteString(fmt.Sprintf("%v", c.NeedRefresh))
	builder.WriteByte(')')
	return builder.String()
}

// Constructions is a parsable slice of Construction.
type Constructions []*Construction

func (c Constructions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
