package entity

// AuthPayload is payload for authentication data
type AuthPayload struct {
	// Login is credentials' login name
	Login string
	// Password is credentials' password
	Password string
}

// NewAuthPayload constructs authentication data payload
func NewAuthPayload(login string, password string) AuthPayload {
	return AuthPayload{
		Login:    login,
		Password: password,
	}
}
