// Code generated by ent, DO NOT EDIT.

package page

import (
	"parsing/orm/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldID, id))
}

// Brand applies equality check predicate on the "brand" field. It's identical to BrandEQ.
func Brand(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldBrand, v))
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldDomain, v))
}

// Job applies equality check predicate on the "job" field. It's identical to JobEQ.
func Job(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldJob, v))
}

// HTML applies equality check predicate on the "html" field. It's identical to HTMLEQ.
func HTML(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldHTML, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldCreated, v))
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldUpdated, v))
}

// Processed applies equality check predicate on the "processed" field. It's identical to ProcessedEQ.
func Processed(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldProcessed, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldURL, v))
}

// BrandEQ applies the EQ predicate on the "brand" field.
func BrandEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldBrand, v))
}

// BrandNEQ applies the NEQ predicate on the "brand" field.
func BrandNEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldBrand, v))
}

// BrandIn applies the In predicate on the "brand" field.
func BrandIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldBrand, vs...))
}

// BrandNotIn applies the NotIn predicate on the "brand" field.
func BrandNotIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldBrand, vs...))
}

// BrandGT applies the GT predicate on the "brand" field.
func BrandGT(v string) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldBrand, v))
}

// BrandGTE applies the GTE predicate on the "brand" field.
func BrandGTE(v string) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldBrand, v))
}

// BrandLT applies the LT predicate on the "brand" field.
func BrandLT(v string) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldBrand, v))
}

// BrandLTE applies the LTE predicate on the "brand" field.
func BrandLTE(v string) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldBrand, v))
}

// BrandContains applies the Contains predicate on the "brand" field.
func BrandContains(v string) predicate.Page {
	return predicate.Page(sql.FieldContains(FieldBrand, v))
}

// BrandHasPrefix applies the HasPrefix predicate on the "brand" field.
func BrandHasPrefix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasPrefix(FieldBrand, v))
}

// BrandHasSuffix applies the HasSuffix predicate on the "brand" field.
func BrandHasSuffix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasSuffix(FieldBrand, v))
}

// BrandEqualFold applies the EqualFold predicate on the "brand" field.
func BrandEqualFold(v string) predicate.Page {
	return predicate.Page(sql.FieldEqualFold(FieldBrand, v))
}

// BrandContainsFold applies the ContainsFold predicate on the "brand" field.
func BrandContainsFold(v string) predicate.Page {
	return predicate.Page(sql.FieldContainsFold(FieldBrand, v))
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldDomain, v))
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldDomain, v))
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldDomain, vs...))
}

// DomainNotIn applies the NotIn predicate on the "domain" field.
func DomainNotIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldDomain, vs...))
}

// DomainGT applies the GT predicate on the "domain" field.
func DomainGT(v string) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldDomain, v))
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldDomain, v))
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldDomain, v))
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldDomain, v))
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.Page {
	return predicate.Page(sql.FieldContains(FieldDomain, v))
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasPrefix(FieldDomain, v))
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasSuffix(FieldDomain, v))
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.Page {
	return predicate.Page(sql.FieldEqualFold(FieldDomain, v))
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.Page {
	return predicate.Page(sql.FieldContainsFold(FieldDomain, v))
}

// JobEQ applies the EQ predicate on the "job" field.
func JobEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldJob, v))
}

// JobNEQ applies the NEQ predicate on the "job" field.
func JobNEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldJob, v))
}

// JobIn applies the In predicate on the "job" field.
func JobIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldJob, vs...))
}

// JobNotIn applies the NotIn predicate on the "job" field.
func JobNotIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldJob, vs...))
}

// JobGT applies the GT predicate on the "job" field.
func JobGT(v string) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldJob, v))
}

// JobGTE applies the GTE predicate on the "job" field.
func JobGTE(v string) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldJob, v))
}

// JobLT applies the LT predicate on the "job" field.
func JobLT(v string) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldJob, v))
}

// JobLTE applies the LTE predicate on the "job" field.
func JobLTE(v string) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldJob, v))
}

// JobContains applies the Contains predicate on the "job" field.
func JobContains(v string) predicate.Page {
	return predicate.Page(sql.FieldContains(FieldJob, v))
}

// JobHasPrefix applies the HasPrefix predicate on the "job" field.
func JobHasPrefix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasPrefix(FieldJob, v))
}

// JobHasSuffix applies the HasSuffix predicate on the "job" field.
func JobHasSuffix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasSuffix(FieldJob, v))
}

// JobEqualFold applies the EqualFold predicate on the "job" field.
func JobEqualFold(v string) predicate.Page {
	return predicate.Page(sql.FieldEqualFold(FieldJob, v))
}

// JobContainsFold applies the ContainsFold predicate on the "job" field.
func JobContainsFold(v string) predicate.Page {
	return predicate.Page(sql.FieldContainsFold(FieldJob, v))
}

// HTMLEQ applies the EQ predicate on the "html" field.
func HTMLEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldHTML, v))
}

// HTMLNEQ applies the NEQ predicate on the "html" field.
func HTMLNEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldHTML, v))
}

// HTMLIn applies the In predicate on the "html" field.
func HTMLIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldHTML, vs...))
}

// HTMLNotIn applies the NotIn predicate on the "html" field.
func HTMLNotIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldHTML, vs...))
}

// HTMLGT applies the GT predicate on the "html" field.
func HTMLGT(v string) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldHTML, v))
}

// HTMLGTE applies the GTE predicate on the "html" field.
func HTMLGTE(v string) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldHTML, v))
}

// HTMLLT applies the LT predicate on the "html" field.
func HTMLLT(v string) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldHTML, v))
}

// HTMLLTE applies the LTE predicate on the "html" field.
func HTMLLTE(v string) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldHTML, v))
}

// HTMLContains applies the Contains predicate on the "html" field.
func HTMLContains(v string) predicate.Page {
	return predicate.Page(sql.FieldContains(FieldHTML, v))
}

// HTMLHasPrefix applies the HasPrefix predicate on the "html" field.
func HTMLHasPrefix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasPrefix(FieldHTML, v))
}

// HTMLHasSuffix applies the HasSuffix predicate on the "html" field.
func HTMLHasSuffix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasSuffix(FieldHTML, v))
}

// HTMLEqualFold applies the EqualFold predicate on the "html" field.
func HTMLEqualFold(v string) predicate.Page {
	return predicate.Page(sql.FieldEqualFold(FieldHTML, v))
}

// HTMLContainsFold applies the ContainsFold predicate on the "html" field.
func HTMLContainsFold(v string) predicate.Page {
	return predicate.Page(sql.FieldContainsFold(FieldHTML, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldCreated, v))
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldUpdated, v))
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldUpdated, v))
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldUpdated, vs...))
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldUpdated, vs...))
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldUpdated, v))
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldUpdated, v))
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldUpdated, v))
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldUpdated, v))
}

// ProcessedEQ applies the EQ predicate on the "processed" field.
func ProcessedEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldProcessed, v))
}

// ProcessedNEQ applies the NEQ predicate on the "processed" field.
func ProcessedNEQ(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldProcessed, v))
}

// ProcessedIn applies the In predicate on the "processed" field.
func ProcessedIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldProcessed, vs...))
}

// ProcessedNotIn applies the NotIn predicate on the "processed" field.
func ProcessedNotIn(vs ...time.Time) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldProcessed, vs...))
}

// ProcessedGT applies the GT predicate on the "processed" field.
func ProcessedGT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldProcessed, v))
}

// ProcessedGTE applies the GTE predicate on the "processed" field.
func ProcessedGTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldProcessed, v))
}

// ProcessedLT applies the LT predicate on the "processed" field.
func ProcessedLT(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldProcessed, v))
}

// ProcessedLTE applies the LTE predicate on the "processed" field.
func ProcessedLTE(v time.Time) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldProcessed, v))
}

// ProcessedIsNil applies the IsNil predicate on the "processed" field.
func ProcessedIsNil() predicate.Page {
	return predicate.Page(sql.FieldIsNull(FieldProcessed))
}

// ProcessedNotNil applies the NotNil predicate on the "processed" field.
func ProcessedNotNil() predicate.Page {
	return predicate.Page(sql.FieldNotNull(FieldProcessed))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Page {
	return predicate.Page(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Page {
	return predicate.Page(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Page {
	return predicate.Page(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Page {
	return predicate.Page(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Page {
	return predicate.Page(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Page {
	return predicate.Page(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Page {
	return predicate.Page(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Page {
	return predicate.Page(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Page {
	return predicate.Page(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Page {
	return predicate.Page(sql.FieldContainsFold(FieldURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Page) predicate.Page {
	return predicate.Page(sql.NotPredicates(p))
}
