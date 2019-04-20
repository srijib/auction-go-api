package offer

import (
	"github.com/jinzhu/gorm"
	e "github.com/urmilagera/auction/pkg/entity_objects"
)

type MysqlRepository struct {
	DB *gorm.DB
}

func CreateMysqlRepository(r *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		DB: r,
	}
}

func (r *MysqlRepository) Find(id int) (*e.Offer, error) {
	result := e.Offer{Id: id}
	err := r.DB.Find(&result)
	if err.RowsAffected == 0 {
		return nil, nil
	} else if err == nil {
		return &result, nil
	} else {
		return &result, nil
	}
}

func (r *MysqlRepository) Save(o *e.Offer) (*e.Offer, error) {
	err := r.DB.Save(&o)
	if err != nil {
		return nil, err.Error
	}
	return o, nil
}

func (r *MysqlRepository) Query(page int, size int, sortkey string) ([]*e.Offer, error) {
	var offset int
	if size == 0 {
		size = 10
	}

	if sortkey == "" {
		sortkey = "go_live"
	}

	if page == 0 {
		offset = 0
	} else {
		offset = size * page
	}
	var result []*e.Offer
	err := r.DB.Order(sortkey).Limit(size).Offset(offset).Find(&result)
	if err.RecordNotFound() {
		return nil, e.ErrNotFound
	} else if err.Error == nil {
		return result, nil
	} else {
		return nil, err.Error
	}
}

func (r *MysqlRepository) Update(id int, key string, val interface{}) (*e.Offer, error) {
	var offer e.Offer
	if err := r.DB.Where("id = ?", id).First(&offer).Error; err != nil {
		return nil, e.ErrNotFound
	}
	r.DB.Model(&offer).Update(key, val)
	return &offer, nil
}

func (app *MysqlRepository) SoldOffers() ([]*e.Offer, error) {
	var result []*e.Offer
	if err := app.DB.Where("sold = ?", true).Find(&result).Error; err != nil {
		return nil, e.ErrNotFound
	}
	return result, nil
}
