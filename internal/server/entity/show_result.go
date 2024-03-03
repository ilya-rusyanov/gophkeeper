package entity

// ShowResult is result of 'show' (data revealing) operation
type ShowResult struct {
	Type    string
	Name    string
	Meta    string
	Payload []byte
}
