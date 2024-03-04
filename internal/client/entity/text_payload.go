package entity

// TextPayload is payload for authentication data
type TextPayload string

// NewTextPayload constructs text data payload
func NewTextPayload(t string) TextPayload {
	return TextPayload(t)
}
