package dto

import (
	"sync"
	"time"
)

type Page struct {
	Job       string
	Brand     string
	Domain    string
	URL       string
	HTML      string
	Created   time.Time
	Updated   time.Time
	Processed time.Time
}

// Pages is not safe for concurrent use collection of Page based on map with URL as a key
type Pages struct {
	items map[string]*Page
	mux   sync.RWMutex
}

// NewPages creates new Pages
func NewPages() *Pages {
	return &Pages{
		items: make(map[string]*Page),
		mux:   sync.RWMutex{},
	}
}

// Add adds page to collection
func (p *Pages) Add(page *Page) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.items[page.URL] = page

}

// All returns all pages
func (p *Pages) All() []*Page {
	var pages []*Page
	for _, page := range p.items {
		pages = append(pages, page)
	}
	return pages
}

// Len returns number of pages
func (p *Pages) Len() int {
	return len(p.items)
}

// URLs returns all urls
func (p *Pages) URLs() []string {
	var urls []string
	for url := range p.items {
		urls = append(urls, url)
	}
	return urls
}

func (p *Pages) Remove(url string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	delete(p.items, url)
}
