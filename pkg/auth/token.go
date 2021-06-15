package auth

import (
	"fmt"
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

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return signedToken, nil
}