// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"parsing/orm/ent/page"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Page is the model entity for the Page schema.
type Page struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Brand holds the value of the "brand" field.
	Brand string `json:"brand,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// Job holds the value of the "job" field.
	Job string `json:"job,omitempty"`
	// HTML holds the value of the "html" field.
	HTML string `json:"html,omitempty" validate:"required"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Updated holds the value of the "updated" field.
	Updated time.Time `json:"updated,omitempty"`
	// Processed holds the value of the "processed" field.
	Processed time.Time `json:"processed,omitempty"`
	// URL holds the value of the "url" field.
	URL          string `json:"url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Page) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			values[i] = new(sql.NullInt64)
		case page.FieldBrand, page.FieldDomain, page.FieldJob, page.FieldHTML, page.FieldURL:
			values[i] = new(sql.NullString)
		case page.FieldCreated, page.FieldUpdated, page.FieldProcessed:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Page fields.
func (pa *Page) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case page.FieldBrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field brand", values[i])
			} else if value.Valid {
				pa.Brand = value.String
			}
		case page.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				pa.Domain = value.String
			}
		case page.FieldJob:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job", values[i])
			} else if value.Valid {
				pa.Job = value.String
			}
		case page.FieldHTML:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html", values[i])
			} else if value.Valid {
				pa.HTML = value.String
			}
		case page.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				pa.Created = value.Time
			}
		case page.FieldUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated", values[i])
			} else if value.Valid {
				pa.Updated = value.Time
			}
		case page.FieldProcessed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field processed", values[i])
			} else if value.Valid {
				pa.Processed = value.Time
			}
		case page.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pa.URL = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Page.
// This includes values selected through modifiers, order, etc.
func (pa *Page) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// Update returns a builder for updating this Page.
// Note that you need to call Page.Unwrap() before calling this method if this Page
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Page) Update() *PageUpdateOne {
	return NewPageClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Page entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Page) Unwrap() *Page {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Page is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Page) String() string {
	var builder strings.Builder
	builder.WriteString("Page(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("brand=")
	builder.WriteString(pa.Brand)
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(pa.Domain)
	builder.WriteString(", ")
	builder.WriteString("job=")
	builder.WriteString(pa.Job)
	builder.WriteString(", ")
	builder.WriteString("html=")
	builder.WriteString(pa.HTML)
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(pa.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated=")
	builder.WriteString(pa.Updated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("processed=")
	builder.WriteString(pa.Processed.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(pa.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Pages is a parsable slice of Page.
type Pages []*Page
