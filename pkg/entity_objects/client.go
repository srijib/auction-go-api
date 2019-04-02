package entity_objects

type Client struct {
	Id       ID     `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
