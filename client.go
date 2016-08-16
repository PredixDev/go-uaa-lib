package lib

type Client struct {
	ID                  string   `json:"client_id,omitempty"`
	Secret              string   `json:"client_secret,omitempty"`
	Name                string   `json:"name,omitempty"`
	Scopes              []string `json:"scope,omitempty"`
	ResourceIDs         []string `json:"resource_ids,omitempty"`
	GrantTypes          []string `json:"authorized_grant_types,omitempty"`
	Authorities         []string `json:"authorities,omitempty"`
	AutoApprove         []string `json:"autoapprove,omitempty"`
	RedirectURI         []string `json:"redirect_uri,omitempty"`
	AccessTokenTimeout  int      `json:"access_token_validity,omitempty"`
	RefreshTokenTimeout int      `json:"refresh_token_validity,omitempty"`
	SignupRedirect      string   `json:"signup_redirect_url,omitempty"`
	Action              string   `json:"action,omitempty"`
}

type ClientResources struct {
	Clients []Client `json:"resources,omitempty"`
}
