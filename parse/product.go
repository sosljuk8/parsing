package parse

import (
	"parsing/dto"
)

type MapperStore interface {
	LoadUnprocessed(jobName string, limit int) []dto.Page
	SaveProducts(...dto.Product) error
}

type ProductMapper struct {
	store MapperStore
}

func NewProductMapper(store MapperStore) ProductMapper {
	return ProductMapper{store: store}
}

func (p ProductMapper) Run(jobName string, limit int) error {
	// load first 100 unprocessed pages
	pages := p.store.LoadUnprocessed(jobName, limit)

	// for each page
	for _, page := range pages {
		product := p.parse(page)
		p.save(product)
	}

	return nil
}

func (p ProductMapper) parse(page dto.Page) dto.Product {
	return dto.Product{
		Series:     "test",
		Model:      "test",
		Sku:        "test",
		Properties: map[string]string{"test": "test"},
	}
}

func (p ProductMapper) save(product dto.Product) error {
	return p.store.SaveProducts(product)
}
