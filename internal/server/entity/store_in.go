package entity

// StoreIn is carrier of store use case arguments
type StoreIn struct {
	// Login is user's login
	Login string
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
	login string,
	typ,
	name,
	meta string,
	payload []byte,
) *StoreIn {
	return &StoreIn{
		Login:   login,
		Type:    typ,
		Name:    name,
		Meta:    meta,
		Payload: payload,
	}
}
