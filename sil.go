package sil

// SIL is the structure of a SIL file
type SIL struct {
	Header    Header
	View      View
	Footer    Footer
	TableType interface{}
}

// View holds the data
type View struct {
	Name    string
	Columns []Column
	Data    []interface{}
}

// Column is each column in a SIL file containing both the name and the type
// contained.
type Column struct {
	Name string
	Type string
}

// Some Constants
const (
	crlf   = "\r\n"
	sqlInt = "INTEGER"
)

// Make makes a sil file of the definiton (as struct) passed
func Make(name string, definition interface{}) (s *SIL) {
	// store the name of the table in the returned sil file
	s.View.Name = name
	return s
}

// Bytes creates the SIL structure from the information in the SIL type
func (s *SIL) Bytes() (data []byte, err error) {

	return data, nil
}

// String returns a string of a SIL file, wrapper for Bytes()
func (s *SIL) String() (string, error) {
	b, e := s.Bytes()
	return string(b), e
}
