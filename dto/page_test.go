package dto_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"parsing/dto"
	"testing"
)

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
