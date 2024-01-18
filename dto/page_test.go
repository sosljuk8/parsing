package dto_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"parsing/dto"
	"testing"
)

func TestNewPages(t *testing.T) {

	// create empty pages
	pages := dto.NewPages()
	count := 0

	// add 10 pages
	for i := 1; i <= 10; i++ {

		// add page
		pages.Add(NewPage(t))
		count++

		// check by length
		assert.Equal(t, count, pages.Len())
		assert.Equal(t, count, len(pages.URLs()))
		assert.Equal(t, count, len(pages.All()))
	}

	// remove 5 pages
	for i := 1; i <= 5; i++ {

		// remove page
		pages.Remove(pages.All()[0].URL)
		count--

		// check by length
		assert.Equal(t, count, pages.Len())
		assert.Equal(t, count, len(pages.URLs()))
		assert.Equal(t, count, len(pages.All()))
	}
}

func NewPage(t *testing.T) *dto.Page {
	return &dto.Page{
		Job:       gofakeit.UUID(),
		Brand:     gofakeit.Company(),
		Domain:    gofakeit.DomainName(),
		URL:       gofakeit.URL(),
		HTML:      gofakeit.Paragraph(1, 10, 10, " "),
		Created:   gofakeit.Date(),
		Updated:   gofakeit.Date(),
		Processed: gofakeit.Date(),
	}
}
