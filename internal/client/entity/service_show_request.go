package entity

// ServiceShowRequest is request to gateway to reveal user's data
type ServiceShowRequest struct {
	AuthData MyAuthentication
	Type     RecordType
	Name     string
}
