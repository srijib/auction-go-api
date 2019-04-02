package user

import e "github.com/urmilagera/auction/pkg/entity_objects"

type Repository interface {
	Find(id e.ID) (*e.Client, error)
	Save(user *e.Client) (e.ID, error)
	FindByKey(key string, val interface{}) ([]*e.Client, error)
}

type UseCase interface {
	Find(id e.ID) (*e.Client, error)
	Save(user *e.Client) (e.ID, error)
	FindByKey(key string, val interface{}) ([]*e.Client, error)
	FindByUsername(username string) ([]*e.Client, error)
}
