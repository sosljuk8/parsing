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

func ProductDTO(p *ent.Product) *dto.Product {
	return &dto.Product{
		ID:         p.ID,
		Series:     p.Series,
		Model:      p.Model,
		Sku:        p.Sku,
		Properties: p.Properties,
	}
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
