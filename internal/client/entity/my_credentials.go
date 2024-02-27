package entity

// MyCredentials represents user's login credentials
type MyCredentials struct {
	Login    string
	Password string
}

// NewMyCredentials constructs new user credentials
func NewMyCredentials(login, password string) *MyCredentials {
	return &MyCredentials{
		Login:    login,
		Password: password,
	}
}
