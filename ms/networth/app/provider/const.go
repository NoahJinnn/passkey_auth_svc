package provider

const (
	Finverse = "finverse"
	Manual   = "manual"
)

func ValidateProvider(input string) bool {
	switch input {
	case Finverse:
		return true
	case Manual:
		return true
	}

	return false
}
