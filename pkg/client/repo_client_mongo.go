package client

import (
	"github.com/juju/mgosession"
	e "github.com/urmilagera/auction/pkg/entity_objects"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

func CreateMongoRepo(p *mgosession.Pool, db string) Repository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

func (r *MongoRepository) Find(id e.ID) (*e.Client, error) {
	result := e.Client{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("clients")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}

func (r *MongoRepository) Save(client *e.Client) (e.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("clients")
	err := coll.Insert(client)
	if err != nil {
		return e.ID(0), err
	}
	return client.Id, nil
}

func (r *MongoRepository) FindByKey(key string, val interface{}) ([]*e.Client, error) {
	var result []*e.Client
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("clients")
	err := coll.Find(bson.M{key: val}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}
