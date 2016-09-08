package lib_test

import (
	"testing"
	"time"

	lib "github.com/PredixDev/go-uaa-lib"
	. "github.com/onsi/gomega"
)

func TestGetTokenClaimsWithLessThan3Parts(t *testing.T) {
	RegisterTestingT(t)

	claims, err := lib.TokenClaimsFetcher.New("part1.part2")
	Expect(claims).To(BeNil())
	Expect(err.Error()).To(Equal("Token contains invalid number of segments"))
}

func TestGetTokenClaimsWithMoreThan3Parts(t *testing.T) {
	RegisterTestingT(t)

	claims, err := lib.TokenClaimsFetcher.New("part1.part2.part3.part4")
	Expect(claims).To(BeNil())
	Expect(err.Error()).To(Equal("Token contains invalid number of segments"))
}

func TestGetTokenClaimsWithInvalidEncoding(t *testing.T) {
	RegisterTestingT(t)

	claims, err := lib.TokenClaimsFetcher.New("encoded_part1.encoded_part2.encoded_part3")
	Expect(claims).To(BeNil())
	Expect(err.Error()).To(Equal("Failed to decode token claims"))
}

func TestGetTokenClaimsWithMalformedToken(t *testing.T) {
	RegisterTestingT(t)

	claims, err := lib.TokenClaimsFetcher.New("encoded_part1.ZW5jb2RlZF9wYXJ0Mg==.encoded_part3")
	Expect(claims).To(BeNil())
	Expect(err.Error()).To(Equal("Malformed token"))
}

func TestGetTokenClaimsWithValidToken(t *testing.T) {
	RegisterTestingT(t)

	claims, err := lib.TokenClaimsFetcher.New("encoded_part1.ew0KICAianRpIjogIjEyMzQ1Njc4OTAiLA0KICAic3ViIjogIkFCQ0QiLA0KICAic2NvcGUiOiBbDQogICAgIm9wZW5pZCINCiAgXSwNCiAgImNsaWVudF9pZCI6ICJzb21lLWNsaWVudCIsDQogICJjaWQiOiAic29tZS1jbGllbnQiLA0KICAiYXpwIjogInNvbWUtY2xpZW50IiwNCiAgImdyYW50X3R5cGUiOiAicGFzc3dvcmQiLA0KICAidXNlcl9pZCI6ICJBQkNEIiwNCiAgIm9yaWdpbiI6ICJ1YWEiLA0KICAidXNlcl9uYW1lIjogInNvbWUtdXNlciIsDQogICJlbWFpbCI6ICJzb21lQHVzZXIuY29tIiwNCiAgImF1dGhfdGltZSI6IDE0NzMyMDUzMDUsDQogICJyZXZfc2lnIjogIjEyMzQiLA0KICAiaWF0IjogMTQ3MzIwNTMwNSwNCiAgImV4cCI6IDE0NzMyNDg1MDUsDQogICJpc3MiOiAiaHR0cHM6Ly9YWVoudWFhLmNvbS9vYXV0aC90b2tlbiIsDQogICJ6aWQiOiAiWFlaIiwNCiAgImF1ZCI6IFsNCiAgICAib3BlbmlkIiwNCiAgICAic29tZS1jbGllbnQiDQogIF0NCn0=.encoded_part3")
	Expect(claims).ToNot(BeNil())
	Expect(err).To(BeNil())

	Expect(claims.ID).To(Equal("1234567890"))
	Expect(claims.Subject).To(Equal("ABCD"))
	Expect(claims.Scopes).To(ConsistOf("openid"))
	Expect(claims.ClientID).To(Equal("some-client"))
	Expect(claims.GrantType).To(Equal("password"))
	Expect(claims.UserID).To(Equal("ABCD"))
	Expect(claims.Origin).To(Equal("uaa"))
	Expect(claims.UserName).To(Equal("some-user"))
	Expect(claims.Email).To(Equal("some@user.com"))
	Expect(claims.AuthTime).To(Equal(int64(1473205305)))
	Expect(claims.IssuedAt).To(Equal(int64(1473205305)))
	Expect(claims.ExpiresAt).To(Equal(int64(1473248505)))
	Expect(claims.Issuer).To(Equal("https://XYZ.uaa.com/oauth/token"))
	Expect(claims.ZoneID).To(Equal("XYZ"))
	Expect(claims.Audience).To(ConsistOf("openid", "some-client"))
}

func TestTokenIsValidExpiresAt(t *testing.T) {
	RegisterTestingT(t)

	tc := lib.TokenClaims{}
	tc.ExpiresAt = -1
	Expect(tc.IsValid().Error()).To(Equal("Token is expired"))

	tc = lib.TokenClaims{}
	tc.ExpiresAt = time.Now().Unix() - 1
	Expect(tc.IsValid().Error()).To(Equal("Token is expired"))
}

func TestTokenIsValidIssuedAt(t *testing.T) {
	RegisterTestingT(t)

	tc := lib.TokenClaims{}
	tc.ExpiresAt = time.Now().Unix() + 100000
	tc.IssuedAt = -1
	Expect(tc.IsValid().Error()).To(Equal("Token used before issued"))

	tc = lib.TokenClaims{}
	tc.ExpiresAt = time.Now().Unix() + 100000
	tc.IssuedAt = time.Now().Unix() + 10000
	Expect(tc.IsValid().Error()).To(Equal("Token used before issued"))
}

func TestTokenIsValidNotBefore(t *testing.T) {
	RegisterTestingT(t)

	tc := lib.TokenClaims{}
	tc.ExpiresAt = time.Now().Unix() + 100000
	tc.IssuedAt = time.Now().Unix() - 1
	tc.NotBefore = -1
	Expect(tc.IsValid().Error()).To(Equal("Token is not valid yet"))

	tc = lib.TokenClaims{}
	tc.ExpiresAt = time.Now().Unix() + 100000
	tc.IssuedAt = time.Now().Unix() - 1
	tc.NotBefore = time.Now().Unix() + 100000
	Expect(tc.IsValid().Error()).To(Equal("Token is not valid yet"))
}

func TestTokenIsValid(t *testing.T) {
	RegisterTestingT(t)

	tc := lib.TokenClaims{}
	tc.ExpiresAt = time.Now().Unix() + 100000
	tc.IssuedAt = time.Now().Unix() - 1
	tc.NotBefore = time.Now().Unix() - 1
	Expect(tc.IsValid()).To(BeNil())
}
