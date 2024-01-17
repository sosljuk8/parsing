// Code generated by ent, DO NOT EDIT.

package category

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBrand holds the string denoting the brand field in the database.
	FieldBrand = "brand"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldUpdated holds the string denoting the updated field in the database.
	FieldUpdated = "updated"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParents holds the string denoting the parents edge name in mutations.
	EdgeParents = "parents"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// ProductsTable is the table that holds the products relation/edge. The primary key declared below.
	ProductsTable = "category_products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ChildrenTable is the table that holds the children relation/edge. The primary key declared below.
	ChildrenTable = "category_parents"
	// ParentsTable is the table that holds the parents relation/edge. The primary key declared below.
	ParentsTable = "category_parents"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldBrand,
	FieldName,
	FieldCreated,
	FieldUpdated,
}

var (
	// ProductsPrimaryKey and ProductsColumn2 are the table columns denoting the
	// primary key for the products relation (M2M).
	ProductsPrimaryKey = []string{"category_id", "product_id"}
	// ChildrenPrimaryKey and ChildrenColumn2 are the table columns denoting the
	// primary key for the children relation (M2M).
	ChildrenPrimaryKey = []string{"category_id", "child_id"}
	// ParentsPrimaryKey and ParentsColumn2 are the table columns denoting the
	// primary key for the parents relation (M2M).
	ParentsPrimaryKey = []string{"category_id", "child_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// BrandValidator is a validator for the "brand" field. It is called by the builders before save.
	BrandValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultUpdated holds the default value on creation for the "updated" field.
	DefaultUpdated func() time.Time
)

// OrderOption defines the ordering options for the Category queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBrand orders the results by the brand field.
func ByBrand(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrand, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreated orders the results by the created field.
func ByCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreated, opts...).ToFunc()
}

// ByUpdated orders the results by the updated field.
func ByUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdated, opts...).ToFunc()
}

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParentsCount orders the results by parents count.
func ByParentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParentsStep(), opts...)
	}
}

// ByParents orders the results by parents terms.
func ByParents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProductsTable, ProductsPrimaryKey...),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ChildrenTable, ChildrenPrimaryKey...),
	)
}
func newParentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ParentsTable, ParentsPrimaryKey...),
	)
}
