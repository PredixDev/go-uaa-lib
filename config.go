package lib

type Target interface {
	URL() string
	CaCertFile() string
	SkipSslVerify() bool
}

type Context interface {
	AccessToken() string
	TokenType() string
}
