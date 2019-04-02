package offer

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

func (s *Service) Save(offer *e.Offer) (e.ID, error) {
	offer.Id = e.NewID()
	return s.repo.Save(offer)
}

func (s *Service) Find(id e.ID) (*e.Offer, error) {
	return s.repo.Find(id)
}

func (s *Service) FindByKey(key string, val interface{}, page int, size int) ([]*e.Offer, error) {
	return s.repo.FindByKey(key, val, page, size)
}

func (s *Service) Query(page int, size int, sortkey string) ([]*e.Offer, error) {
	return s.repo.Query(page, size, sortkey)
}

func (s *Service) Update(id e.ID, key string, val interface{}) (*e.Offer, error) {
	return s.repo.Update(id, key, val)
}
