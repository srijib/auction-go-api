package client

import e "github.com/urmilagera/auction/pkg/entity_objects"

type Repository interface {
	Find(id int) (*e.Client, error)
	Save(client *e.Client) (*e.Client, error)
	FindByKey(key string, val interface{}) ([]*e.Client, error)
}

type UseCase interface {
	Find(id int) (*e.Client, error)
	Save(client *e.Client) (*e.Client, error)
	FindByKey(key string, val interface{}) ([]*e.Client, error)
	FindByUsername(username string) ([]*e.Client, error)
}
