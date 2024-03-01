package entity

// ServiceStoreRequest is a request to remote service for data storage
type ServiceStoreRequest struct {
	AuthData MyAuthentication
	Record   Record
}

// NewServiceStoreRequest constructs the request
func NewServiceStoreRequest(
	auth MyAuthentication,
	rec Record,
) *ServiceStoreRequest {
	return &ServiceStoreRequest{
		AuthData: auth,
		Record:   rec,
	}
}
