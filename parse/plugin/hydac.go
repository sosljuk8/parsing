package plugin

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
)

func NewHydacJob() Job {

	// if url following mask /shop/en/{number}{end of string} with regex
	exp, err := regexp.Compile(`\/shop\/en\/\d+$`)
	if err != nil {
		log.Fatal(err)
	}

	job := NewDefaultJob()
	job.AllowedDomains = []string{"www.hydac.com"}
	job.StartingURL = "https://www.hydac.com/shop/en"
	job.OnLink = func(e *colly.HTMLElement) error {

		url := e.Attr("href")
		if exp.MatchString(url) {
			return nil
		}

		return errors.New("url not following mask /shop/en/{number}")
	}
	job.OnPage = func(e *colly.Response) (Page, error) {

		// TODO: Save body to db

		fmt.Println("OnPage")
		return Page{}, nil
	}

	return job
}
