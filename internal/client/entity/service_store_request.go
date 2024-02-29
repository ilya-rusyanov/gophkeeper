package entity

// ServiceStoreRequest is a request to remote service for data storage
type ServiceStoreRequest struct {
	UserCredentials MyCredentials
	Record          Record
}

// NewServiceStoreRequest constructs the request
func NewServiceStoreRequest(
	cred MyCredentials,
	rec Record,
) *ServiceStoreRequest {
	return &ServiceStoreRequest{
		UserCredentials: cred,
		Record:          rec,
	}
}
