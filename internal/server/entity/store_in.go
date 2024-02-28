package entity

// StoreIn is carrier of store use case arguments
type StoreIn struct {
	// Credentials depicts credentials of the user who submits store request
	Credentials UserCredentials
	// Type is the type of data being submitted
	Type string
	// Name is the title of the data that is being submitted
	Name string
	// Meta is additional information about data
	Meta string
	// Payload is the actual data being submitted
	Payload []byte
}

// NewStoreIn creates and initializes StoreIn
func NewStoreIn(
	creds UserCredentials,
	typ,
	name,
	meta string,
	payload []byte,
) *StoreIn {
	return &StoreIn{
		Credentials: creds,
		Type:        typ,
		Name:        name,
		Meta:        meta,
		Payload:     payload,
	}
}
