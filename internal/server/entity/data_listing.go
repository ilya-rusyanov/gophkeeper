package entity

// DataListing is the result of listing data operation
type DataListing []DataListEntry

// DataListEntry is single entry of data listing operation
type DataListEntry struct {
	Type string
	Name string
}

// NewDataListEntry creates DataListEntry
func NewDataListEntry(typ string, name string) DataListEntry {
	return DataListEntry{
		Type: typ,
		Name: name,
	}
}
