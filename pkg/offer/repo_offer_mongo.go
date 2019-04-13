package offer

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

func CreateMongoRepository(p *mgosession.Pool, db string) Repository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

func (r *MongoRepository) Find(id e.ID) (*e.Offer, error) {
	result := e.Offer{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
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

func (r *MongoRepository) Save(o *e.Offer) (e.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Insert(o)
	if err != nil {
		return e.ID(0), err
	}
	return o.Id, nil
}

func (r *MongoRepository) Query(page int, size int, sortkey string) ([]*e.Offer, error) {

	if size == 0 {
		size = 10
	}

	if sortkey == "" {
		sortkey = "go_live"
	}

	var res []*e.Offer
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Find(nil).Sort(sortkey).Limit(size).Skip(page).All(&res)
	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}

func (r *MongoRepository) FindByKey(key string, val interface{}, page int, size int) ([]*e.Offer, error) {

	if size == 0 {
		size = 10
	}

	var res []*e.Offer
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	err := coll.Find(bson.M{key: val}).Limit(size).Skip(page).All(&res)
	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}

func (r *MongoRepository) Update(id e.ID, key string, val interface{}) (*e.Offer, error) {
	result := e.Offer{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("offers")
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{key: val}},
		ReturnNew: true,
	}
	_, err := coll.Find(bson.M{"_id": id}).Apply(change, &result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, err
	}
}
