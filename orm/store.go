package orm

import (
	"context"
	"parsing/dto"
	"parsing/orm/ent"
)

type Store struct {
	ent *ent.Client
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
		SetURL(page.Url).
		SetHTML(page.HTML).
		Save(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
