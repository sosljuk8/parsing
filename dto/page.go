package dto

import (
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
