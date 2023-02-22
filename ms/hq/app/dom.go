package app

type User struct {
	ID          uint
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
}

type AssetInfo struct {
	ID uint `json:"id,omitempty"`
	// AccountInfo holds the value of the "account_info" field.
	AccountInfo struct{} `json:"account_info,omitempty"`
	// InstitutionInfo holds the value of the "institution_info" field.
	InstitutionInfo struct{} `json:"institution_info,omitempty"`
	// AssetInfo holds the value of the "asset_info" field.
	AssetInfo struct{} `json:"asset_info,omitempty"`
	// SensibleData holds the value of the "sensible_data" field.
	SensibleData string `json:"sensible_data,omitempty"`
	// Descriptions holds the value of the "descriptions" field.
	Descriptions string `json:"descriptions,omitempty"`
}

type BankAccount struct {
	ID              uint
	UserID          uint
	AccountID       string
	InstitutionInfo struct{} `json:"institution_info,omitempty"`
	// AccountInfo holds the value of the "account_info" field.
	AccountInfo struct{} `json:"account_info,omitempty"`
	// SensibleData holds the value of the "sensible_data" field.
	SensibleData string `json:"sensible_data,omitempty"`
}

type Car struct {
	ID     uint
	UserID uint
	Make   string
	Model  string
	Year   int32
}

type Collectible struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
}

type CryptoAccount struct {
	ID       uint
	UserID   uint
	Name     string
	CoinType string
	Balance  float64
}
