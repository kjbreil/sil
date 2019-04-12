package sil

// View is the data of the SIL file
// Name is the table name
// Data is an array of interfaces
type View struct {
	Name     string
	Required bool
	Data     []interface{}
}
