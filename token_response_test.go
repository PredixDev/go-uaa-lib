package lib_test

import (
	"fmt"
	"testing"
	"time"

	lib "github.com/PredixDev/go-uaa-lib"
	. "github.com/onsi/gomega"
)

func TestTokenResponseIsPresent(t *testing.T) {
	RegisterTestingT(t)

	tr := lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = "Refresh Token"
	Expect(tr.IsPresent()).To(BeTrue())

	tr = lib.NewTokenResponse()
	tr.Access = ""
	tr.Refresh = "Refresh Token"
	Expect(tr.IsPresent()).To(BeFalse())

	tr = lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = ""
	Expect(tr.IsPresent()).To(BeFalse())

	tr = lib.NewTokenResponse()
	tr.Access = ""
	tr.Refresh = ""
	Expect(tr.IsPresent()).To(BeFalse())
}

func TestTokenResponseIsValid(t *testing.T) {
	RegisterTestingT(t)

	var oldTokenClaimsFactory = lib.TokenClaimsFactory
	defer func() { lib.TokenClaimsFactory = oldTokenClaimsFactory }()

	tr := lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = ""
	Expect(tr.IsValid()).To(BeFalse())

	tr = lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = "Refresh Token"
	lib.TokenClaimsFactory = fakeTokenClaimsFactory{
		tc:  nil,
		err: nil,
	}
	Expect(tr.IsValid()).To(BeFalse())

	tr = lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = "Refresh Token"
	lib.TokenClaimsFactory = fakeTokenClaimsFactory{
		tc:  nil,
		err: fmt.Errorf("Error"),
	}
	Expect(tr.IsValid()).To(BeFalse())

	tr = lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = "Refresh Token"
	lib.TokenClaimsFactory = fakeTokenClaimsFactory{
		tc:  &lib.TokenClaims{},
		err: fmt.Errorf("Error"),
	}
	Expect(tr.IsValid()).To(BeFalse())

	tr = lib.NewTokenResponse()
	tr.Access = "Access Token"
	tr.Refresh = "Refresh Token"
	lib.TokenClaimsFactory = fakeTokenClaimsFactory{
		tc: &lib.TokenClaims{
			ExpiresAt: time.Now().Unix() + 100000,
			IssuedAt:  time.Now().Unix() - 1,
			NotBefore: time.Now().Unix() - 1,
		},
		err: nil,
	}
	Expect(tr.IsValid()).To(BeTrue())
}

func TestTokenResponseGetAccessToken(t *testing.T) {
	RegisterTestingT(t)

	tr := lib.NewTokenResponse()
	tr.Type = "Token_Type"
	tr.Access = "Access_Token"
	Expect(tr.GetAccessToken()).To(Equal("Token_Type Access_Token"))
}

type fakeTokenClaimsFactory struct {
	tc  *lib.TokenClaims
	err error
}

func (f fakeTokenClaimsFactory) New(token string) (*lib.TokenClaims, error) {
	return f.tc, f.err
}
