package orm

import (
	"context"
	"parsing/dto"
	"parsing/orm/ent"
	"parsing/orm/ent/page"
	"time"
)

type Store struct {
	ent     *ent.Client
	visited map[string]bool
}

func NewStore(ent *ent.Client) *Store {
	return &Store{ent: ent}
}

func (s *Store) Close() error {
	return s.ent.Close()
}

// Save dto.Page
func (s *Store) Save(page dto.Page) error {

	_, err := s.ent.Page.
		Create().
		SetURL(page.URL).
		SetHTML(page.HTML).
		SetBrand(page.Brand).
		SetDomain(page.Domain).
		Save(context.TODO())

	// set visited
	if err != nil {
		s.visited[page.URL] = true
	}

	return err
}

// loadVisited loads all visited pages from db
func (s *Store) loadVisited() error {

	s.visited = make(map[string]bool)

	pages, err := s.ent.Page.Query().All(context.Background())
	if err != nil {
		return err
	}

	for _, p := range pages {
		s.visited[p.URL] = true
	}

	return nil
}

// IsVisited checks if page with url already exists in db
func (s *Store) IsVisited(url string) bool {

	if s.visited == nil {
		s.loadVisited()
	}
	return s.visited[url]
}

func (s *Store) UnprocessedPages(jobName string, limit int) ([]*dto.Page, error) {

	entities, err := s.ent.Page.Query().
		Where(page.Job(jobName), page.Processed(time.Time{})).
		Order(ent.Desc(page.FieldCreated)).
		Limit(limit).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return PagesDTO(entities), nil
}

func (s *Store) SaveProducts(products ...*dto.Product) error {

	// ent.ProductCreate used for bulk insert
	// https://entgo.io/docs/crud/#create-many
	entities := make([]*ent.ProductCreate, len(products))

	for _, product := range products {

		// create empty entity
		entity := s.ent.Product.Create()

		// set fields from dto.Product to the entity
		ProductCreateEntity(product, entity)

		// append entity to the slice
		entities = append(entities, entity)
	}

	// use slice for bulk insert
	_, err := s.ent.Product.
		CreateBulk(entities...).
		Save(context.TODO())

	return err
}
