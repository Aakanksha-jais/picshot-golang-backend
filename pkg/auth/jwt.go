package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
)

type JWTContextKey string

type JWT struct {
	payload   string
	header    header
	signature string
	token     string
}

type header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	URL       string `json:"jku"`
	KeyID     string `json:"kid"`
}

func getJWT(r *http.Request) (JWT, error) {
	authHeader := r.Header.Get("Authorization")
	fields := strings.Fields(authHeader)

	if authHeader == "" || len(fields) != 2 || !strings.EqualFold("bearer", fields[0]) {
		return JWT{}, errors.AuthError{Msg: fmt.Sprintf("invalid auth-header: %v", authHeader)}
	}

	jwtParts := strings.Split(fields[1], ".")
	if len(jwtParts) != 3 {
		return JWT{}, errors.AuthError{Msg: fmt.Sprintf("auth-header not in the format [header.payload.signature] : %v", authHeader)}
	}

	var h header

	decodedHeader, err := base64.RawStdEncoding.DecodeString(jwtParts[0])
	if err != nil {
		return JWT{}, errors.AuthError{Msg: fmt.Sprintf("cannot decode jwt-header: %v", err)}
	}

	err = json.Unmarshal(decodedHeader, &h)
	if err != nil {
		return JWT{}, errors.AuthError{Msg: fmt.Sprintf("error unmarshalling the decoded jwt-header: %v", err)}
	}

	return JWT{header: h, payload: jwtParts[1], signature: jwtParts[2], token: fields[1]}, err
}
