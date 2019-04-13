package client

import (
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func CreateService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Save
func (s *Service) Save(client *e.Client) (e.ID, error) {
	client.Id = e.NewID()
	return s.repo.Save(client)
}

//Find
func (s *Service) Find(id e.ID) (*e.Client, error) {
	return s.repo.Find(id)
}

//FindByKey
func (s *Service) FindByKey(key string, val interface{}) ([]*e.Client, error) {
	return s.repo.FindByKey(key, val)
}

//FindByUsername
func (s *Service) FindByUsername(username string) ([]*e.Client, error) {
	return s.repo.FindByKey("username", username)
}
