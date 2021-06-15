package auth

// JWKS (JSON web key set) is public key set. Refer to JWK Set: https://datatracker.ietf.org/doc/html/rfc7517#section-5
type JWKS struct {
	// Keys Prameter; Refer to keys: https://datatracker.ietf.org/doc/html/rfc7517#section-5.1
	Keys []JWK `json:"keys"`
}

// Get fetches the Public Key corresponding to given KeyID
func (jwks *JWKS) Get(kID string) *JWK {
	for _, key := range jwks.Keys {
		if key.ID == kID {
			return &key
		}
	}

	return &JWK{}
}
