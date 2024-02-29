package entity

// MyAuthentication holds user's own authentication data
type MyAuthentication string

// NewMyAuthentication constructs MyAuthentication
func NewMyAuthentication(token string) MyAuthentication {
	return MyAuthentication(token)
}
