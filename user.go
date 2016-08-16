package lib

type Name struct {
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
}

type Value struct {
	Value   string `json:"value,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type Meta struct {
	Version      int    `json:"version,omitempty"`
	Created      string `json:"created,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
}

type User struct {
	ID                   string   `json:"id"`
	Meta                 Meta     `json:"meta"`
	UserName             string   `json:"userName,omitempty"`
	Password             string   `json:"password,omitempty"`
	Name                 Name     `json:"name,omitempty"`
	Emails               []Value  `json:"emails,omitempty"`
	Groups               []string `json:"groups,omitempty"`
	Approval             []string `json:"approval,omitempty"`
	Phones               []Value  `json:"phoneNumbers,omitempty"`
	Active               bool     `json:"active,omitempty"`
	Verified             bool     `json:"verified,omitempty"`
	Origin               string   `json:"origin,omitempty"`
	ZoneID               string   `json:"zoneId,omitempty"`
	PasswordLastModified string   `json:"passwordLastModified,omitempty"`
}

type UserResources struct {
	Users []User `json:"resources,omitempty"`
}
