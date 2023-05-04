package config

// WebauthnSettings defines the settings for the webauthn authentication mechanism
type WebauthnSettings struct {
	RelyingParty RelyingParty `split_words:"true"`
	Timeout      int
}

// Validate does not need to validate the config, because the library does this already
func (r *WebauthnSettings) Validate() error {
	return nil
}

// RelyingParty webauthn settings for your application.
type RelyingParty struct {
	Id          string
	DisplayName string `split_words:"true"`
	Icon        string
	// Deprecated: Use Origins instead
	Origin  string
	Origins []string
}
