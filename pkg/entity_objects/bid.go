package entity_objects

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Bid struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice  float64
	OfferId   int
	Client    Client `gorm:"foreignkey:ClientId"`
	Timestamp time.Time
	Accepted  bool
	ClientId  int
}

func (bid *Bid) Validate() bool {
	if bid.BidPrice <= 0 {
		return false
	}
	bid.Timestamp = time.Now()
	return true
}
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Bid{})
	db.AutoMigrate(&Offer{})
	db.AutoMigrate(&Client{})
	return db
}
