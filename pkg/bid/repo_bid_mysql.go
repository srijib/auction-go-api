package bid

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

func (r *MysqlRepository) Find(id int) (*e.Bid, error) {
	bid := e.Bid{Id: id}
	err := r.DB.Find(&bid)

	if err.RowsAffected == 0 {
		return nil, nil
	} else if err == nil {
		return &bid, nil
	} else {
		return &bid, nil
	}
}

func (r *MysqlRepository) Save(b *e.Bid) (*e.Bid, error) {
	err := r.DB.Save(&b)
	if err.RowsAffected == 0 {
		return nil, nil
	} else if err == nil {
		return b, nil
	} else {
		return b, nil
	}
}

func (r *MysqlRepository) Update(id int, key string, val interface{}) (*e.Bid, error) {
	var bid e.Bid
	if err := r.DB.Where("id = ?", id).First(&bid).Error; err != nil {
		return nil, e.ErrNotFound
	}
	r.DB.Model(&bid).Update(key, val)
	return &bid, nil
}
