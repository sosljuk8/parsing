package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("series").
			Optional(),
		field.String("model").
			Optional(),
		field.String("sku").
			Optional(),
		field.JSON("properties", map[string]string{}).
			Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("products"),
		edge.To("page", Page.Type).Unique(),
	}
}
