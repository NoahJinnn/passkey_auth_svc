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
	ID              uint
	AccountInfo     struct{}
	InstitutionInfo struct{}
	AssetInfo       struct{}
	SensibleData    string
	Descriptions    string
}

type BankAccount struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type Car struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type Collectible struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type CryptoAccount struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type Loan struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}

type PrivateShare struct {
	ID          uint
	UserID      uint
	AssetInfoID uint
}
