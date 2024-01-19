package plugin

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"parsing/dto"
	"parsing/parse"
	"regexp"
	"strings"
)

const BrandHydac = "HYDAC"
const BrandDomain = "www.hydac.com"

func NewHydacJob() parse.Job {

	// if url following mask /shop/en/{number}{end of string} with regex
	exp, err := regexp.Compile(`\/shop\/en\/\d+$`)
	if err != nil {
		log.Fatal(err)
	}

	job := parse.NewDefaultJob()
	job.AllowedDomains = []string{BrandDomain}
	job.StartingURL = "https://www.hydac.com/shop/en/hps-2400-1000496612"
	job.OnLink = func(e *colly.HTMLElement) error {

		url := e.Attr("href")
		if strings.Contains(url, "shop/en") {
			return nil
		}

		return errors.New("url not following mask /shop/en")
	}

	job.OnPage = func(e *colly.Response) (dto.Page, error) {

		// TODO: Save body to db

		url := e.Request.URL.String()
		if !exp.MatchString(url) {
			return dto.Page{}, errors.New("url not following mask /shop/en/{number}")
		}

		fmt.Println("OnPage")
		return dto.Page{
			URL:    url,
			HTML:   string(e.Body),
			Brand:  BrandHydac,
			Domain: BrandDomain,
		}, nil
	}

	job.OnProduct = func(p dto.Page) (dto.Product, error) {

		fmt.Println("OnProduct")
		return ParsePageHydac(p), nil
	}
	return job
}

func ParsePageHydac(page dto.Page) dto.Product {
	return dto.Product{
		Series:     "test",
		Model:      "test",
		Sku:        "test",
		Properties: map[string]string{"test": "test"},
	}
}
