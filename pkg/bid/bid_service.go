package bid

import (
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

type Service struct {
	repo Repository
}

func CreateService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Save(bid *e.Bid) (*e.Bid, error) {
	return s.repo.Save(bid)
}

func (s *Service) Find(id int) (*e.Bid, error) {
	return s.repo.Find(id)
}

func (s *Service) Update(id int, key string, val interface{}) (*e.Bid, error) {
	return s.repo.Update(id, key, val)
}
