package config

// WebauthnSettings defines the settings for the webauthn authentication mechanism
type WebauthnSettings struct {
	RelyingParty RelyingParty `yaml:"relying_party" json:"relying_party" koanf:"relying_party" split_words:"true"`
	Timeout      int          `yaml:"timeout" json:"timeout" koanf:"timeout"`
}

// Validate does not need to validate the config, because the library does this already
func (r *WebauthnSettings) Validate() error {
	return nil
}

// RelyingParty webauthn settings for your application.
type RelyingParty struct {
	Id          string `yaml:"id" json:"id" koanf:"id"`
	DisplayName string `yaml:"display_name" json:"display_name" koanf:"display_name" split_words:"true"`
	Icon        string `yaml:"icon" json:"icon" koanf:"icon"`
	// Deprecated: Use Origins instead
	Origin  string   `yaml:"origin" json:"origin" koanf:"origin"`
	Origins []string `yaml:"origins" json:"origins" koanf:"origins"`
}
