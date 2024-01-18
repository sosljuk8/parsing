package parse

import (
	"fmt"
	"parsing/dto"
	"parsing/parse/plugin"
	"time"
)

type MapperStore interface {
	LoadUnprocessed(jobName string, limit int) []dto.Page
	SaveProducts(...dto.Product) error
}

type ProductMapper struct {
	store MapperStore
	jobs  plugin.Jobs
}

func NewProductMapper(store MapperStore, j plugin.Jobs) ProductMapper {
	return ProductMapper{store: store, jobs: j}
}

func (p ProductMapper) Run() {
	for _, j := range p.jobs.All() {
		go p.Processor(j)
	}
}

func (p ProductMapper) Processor(j plugin.Job) {

	for {
		err := p.Process(j)
		if err != nil {
			fmt.Println("Processor Job", err.Error())
			time.Sleep(time.Hour)
		}
	}
}

func (p ProductMapper) Process(j plugin.Job) error {

	// load first 100 unprocessed pages
	pages := p.store.LoadUnprocessed(j.Name, 100)

	// for each page
	for _, page := range pages {
		// parse page
		product, err := j.OnProduct(page)
		if err != nil {
			return err
		}

		// save product
		err = p.save(product)
		if err != nil {
			return err
		}

		// mark page as processed
		//err = p.store.MarkAsProcessed(page)
		//if err != nil {
		//	return err
		//}
	}

	return nil
}

func (p ProductMapper) save(product dto.Product) error {
	return p.store.SaveProducts(product)
}
