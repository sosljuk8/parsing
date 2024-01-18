// Code generated by ent, DO NOT EDIT.

package ent

import (
	"parsing/orm/ent/category"
	"parsing/orm/ent/page"
	"parsing/orm/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescBrand is the schema descriptor for brand field.
	categoryDescBrand := categoryFields[0].Descriptor()
	// category.BrandValidator is a validator for the "brand" field. It is called by the builders before save.
	category.BrandValidator = categoryDescBrand.Validators[0].(func(string) error)
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[1].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	// categoryDescCreated is the schema descriptor for created field.
	categoryDescCreated := categoryFields[2].Descriptor()
	// category.DefaultCreated holds the default value on creation for the created field.
	category.DefaultCreated = categoryDescCreated.Default.(func() time.Time)
	// categoryDescUpdated is the schema descriptor for updated field.
	categoryDescUpdated := categoryFields[3].Descriptor()
	// category.DefaultUpdated holds the default value on creation for the updated field.
	category.DefaultUpdated = categoryDescUpdated.Default.(func() time.Time)
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescBrand is the schema descriptor for brand field.
	pageDescBrand := pageFields[0].Descriptor()
	// page.BrandValidator is a validator for the "brand" field. It is called by the builders before save.
	page.BrandValidator = pageDescBrand.Validators[0].(func(string) error)
	// pageDescDomain is the schema descriptor for domain field.
	pageDescDomain := pageFields[1].Descriptor()
	// page.DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	page.DomainValidator = pageDescDomain.Validators[0].(func(string) error)
	// pageDescJob is the schema descriptor for job field.
	pageDescJob := pageFields[2].Descriptor()
	// page.JobValidator is a validator for the "job" field. It is called by the builders before save.
	page.JobValidator = func() func(string) error {
		validators := pageDescJob.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(job string) error {
			for _, fn := range fns {
				if err := fn(job); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// pageDescHTML is the schema descriptor for html field.
	pageDescHTML := pageFields[3].Descriptor()
	// page.HTMLValidator is a validator for the "html" field. It is called by the builders before save.
	page.HTMLValidator = pageDescHTML.Validators[0].(func(string) error)
	// pageDescCreated is the schema descriptor for created field.
	pageDescCreated := pageFields[4].Descriptor()
	// page.DefaultCreated holds the default value on creation for the created field.
	page.DefaultCreated = pageDescCreated.Default.(func() time.Time)
	// pageDescUpdated is the schema descriptor for updated field.
	pageDescUpdated := pageFields[5].Descriptor()
	// page.DefaultUpdated holds the default value on creation for the updated field.
	page.DefaultUpdated = pageDescUpdated.Default.(func() time.Time)
	// pageDescURL is the schema descriptor for url field.
	pageDescURL := pageFields[7].Descriptor()
	// page.URLValidator is a validator for the "url" field. It is called by the builders before save.
	page.URLValidator = func() func(string) error {
		validators := pageDescURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(url string) error {
			for _, fn := range fns {
				if err := fn(url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
