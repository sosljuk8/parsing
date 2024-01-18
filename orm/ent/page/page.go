// Code generated by ent, DO NOT EDIT.

package page

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the page type in the database.
	Label = "page"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBrand holds the string denoting the brand field in the database.
	FieldBrand = "brand"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldJob holds the string denoting the job field in the database.
	FieldJob = "job"
	// FieldHTML holds the string denoting the html field in the database.
	FieldHTML = "html"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldUpdated holds the string denoting the updated field in the database.
	FieldUpdated = "updated"
	// FieldProcessed holds the string denoting the processed field in the database.
	FieldProcessed = "processed"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// Table holds the table name of the page in the database.
	Table = "pages"
)

// Columns holds all SQL columns for page fields.
var Columns = []string{
	FieldID,
	FieldBrand,
	FieldDomain,
	FieldJob,
	FieldHTML,
	FieldCreated,
	FieldUpdated,
	FieldProcessed,
	FieldURL,
}

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
	// DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	DomainValidator func(string) error
	// JobValidator is a validator for the "job" field. It is called by the builders before save.
	JobValidator func(string) error
	// HTMLValidator is a validator for the "html" field. It is called by the builders before save.
	HTMLValidator func(string) error
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultUpdated holds the default value on creation for the "updated" field.
	DefaultUpdated func() time.Time
	// DefaultProcessed holds the default value on creation for the "processed" field.
	DefaultProcessed time.Time
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
)

// OrderOption defines the ordering options for the Page queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBrand orders the results by the brand field.
func ByBrand(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrand, opts...).ToFunc()
}

// ByDomain orders the results by the domain field.
func ByDomain(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDomain, opts...).ToFunc()
}

// ByJob orders the results by the job field.
func ByJob(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJob, opts...).ToFunc()
}

// ByHTML orders the results by the html field.
func ByHTML(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTML, opts...).ToFunc()
}

// ByCreated orders the results by the created field.
func ByCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreated, opts...).ToFunc()
}

// ByUpdated orders the results by the updated field.
func ByUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdated, opts...).ToFunc()
}

// ByProcessed orders the results by the processed field.
func ByProcessed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProcessed, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}
