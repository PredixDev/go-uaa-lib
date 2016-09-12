package lib_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lib "github.com/PredixDev/go-uaa-lib"
	. "github.com/onsi/gomega"
)

func TestInfoServerSuccess(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Expect(r.URL.RequestURI()).To(Equal("/login"))
		Expect(r.Header.Get("Accept")).To(Equal("application/json"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	info := lib.NewInfoClient(server.URL, false, "")

	Expect(info.Server()).To(BeNil())
}

func TestInfoServerError(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	info := lib.NewInfoClient(server.URL, false, "")

	err := info.Server()
	Expect(err.Error()).To(Equal("Invalid status response: 500"))
}

func TestInfoServerInvalidURL(t *testing.T) {
	RegisterTestingT(t)

	file, _ := ioutil.TempFile(os.TempDir(), "test-cert-file")
	defer os.Remove(file.Name())

	_, _ = file.WriteString(sampleCertificate)
	_ = file.Close()

	info := lib.NewInfoClient("Invalid URL", false, file.Name())

	err := info.Server()
	Expect(err).ToNot(BeNil())
}

var sampleCertificate = `-----BEGIN CERTIFICATE-----
MIIGOTCCBCGgAwIBAgIJAOE/vJd8EB24MA0GCSqGSIb3DQEBBQUAMIGyMQswCQYD
VQQGEwJGUjEPMA0GA1UECAwGQWxzYWNlMRMwEQYDVQQHDApTdHJhc2JvdXJnMRgw
FgYDVQQKDA93d3cuZnJlZWxhbi5vcmcxEDAOBgNVBAsMB2ZyZWVsYW4xLTArBgNV
BAMMJEZyZWVsYW4gU2FtcGxlIENlcnRpZmljYXRlIEF1dGhvcml0eTEiMCAGCSqG
SIb3DQEJARYTY29udGFjdEBmcmVlbGFuLm9yZzAeFw0xMjA0MjcxMDE3NDRaFw0x
MjA1MjcxMDE3NDRaMIGyMQswCQYDVQQGEwJGUjEPMA0GA1UECAwGQWxzYWNlMRMw
EQYDVQQHDApTdHJhc2JvdXJnMRgwFgYDVQQKDA93d3cuZnJlZWxhbi5vcmcxEDAO
BgNVBAsMB2ZyZWVsYW4xLTArBgNVBAMMJEZyZWVsYW4gU2FtcGxlIENlcnRpZmlj
YXRlIEF1dGhvcml0eTEiMCAGCSqGSIb3DQEJARYTY29udGFjdEBmcmVlbGFuLm9y
ZzCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAODp+8oQcK+MTuWPZVxJ
ZR75paK4zcUngupYXWSGWFXPTV7vssFk6vInePArTL+T9KwHfiZ29Pp3UbzDlysY
Kz9f9Ae50jGD6xVPwXgQ/VI979GyFXzhiEMtSYykF04tBJiDl2/FZxbHPpNxC39t
14kwuDqBin9N/ZbT5+45tbbS8ziXS+QgL5hD2q2eYCWayrGEt1Y+jDAdHDHmGnZ8
d4hbgILJAs3IInOCDjC4c1gwHFb8G4QHHTwVhjhqpkq2hQHgzWBC1l2Dku/oDYev
Zu/pfpTo3z6+NOYBrUWseQmIuG+DGMQA9KOuSQveyTywBm4G4vZKn0sCu1/v2+9T
BGv41tgS/Yf6oeeQVrbS4RFY1r9qTK6DW9wkTTesa4xoDKQrWjSJ7+aa8tvBXLGX
x2xdRNWLeRMuGBSOihwXmDr+rCJRauT7pItN5X+uWNTX1ofNksQSUMaFJ5K7L0LU
iQqU2Yyt/8UphdVZL4EFkGSA13UDWtb9mM1hY0h65LlSYwCchEphrtI9cuV+ITrS
NcN6cP/dqDx1/jWd6dqjNu7+dugwX5elQS9uUYCFmugR5s1m2eeBg3QuC7gZLE0N
NbgS7oSxKJe9KeOcw68jHWfBKsCfBfQ4fU2t/ntMybT3hCdEMQu4dgM5Tyw/UeFq
0SaJyTl+G1bTzS0FW6uUp6NLAgMBAAGjUDBOMB0GA1UdDgQWBBQjbC09PildeLhs
Pqriuy4ebIfyUzAfBgNVHSMEGDAWgBQjbC09PildeLhsPqriuy4ebIfyUzAMBgNV
HRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4ICAQCwRJpJCgp7S+k9BT6X3kBefonE
EOYtyWXBPpuyG3Qlm1rdhc66DCGForDmTxjMmHYtNmAVnM37ILW7MoflWrAkaY19
gv88Fzwa5e6rWK4fTSpiEOc5WB2A3HPN9wJnhQXt1WWMDD7jJSLxLIwFqkzpDbDE
9122TtnIbmKNv0UQpzPV3Ygbqojy6eZHUOT05NaOT7vviv5QwMAH5WeRfiCys8CG
Sno/o830OniEHvePTYswLlX22LyfSHeoTQCCI8pocytl7IwARKCvBgeFqvPrMiqP
ch16FiU9II8KaMgpebrUSz3J1BApOOd1LBd42BeTAkNSxjRvbh8/lDWfnE7ODbKc
b6Ad3V9flFb5OBZH4aTi6QfrDnBmbLgLL8o/MLM+d3Kg94XRU9LjC2rjivQ6MC53
EnWNobcJFY+soXsJokGtFxKgIx8XrhF5GOsT2f1pmMlYL4cjlU0uWkPOOkhq8tIp
R8cBYphzXu1v6h2AaZLRq184e30ZO98omKyQoQ2KAm5AZayRrZZtjvEZPNamSuVQ
iPe3o/4tyQGq+jEMAEjLlDECu0dEa6RFntcbBPMBP3wZwE2bI9GYgvyaZd63DNdm
Xd65m0mmfOWYttfrDT3Q95YP54nHpIxKBw1eFOzrnXOqbKVmJ/1FDP2yWeooKVLf
KvbxUcDaVvXB0EU0bg==
-----END CERTIFICATE-----
`
