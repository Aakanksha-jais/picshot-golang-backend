package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
)

type Claims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func newClaims(userID int64) *Claims {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.Must(uuid.NewUUID()).String(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // expiration time after which the token must be disregarded
			IssuedAt:  time.Now().Unix(),                     // time at which the token was issued
			NotBefore: time.Now().Unix(),                     // time before which the token must be disregarded
		},
		UserID: userID, // id of the user
	}
	return claims
}

func (c *Claims) Valid() error {
	if !c.VerifyExpiresAt(time.Now().Unix(), true) {
		return errors.AuthError{Msg: "token expired. log in again"}
	}

	if c.UserID == 0 {
		return errors.AuthError{Msg: "user id not present in claims"}
	}

	return nil
}
