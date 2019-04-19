package client

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

func (r *MysqlRepository) Find(id int) (*e.Client, error) {
	client := e.Client{Id: id}
	err := r.DB.Find(&client)

	if err.RecordNotFound() {
		return nil, e.ErrNotFound
	} else if err == nil {
		return &client, nil
	} else {
		return nil, err.Error
	}
}

func (r *MysqlRepository) Save(client *e.Client) (*e.Client, error) {
	err := r.DB.Save(&client)
	if err != nil {
		return nil, err.Error
	}
	return client, nil
}

func (r *MysqlRepository) FindByKey(key string, val interface{}) ([]*e.Client, error) {
	var result []*e.Client
	err := r.DB.Where(key+"= ?", val).Find(&result)

	if err.RowsAffected == 0 {
		return nil, nil
	} else if err == nil {
		return result, nil
	} else {
		return result, nil
	}
}
