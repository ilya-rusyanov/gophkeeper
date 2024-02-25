package usercred

// UserCred is storage for app user's credentials
type UserCred struct{}

// New creates storage
func New(usernameFilename, appName string) *UserCred {
	// TODO: implement actual storage
	return &UserCred{}
}

func (c *UserCred) Store(login, password string) error {
	// TODO
	return nil
}
