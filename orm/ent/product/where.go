// Code generated by ent, DO NOT EDIT.

package product

import (
	"parsing/orm/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// Series applies equality check predicate on the "series" field. It's identical to SeriesEQ.
func Series(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSeries, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldModel, v))
}

// Sku applies equality check predicate on the "sku" field. It's identical to SkuEQ.
func Sku(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSku, v))
}

// SeriesEQ applies the EQ predicate on the "series" field.
func SeriesEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSeries, v))
}

// SeriesNEQ applies the NEQ predicate on the "series" field.
func SeriesNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldSeries, v))
}

// SeriesIn applies the In predicate on the "series" field.
func SeriesIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldSeries, vs...))
}

// SeriesNotIn applies the NotIn predicate on the "series" field.
func SeriesNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldSeries, vs...))
}

// SeriesGT applies the GT predicate on the "series" field.
func SeriesGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldSeries, v))
}

// SeriesGTE applies the GTE predicate on the "series" field.
func SeriesGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldSeries, v))
}

// SeriesLT applies the LT predicate on the "series" field.
func SeriesLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldSeries, v))
}

// SeriesLTE applies the LTE predicate on the "series" field.
func SeriesLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldSeries, v))
}

// SeriesContains applies the Contains predicate on the "series" field.
func SeriesContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldSeries, v))
}

// SeriesHasPrefix applies the HasPrefix predicate on the "series" field.
func SeriesHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldSeries, v))
}

// SeriesHasSuffix applies the HasSuffix predicate on the "series" field.
func SeriesHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldSeries, v))
}

// SeriesIsNil applies the IsNil predicate on the "series" field.
func SeriesIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldSeries))
}

// SeriesNotNil applies the NotNil predicate on the "series" field.
func SeriesNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldSeries))
}

// SeriesEqualFold applies the EqualFold predicate on the "series" field.
func SeriesEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldSeries, v))
}

// SeriesContainsFold applies the ContainsFold predicate on the "series" field.
func SeriesContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldSeries, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldModel, v))
}

// ModelIsNil applies the IsNil predicate on the "model" field.
func ModelIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldModel))
}

// ModelNotNil applies the NotNil predicate on the "model" field.
func ModelNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldModel))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldModel, v))
}

// SkuEQ applies the EQ predicate on the "sku" field.
func SkuEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldSku, v))
}

// SkuNEQ applies the NEQ predicate on the "sku" field.
func SkuNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldSku, v))
}

// SkuIn applies the In predicate on the "sku" field.
func SkuIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldSku, vs...))
}

// SkuNotIn applies the NotIn predicate on the "sku" field.
func SkuNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldSku, vs...))
}

// SkuGT applies the GT predicate on the "sku" field.
func SkuGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldSku, v))
}

// SkuGTE applies the GTE predicate on the "sku" field.
func SkuGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldSku, v))
}

// SkuLT applies the LT predicate on the "sku" field.
func SkuLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldSku, v))
}

// SkuLTE applies the LTE predicate on the "sku" field.
func SkuLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldSku, v))
}

// SkuContains applies the Contains predicate on the "sku" field.
func SkuContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldSku, v))
}

// SkuHasPrefix applies the HasPrefix predicate on the "sku" field.
func SkuHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldSku, v))
}

// SkuHasSuffix applies the HasSuffix predicate on the "sku" field.
func SkuHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldSku, v))
}

// SkuIsNil applies the IsNil predicate on the "sku" field.
func SkuIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldSku))
}

// SkuNotNil applies the NotNil predicate on the "sku" field.
func SkuNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldSku))
}

// SkuEqualFold applies the EqualFold predicate on the "sku" field.
func SkuEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldSku, v))
}

// SkuContainsFold applies the ContainsFold predicate on the "sku" field.
func SkuContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldSku, v))
}

// PropertiesIsNil applies the IsNil predicate on the "properties" field.
func PropertiesIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldProperties))
}

// PropertiesNotNil applies the NotNil predicate on the "properties" field.
func PropertiesNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldProperties))
}

// HasCategories applies the HasEdge predicate on the "categories" edge.
func HasCategories() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoriesTable, CategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoriesWith applies the HasEdge predicate on the "categories" edge with a given conditions (other predicates).
func HasCategoriesWith(preds ...predicate.Category) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPage applies the HasEdge predicate on the "page" edge.
func HasPage() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PageTable, PageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPageWith applies the HasEdge predicate on the "page" edge with a given conditions (other predicates).
func HasPageWith(preds ...predicate.Page) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newPageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
