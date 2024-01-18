package orm

import (
	"parsing/dto"
	"parsing/orm/ent"
)

func PageDTO(p *ent.Page) *dto.Page {
	return &dto.Page{
		Job:       p.Job,
		Brand:     p.Brand,
		Domain:    p.Domain,
		URL:       p.URL,
		HTML:      p.HTML,
		Created:   p.Created,
		Updated:   p.Updated,
		Processed: p.Processed,
	}
}

func PagesDTO(p []*ent.Page) []*dto.Page {

	var pages []*dto.Page

	for _, page := range p {
		pages = append(pages, PageDTO(page))
	}

	return pages
}

func PageEntity(p *dto.Page) *ent.Page {
	return &ent.Page{
		Job:       p.Job,
		Brand:     p.Brand,
		Domain:    p.Domain,
		URL:       p.URL,
		HTML:      p.HTML,
		Created:   p.Created,
		Updated:   p.Updated,
		Processed: p.Processed,
	}
}

func PagesEntity(p []*dto.Page) []*ent.Page {

	var pages []*ent.Page

	for _, page := range p {
		pages = append(pages, PageEntity(page))
	}

	return pages
}

func ProductDTO(p *ent.Product) *dto.Product {
	return &dto.Product{
		ID:         p.ID,
		Series:     p.Series,
		Model:      p.Model,
		Sku:        p.Sku,
		Properties: p.Properties,
	}
}

func ProductsDTO(p []*ent.Product) []*dto.Product {

	var products []*dto.Product

	for _, product := range p {
		products = append(products, ProductDTO(product))
	}

	return products
}

func ProductEntity(p *dto.Product) *ent.Product {
	return &ent.Product{
		ID:         p.ID,
		Series:     p.Series,
		Model:      p.Model,
		Sku:        p.Sku,
		Properties: p.Properties,
	}
}

func ProductsEntity(p []*dto.Product) []*ent.Product {

	var products []*ent.Product

	for _, product := range p {
		products = append(products, ProductEntity(product))
	}

	return products
}
