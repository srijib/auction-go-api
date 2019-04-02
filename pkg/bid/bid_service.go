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

func (s *Service) Save(bid *e.Bid) (e.ID, error) {
	bid.Id = e.NewID()
	return s.repo.Save(bid)
}

func (s *Service) Find(id e.ID) (*e.Bid, error) {
	return s.repo.Find(id)
}

func (s *Service) FindByKey(key string, val interface{}, page int, size int) ([]*e.Bid, error) {
	return s.repo.FindByKey(key, val, page, size)
}

func (s *Service) Update(id e.ID, key string, val interface{}) (*e.Bid, error) {
	return s.repo.Update(id, key, val)
}
