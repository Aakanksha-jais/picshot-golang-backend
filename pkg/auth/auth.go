package auth

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTContextKey string

type Claims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func NewClaim(exp, userID int64) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
		UserID: userID,
	}
}

func (c *Claims) Valid() error {
	if !c.VerifyExpiresAt(time.Now().Unix(), true) {
		return errors.AuthError{Message: "token expired. log in again"}
	}

	if c.UserID == 0 {
		return errors.AuthError{Message: "user id not present in claims"}
	}

	return nil
}

func CreateToken(config configs.ConfigLoader, c *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)

	signedToken, err := token.SignedString([]byte(config.Get("ACCESS_KEY")))
	if err != nil {
		return "", errors.Error{Message: "error in signing token", Err: err, Type: "token-creation"}
	}

	return signedToken, nil
}

func ParseToken(config configs.ConfigLoader, signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, errors.AuthError{Message: "invalid signing algorithm"}
		}

		return []byte(config.Get("ACCESS_KEY")), nil
	})

	if err == nil && token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, errors.AuthError{Message: "invalid token", Err: err}
}
