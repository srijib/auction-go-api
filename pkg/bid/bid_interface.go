package bid

import (
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

type Repository interface {
	Find(id int) (*e.Bid, error)
	Save(bid *e.Bid) (*e.Bid, error)
	Update(id int, key string, val interface{}) (*e.Bid, error)
}

type UseCase interface {
	Find(id int) (*e.Bid, error)
	Save(bid *e.Bid) (*e.Bid, error)
	Update(id int, key string, val interface{}) (*e.Bid, error)
}
