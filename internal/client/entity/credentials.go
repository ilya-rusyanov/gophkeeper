package entity

// Credentials is auth data that user can store
type Credentials struct {
	// Name is title of the record
	Name string
	// Meta is additional information
	Meta []string
	// Login is credentials' login name
	Login string
	// Password is credentials' password
	Password string
}
