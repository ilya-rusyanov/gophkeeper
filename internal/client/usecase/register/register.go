package register

// CredentialsStorager represents storage of credentials
type CredentialsStorager interface{}

// Servicer is remote service
type Servicer interface{}

// Register is UC for user registration on server
type Register struct{}

// New constructs UC
func New(
	credStorage CredentialsStorager,
	service Servicer,
) *Register {
	// TODO: implement constructor logic
	return &Register{}
}

// Register performs user registration
func (r *Register) Register(login, password string) error {
	// TODO: implement logic
	return nil
}
