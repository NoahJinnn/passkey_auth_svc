package config

type Emails struct {
	RequireVerification bool `yaml:"require_verification" json:"require_verification" koanf:"require_verification" split_words:"true"`
	MaxNumOfAddresses   int  `yaml:"max_num_of_addresses" json:"max_num_of_addresses" koanf:"max_num_of_addresses" split_words:"true"`
}
