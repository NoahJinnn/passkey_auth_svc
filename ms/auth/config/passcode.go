package config

type Email struct {
	FromAddress string `split_words:"true"`
	FromName    string `split_words:"true"`
}

type Passcode struct {
	OneSignalAppKey string
	OneSignalAppID  string
	Email           Email
	TTL             int32
}
