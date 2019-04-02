package user

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

func (s *Service) Save(client *e.Client) (e.ID, error) {
	client.Id = e.NewID()
	return s.repo.Save(client)
}

func (s *Service) Find(id e.ID) (*e.Client, error) {
	return s.repo.Find(id)
}

func (s *Service) FindByKey(key string, val interface{}) ([]*e.Client, error) {
	return s.repo.FindByKey(key, val)
}

func (s *Service) FindByUsername(username string) ([]*e.Client, error) {
	return s.repo.FindByKey("username", username)
}
