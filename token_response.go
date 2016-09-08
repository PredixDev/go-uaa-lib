package lib

type TokenResponse struct {
	Type    string `json:"token_type"`
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func NewTokenResponse() *TokenResponse {
	return &TokenResponse{}
}

func (tr TokenResponse) IsPresent() bool {
	return tr.Access != "" && tr.Refresh != ""
}

func (tr TokenResponse) IsValid() bool {
	if !tr.IsPresent() {
		return false
	}
	tc, err := TokenClaimsFetcher.New(tr.Access)
	if tc == nil || err != nil {
		return false
	}
	return tc.IsValid() == nil
}

func (tr TokenResponse) GetAccessToken() string {
	return tr.Type + " " + tr.Access
}
