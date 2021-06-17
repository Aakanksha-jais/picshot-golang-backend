package auth

import (
	"io/ioutil"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/dgrijalva/jwt-go"
)

// CreateToken parses rsa private key and uses it to generate a JWT token.
// It returns a JWT token in string format (header.payload.signature) that can be added to auth header.
func CreateToken(userID int64) (string, error) {
	const keyPath = "./configs/keys/id_rsa"

	signBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", errors.Error{Msg: "error in reading private key", Err: err, Type: "token-creation"}
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", errors.Error{Msg: "error in parsing private key", Err: err, Type: "token-creation"}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, newClaims(userID))
	token.Header["kid"] = "lVHu/0QjmU/ZFq8oxD9KYnDt6NA="

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", errors.Error{Err: err, Msg: "error in signing token"}
	}

	return signedToken, nil
}
