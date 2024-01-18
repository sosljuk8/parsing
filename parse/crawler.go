package parse

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"parsing/dto"
	"parsing/parse/plugin"
)

type PageStore interface {
	Save(page dto.Page) error
}

type Crawler struct {
	store PageStore
	jobs  plugin.Jobs
}

func NewCrawler(s PageStore, j plugin.Jobs) Crawler {
	return Crawler{
		store: s,
		jobs:  j,
	}
}

func (p Crawler) Run() {

	for _, j := range p.jobs.All() {
		p.Process(j)
	}
}

func (p Crawler) Process(j plugin.Job) {

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(j.AllowedDomains...),
	)

	// On every an element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		err := j.OnLink(e)
		if err != nil {
			fmt.Println("OnLink failed", err.Error())
			return
		}

		link := e.Attr("href")
		err = c.Visit(e.Request.AbsoluteURL(link))
		if err != nil {
			fmt.Println("Visiting failed", err.Error())
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)

		page, err := j.OnPage(r)
		if err != nil {
			fmt.Println("OnResponse failed", err.Error())
			return
		}

		err = p.store.Save(page)
		if err != nil {
			fmt.Println("Page Save failed", err.Error())
			return
		}
	})

	err := c.Visit(j.StartingURL)
	if err != nil {
		fmt.Println("Visiting failed!!!!! to ", err.Error())
	}
}
