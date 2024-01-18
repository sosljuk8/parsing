package plugin

import (
	"github.com/gocolly/colly/v2"
	"parsing/dto"
)

type Job struct {
	AllowedDomains []string
	StartingURL    string
	OnLink         func(e *colly.HTMLElement) error
	OnPage         func(e *colly.Response) (dto.Page, error)
}

func NewDefaultJob() Job {
	return Job{
		AllowedDomains: []string{},
		StartingURL:    "",
		OnLink: func(e *colly.HTMLElement) error {
			return nil
		},
		OnPage: func(e *colly.Response) (dto.Page, error) {
			return dto.Page{
				URL:  e.Request.URL.String(),
				HTML: string(e.Body),
			}, nil
		},
	}
}
