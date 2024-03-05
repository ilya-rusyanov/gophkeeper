package entity

type Meta []string

// Record is piece of data that user can store
type Record struct {
	// Type is type of the record
	Type RecordType
	// Name is title of the record
	Name string
	// Meta is additional information
	Meta Meta
	// Payload is variable part of the record
	Payload any
}

// NewAuthRecord constructs authentication data record
func NewAuthRecord(
	name string,
	meta Meta,
	data AuthPayload,
) *Record {
	return &Record{
		Type:    RecordTypeAuth,
		Name:    name,
		Meta:    meta,
		Payload: &data,
	}
}

// NewTextRecord constructs text data record
func NewTextRecord(
	name string,
	meta Meta,
	data TextPayload,
) *Record {
	return &Record{
		Type:    RecordTypeText,
		Name:    name,
		Meta:    meta,
		Payload: &data,
	}
}

// NewTextRecord constructs text data record
func NewBinRecord(
	name string,
	meta Meta,
	data BinPayload,
) *Record {
	return &Record{
		Type:    RecordTypeBin,
		Name:    name,
		Meta:    meta,
		Payload: &data,
	}
}
