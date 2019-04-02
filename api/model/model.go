package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
type Client struct {
    ClientId int `gorm:"primary_key"`
    Name  string
    Email string
    Bid []Bid `gorm:"ForeignKey:BidId"` //you need to do like this
    BidId int
}
type Offer struct {
    OfferId int `gorm:"primary_key"`
    BidPrice float32
    GoLive time.Time
    LifeTime int
    PhotoUrl string
    Title string
    Bids []Bid `gorm:"ForeignKey:BidId"` //you need to do like this
    BidId int
}
type Bid struct {
    BidId int `gorm:"primary_key"`
    BidPrice string
    Client Client `gorm:"ForeignKey:ClientId"`
    Offer Offer `gorm:"ForeignKey:OfferId"`
    ClientId int
    OfferId int
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Offer{})
	db.AutoMigrate(&Bid{})
	return db
}