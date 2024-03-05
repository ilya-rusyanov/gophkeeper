package entity

// RecordType is type for defining kind of record
type RecordType string

const (
	// RecordTypeAuth is authentication data
	RecordTypeAuth RecordType = "auth"
	// RecordTypeText is text data
	RecordTypeText RecordType = "text"
	// RecordTypeBin is binary data
	RecordTypeBin RecordType = "bin"
	// RecordTypeCard is card data
	RecordTypeCard RecordType = "card"
)
