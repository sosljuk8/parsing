package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("brand").NotEmpty(),
		field.String("name").NotEmpty(),
		field.Time("created").
			Default(time.Now),
		field.Time("updated").
			Default(time.Now),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("parents", Category.Type).
			From("children"),
	}
}
