package entity_objects

import (
	"time"
)

type Offer struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice  float64
	GoLive    time.Time
	Lifetime  int64
	PhotoURL  string
	Title     string
	CreatedBy string
	Sold      bool
}

func (offer *Offer) Validate() bool {
	if offer.BidPrice == 0 || offer.Title == "" || offer.Lifetime < 0 {
		return false
	}
	if offer.GoLive.Before(time.Now()) {
		offer.GoLive = time.Now()
	}
	return true
}
