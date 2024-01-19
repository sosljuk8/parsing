package parse

import (
	"fmt"
	"parsing/dto"
	"time"
)

type MapperStore interface {
	UnprocessedPages(jobName string, limit int) ([]*dto.Page, error)
	SaveProducts(...*dto.Product) error
}

type Mapper struct {
	store MapperStore
	jobs  Jobs
}

func NewMapper(store MapperStore, j Jobs) Mapper {
	return Mapper{store: store, jobs: j}
}

func (p Mapper) Run() {
	for _, j := range p.jobs.All() {
		go p.Processor(j)
	}
}

func (p Mapper) Processor(j Job) {

	for {
		err := p.Process(j)
		if err != nil {
			fmt.Println("Processor Job", err.Error())
			time.Sleep(time.Hour)
		}
	}
}

func (p Mapper) Process(j Job) error {

	// load first 100 unprocessed pages
	pages, loadErr := p.store.UnprocessedPages(j.Name, 100)
	if loadErr != nil {
		return loadErr
	}

	// for each page
	for _, page := range pages {
		// parse page
		product, err := j.OnProduct(*page)
		if err != nil {
			return err
		}

		// save product
		err = p.save(&product)
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

func (p Mapper) save(product *dto.Product) error {
	return p.store.SaveProducts(product)
}
