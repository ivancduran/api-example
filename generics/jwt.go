package generics

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	UserName string
}

func (j JWT) Secret() []byte {
	return []byte("osemidolktlejlebd123x12leaxcmsen")
}

// MapClaims creates a JSON web token
func (j JWT) MapClaims(email string, role int, account_id uint, id uint, account_entry string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	secs := now.Unix()

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["iat"] = secs
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["account_id"] = account_id
	claims["account_entry"] = account_entry
	claims["user_id"] = id

	// Generate encoded token and send it as response.
	t, err := token.SignedString(j.Secret())
	return t, err
}
