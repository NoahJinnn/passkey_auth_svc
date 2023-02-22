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

type BankAccount struct {
	ID              uint
	UserID          uint
	AccountID       string
	InstitutionName string
	AccountType     string
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
