package entity_objects

import (
	"time"
)

type Offer struct {
	Id        int       `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice  float64   `json:"bid_price"`
	GoLive    time.Time `json:"go_live"`
	Lifetime  int       `json:"life_time"`
	PhotoUrl  string    `json:"photo_url"`
	Title     string    `json:"title"`
	Sold      bool      `json:"sold"`
	CreatedBy string    `json:"created_by"`
	BidId     int       `json:"bid_id"`
	Bid       *Bid      `gorm:"foreignkey:OfferId"` //you need to do like this
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
