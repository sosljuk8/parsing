package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

const ChromeUrlMaxLen = 2048

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("brand").NotEmpty(),
		field.String("domain").NotEmpty(),
		field.Text("html").NotEmpty().StructTag(`validate:"required"`),
		field.Time("created").
			Default(time.Now),
		field.Time("updated").
			Default(time.Now),
		field.String("url").Unique().NotEmpty().MaxLen(ChromeUrlMaxLen),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
