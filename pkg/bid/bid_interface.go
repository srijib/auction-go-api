package bid

import (
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

type Repository interface {
	Find(id e.ID) (*e.Bid, error)
	Save(bid *e.Bid) (e.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*e.Bid, error)
	Update(id e.ID, key string, val interface{}) (*e.Bid, error)
}

type UseCase interface {
	Find(id e.ID) (*e.Bid, error)
	Save(bid *e.Bid) (e.ID, error)
	FindByKey(key string, val interface{}, page int, size int) ([]*e.Bid, error)
	Update(id e.ID, key string, val interface{}) (*e.Bid, error)
}
