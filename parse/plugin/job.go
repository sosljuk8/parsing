package plugin

import (
	"errors"
	"github.com/gocolly/colly/v2"
	"parsing/dto"
)

type Job struct {
	// Name of job must be unique
	Name           string
	AllowedDomains []string
	StartingURL    string
	OnLink         func(e *colly.HTMLElement) error
	OnPage         func(e *colly.Response) (dto.Page, error)
	OnProduct      func(p dto.Page) (dto.Product, error)
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
		OnProduct: func(p dto.Page) (dto.Product, error) {
			return dto.Product{}, errors.New("OnProduct not implemented")
		},
	}
}

type Jobs struct {
	jobs map[string]Job
}

func NewJobs() Jobs {
	return Jobs{
		jobs: make(map[string]Job),
	}
}

// Add job to jobs by job name
func (j Jobs) Add(job Job) {
	j.jobs[job.Name] = job
}

// Get job from jobs by job name
func (j Jobs) Get(name string) Job {
	return j.jobs[name]
}

// All jobs
func (j Jobs) All() []Job {
	var jobs []Job
	for _, job := range j.jobs {
		jobs = append(jobs, job)
	}
	return jobs
}

// Remove job from jobs by job name
func (j Jobs) Remove(name string) {
	delete(j.jobs, name)
}
