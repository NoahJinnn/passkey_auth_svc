package dao

type Provider string

const (
	Finverse Provider = "finverse"
	Manual   Provider = "manual"
)

func ValidateProvider(input Provider) bool {
	switch input {
	case Finverse:
		return true
	case Manual:
		return true
	}

	return false
}

func (p Provider) String() string {
	return string(p)
}

type Category string

const (
	Asset Category = "asset"
	Debt  Category = "debt"
)

func ValidateCategory(input Category) bool {
	switch input {
	case Asset:
		return true
	case Debt:
		return true
	}

	return false
}

func (p Category) String() string {
	return string(p)
}
