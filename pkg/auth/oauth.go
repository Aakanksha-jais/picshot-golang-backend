package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
)

type OAuth struct {
	options Options
	cache   PublicKeyCache
}

type Options struct {
	ValidityFrequency int
	JWKPath           string
}

type PublicKeyCache struct {
	publicKeys JWKS
	mu         sync.RWMutex
}

func New(options Options) *OAuth {
	oauth := &OAuth{
		options: options,
		cache: PublicKeyCache{
			publicKeys: JWKS{},
			mu:         sync.RWMutex{},
		},
	}

	_ = oauth.invalidateCache()

	return oauth
}

func (o *OAuth) Validate(r *http.Request) (*jwt.Token, error) {
	token := &jwt.Token{Valid: false}

	// get JSON web token
	jwtObj, err := getJWT(r)
	if err != nil {
		return token, err
	}

	// fetch public key for specified header key id
	publicKey := o.cache.publicKeys.Get(jwtObj.header.KeyID)

	pKey, err := publicKey.getRSAPublicKey()
	if err != nil {
		return nil, err
	}

	token, err = jwt.ParseWithClaims(
		jwtObj.token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) { // keyFunc will receive the parsed token and should return the key for validating.
			if t.Method.Alg() != jwt.SigningMethodRS256.Alg() {
				return nil, errors.AuthError{Msg: "invalid signing algorithm"}
			}

			return &pKey, nil
		},
	)

	// ParseWithClaims calls Valid() to check if the token is valid
	if err != nil || !token.Valid {
		return nil, errors.AuthError{Msg: "invalid token", Err: err} // todo: recheck
	}

	return token, nil
}

func (o *OAuth) invalidateCache() error {
	var err error

	o.cache.mu.Lock()
	var keys []JWK

	duration := o.options.ValidityFrequency
	if keys, err = o.loadJWK(); err != nil {
		duration = 3
	} else {
		// save the public keys in memory
		o.cache.publicKeys.Keys = keys
	}

	if duration > 0 {
		go func() {
			time.Sleep(time.Duration(duration) * time.Second)

			_ = o.invalidateCache()
		}()
	}
	o.cache.mu.Unlock()

	return err
}

func (o *OAuth) loadJWK() ([]JWK, error) {
	// if key is not present in memory, get it from jwk_endpoint
	resp, err := http.Get(o.options.JWKPath)
	if err != nil {
		return nil, errors.Error{Msg: "failed to fetch the response from jwks endpoint", Err: err}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Error{Err: err, Msg: "error in reading response from JWKS endpoint"}
	}

	var jwks JWKS

	err = json.Unmarshal(body, &jwks)
	if err != nil {
		return nil, errors.Error{Err: err, Msg: "error in unmarshalling response from JWKS endpoint"}
	}

	return jwks.Keys, nil
}
