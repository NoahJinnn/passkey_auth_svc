package config

// WebauthnSettings defines the settings for the webauthn authentication mechanism
type WebauthnSettings struct {
	RelyingParty RelyingParty `split_words:"true"`
	Timeout      int
}

// RelyingParty webauthn settings for your application.
type RelyingParty struct {
	Id          string
	DisplayName string `split_words:"true"`
	Icon        string
	Origins     []string
}
