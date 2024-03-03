package entity

// DataList is data listing result
type DataList []DataListEntry

// DataListEntry is entry of data listing result
type DataListEntry struct {
	Type string
	Name string
}

// NewDataListEntry creates one data list entry
func NewDataListEntry(typ string, name string) DataListEntry {
	return DataListEntry{
		Type: typ,
		Name: name,
	}
}
