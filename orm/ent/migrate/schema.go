// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "brand", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeTime},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// PagesColumns holds the columns for the "pages" table.
	PagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "brand", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "job", Type: field.TypeString, Size: 255},
		{Name: "html", Type: field.TypeString, Size: 2147483647},
		{Name: "created", Type: field.TypeTime},
		{Name: "updated", Type: field.TypeTime},
		{Name: "processed", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 2048},
	}
	// PagesTable holds the schema information for the "pages" table.
	PagesTable = &schema.Table{
		Name:       "pages",
		Columns:    PagesColumns,
		PrimaryKey: []*schema.Column{PagesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "series", Type: field.TypeString, Nullable: true},
		{Name: "model", Type: field.TypeString, Nullable: true},
		{Name: "sku", Type: field.TypeString, Nullable: true},
		{Name: "properties", Type: field.TypeJSON, Nullable: true},
		{Name: "product_page", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_pages_page",
				Columns:    []*schema.Column{ProductsColumns[5]},
				RefColumns: []*schema.Column{PagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoryProductsColumns holds the columns for the "category_products" table.
	CategoryProductsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeInt},
		{Name: "product_id", Type: field.TypeInt},
	}
	// CategoryProductsTable holds the schema information for the "category_products" table.
	CategoryProductsTable = &schema.Table{
		Name:       "category_products",
		Columns:    CategoryProductsColumns,
		PrimaryKey: []*schema.Column{CategoryProductsColumns[0], CategoryProductsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_products_category_id",
				Columns:    []*schema.Column{CategoryProductsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_products_product_id",
				Columns:    []*schema.Column{CategoryProductsColumns[1]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CategoryParentsColumns holds the columns for the "category_parents" table.
	CategoryParentsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeInt},
		{Name: "child_id", Type: field.TypeInt},
	}
	// CategoryParentsTable holds the schema information for the "category_parents" table.
	CategoryParentsTable = &schema.Table{
		Name:       "category_parents",
		Columns:    CategoryParentsColumns,
		PrimaryKey: []*schema.Column{CategoryParentsColumns[0], CategoryParentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_parents_category_id",
				Columns:    []*schema.Column{CategoryParentsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_parents_child_id",
				Columns:    []*schema.Column{CategoryParentsColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		PagesTable,
		ProductsTable,
		CategoryProductsTable,
		CategoryParentsTable,
	}
)

func init() {
	ProductsTable.ForeignKeys[0].RefTable = PagesTable
	CategoryProductsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategoryProductsTable.ForeignKeys[1].RefTable = ProductsTable
	CategoryParentsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategoryParentsTable.ForeignKeys[1].RefTable = CategoriesTable
}
