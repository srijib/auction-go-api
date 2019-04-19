package offer

import (
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

//Repository repository interface
type Repository interface {
	Find(id int) (*e.Offer, error)
	Save(offer *e.Offer) (*e.Offer, error)
	Query(page int, size int, sortkey string) ([]*e.Offer, error)
	Update(id int, key string, val interface{}) (*e.Offer, error)
	SoldOffers() ([]*e.Offer, error)
}

//UseCase for offer
type UseCase interface {
	Find(id int) (*e.Offer, error)
	Save(user *e.Offer) (*e.Offer, error)
	Query(page int, size int, sortkey string) ([]*e.Offer, error)
	Update(id int, key string, val interface{}) (*e.Offer, error)
	SoldOffers() ([]*e.Offer, error)
}
