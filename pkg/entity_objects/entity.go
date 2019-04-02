package entity_objects

import (
	"gopkg.in/mgo.v2/bson"
)

type ID bson.ObjectId

func (i ID) String() string {
	return bson.ObjectId(i).Hex()
}

func (i ID) MarshalJSON() ([]byte, error) {
	return bson.ObjectId(i).MarshalJSON()
}

func (i *ID) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = s[1 : len(s)-1]
	if bson.IsObjectIdHex(s) {
		*i = ID(bson.ObjectIdHex(s))
	}

	return nil
}

func (i ID) GetBSON() (interface{}, error) {
	if i == "" {
		return "", nil
	}
	return bson.ObjectId(i), nil
}

func (i *ID) SetBSON(raw bson.Raw) error {
	decoded := new(string)
	bsonErr := raw.Unmarshal(decoded)
	if bsonErr == nil {
		*i = ID(bson.ObjectId(*decoded))
		return nil
	}
	return bsonErr
}

func StringToID(s string) ID {
	return ID(bson.ObjectIdHex(s))
}

func IsValidID(s string) bool {
	return bson.IsObjectIdHex(s)
}

func NewID() ID {
	return StringToID(bson.NewObjectId().Hex())
}
