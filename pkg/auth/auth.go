package auth

import (
	"os"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/dgrijalva/jwt-go"
)

type JWTContextKey string

type Claims interface {
	Valid() error
	CreateToken() (string, error)
	ParseToken(signedToken string) error
	GetUserID() int64
}

type claims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

func New(exp, userID int64) *claims {
	return &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
		UserID: userID,
	}
}

func NewEmptyClaim() *claims {
	return &claims{}
}

func (c *claims) GetUserID() int64 {
	return c.UserID
}

func (c *claims) Valid() error {
	if !c.VerifyExpiresAt(time.Now().Unix(), true) {
		return errors.AuthError{Message: "token expired. log in again"}
	}

	if c.UserID == 0 {
		return errors.AuthError{Message: "user id not present in claims"}
	}

	return nil
}

func (c *claims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)

	signedToken, err := token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", errors.Error{Message: "error in signing token", Err: err, Type: "token-creation"}
	}

	return signedToken, nil
}

func (c *claims) ParseToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(signedToken, &claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, errors.AuthError{Message: "invalid signing algorithm"}
		}

		return []byte(os.Getenv("ACCESS_KEY")), nil
	})

	if err == nil && token != nil {
		if claims, ok := token.Claims.(*claims); ok && token.Valid {
			c = claims
			return nil
		}
	}

	return errors.AuthError{Message: "invalid token", Err: err}
}
