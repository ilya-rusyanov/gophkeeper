package gophkeepergw

// GophKeeperGW is a gateway to the actual service
type GophKeeperGW struct{}

// New creates an instance of the gateway
func New(serverAddr string) *GophKeeperGW {
	// TODO: implement actual constructor
	return &GophKeeperGW{}
}

func (gk *GophKeeperGW) Register(login, password string) error {
	// TODO
	return nil
}
