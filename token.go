package lib

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var TokenClaimsFetcher TokenClaimsFactory = tokenClaimsFactory{}

type TokenClaimsFactory interface {
	New(string) (*TokenClaims, error)
}

type tokenClaimsFactory struct{}

func (f tokenClaimsFactory) New(token string) (*TokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("Token contains invalid number of segments")
	}
	claimsSegment, err := DecodeSegment(parts[1])
	if err != nil {
		return nil, errors.New("Failed to decode token claims")
	}
	tokenClaims := &TokenClaims{}

	err = json.Unmarshal(claimsSegment, tokenClaims)
	if err != nil {
		return nil, errors.New("Malformed token")
	}
	return tokenClaims, nil
}

func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}

type TokenClaims struct {
	ID        string   `json:"jti,omitempty"`
	Subject   string   `json:"sub,omitempty"`
	Scopes    []string `json:"scope,omitempty"`
	ClientID  string   `json:"client_id,omitempty"`
	GrantType string   `json:"grant_type,omitempty"`
	UserID    string   `json:"user_id,omitempty"`
	Origin    string   `json:"origin,omitempty"`
	UserName  string   `json:"user_name,omitempty"`
	Email     string   `json:"email,omitempty"`
	AuthTime  int64    `json:"auth_time,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	ZoneID    string   `json:"zid,omitempty"`
	Audience  []string `json:"aud,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
}

func (tc TokenClaims) IsValid() error {
	now := time.Now().Unix()

	if tc.ExpiresAt <= 0 || now >= tc.ExpiresAt {
		return errors.New("Token is expired")
	}
	if tc.IssuedAt <= 0 || now < tc.IssuedAt {
		return errors.New("Token used before issued")
	}
	if tc.NotBefore <= 0 || now < tc.NotBefore {
		return errors.New("Token is not valid yet")
	}

	return nil
}
