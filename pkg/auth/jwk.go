package auth

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"math/big"
)

// JWK (JSON web key) is a public key
type JWK struct {
	// Key ID Parameter; Refer to kid: https://datatracker.ietf.org/doc/html/rfc7517#section-4.5
	ID string `json:"kid"`

	// Algorithm Parameter; Refer to alg: https://datatracker.ietf.org/doc/html/rfc7517#section-4.4
	Algorithm string `json:"alg"`

	// Key Type Parameter; Refer to kty: https://datatracker.ietf.org/doc/html/rfc7517#section-4.1
	Type string `json:"kty"`

	// Public Key Use Parameter; Refer to use: https://datatracker.ietf.org/doc/html/rfc7517#section-4.2
	Use string `json:"use"`

	// Key Operations Parameter; Refer to key_ops: https://datatracker.ietf.org/doc/html/rfc7517#section-4.3
	Operations []string `json:"key_ops"`

	// rsa fields (the modulus n = pq, the public exponent e, the private exponent d, the two prime numbers p and q)
	Modulus         string `json:"n"`
	PublicExponent  string `json:"e"`
	PrivateExponent string `json:"d"`

	rsaPublicKey rsa.PublicKey
}

func (jwk *JWK) getRSAPublicKey() (rsa.PublicKey, error) {
	if jwk.rsaPublicKey.N != nil {
		return jwk.rsaPublicKey, nil
	}

	rsaPublicKey, err := generateRSAPublicKey(jwk)
	if err != nil {
		return jwk.rsaPublicKey, err
	}

	jwk.rsaPublicKey = rsaPublicKey

	return jwk.rsaPublicKey, nil
}

// generateRSAPublicKey takes JSON Web Key (JWK) and returns RSA Public Key (rsa.PublicKey)
func generateRSAPublicKey(key *JWK) (rsa.PublicKey, error) {
	var publicKey rsa.PublicKey

	// MODULUS
	decN, err := base64.RawURLEncoding.DecodeString(key.Modulus)
	if err != nil {
		return publicKey, err
	}

	n := big.NewInt(0)
	n.SetBytes(decN)

	// EXPONENT
	decE, err := base64.RawURLEncoding.DecodeString(key.PublicExponent)
	if err != nil {
		return publicKey, err
	}

	var eBytes []byte

	const DecStrLen = 8
	if len(decE) < DecStrLen {
		eBytes = make([]byte, DecStrLen-len(decE), DecStrLen)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}

	eReader := bytes.NewReader(eBytes)

	var e uint64

	err = binary.Read(eReader, binary.BigEndian, &e)
	if err != nil {
		return publicKey, err
	}

	publicKey.N = n
	publicKey.E = int(e)

	return publicKey, nil
}
