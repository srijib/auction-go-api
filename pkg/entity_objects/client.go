package entity_objects

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Client struct {
	Id       int `gorm:"primary_key";"AUTO_INCREMENT"`
	Username string
	Email    string
	Password string
}

func (client *Client) GenerateJWT(key []byte) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	clientId := client.Id

	claims["clientId"] = clientId
	claims["created"] = time.Now().Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(key)

	jwtMap := map[string]string{"token": tokenString}

	return jwtMap
}
