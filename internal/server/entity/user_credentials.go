package entity

// UserCredentials represents authorization data for user
type UserCredentials struct {
	Login    string
	Password string
}

// NewUserCredentials constructs user authorization data
func NewUserCredentials(login, password string) *UserCredentials {
	return &UserCredentials{
		Login:    login,
		Password: password,
	}
}
