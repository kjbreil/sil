package sil

import "fmt"

// SIL is the structure of a SIL file
type SIL struct {
	Header    Header
	View      View
	Footer    Footer
	TableType interface{}
}

// Some Constants
const (
	crlf   = "\r\n"
	sqlInt = "INTEGER"
)

// Make makes a sil file of the definiton (as struct) passed
func Make(name string, definition interface{}) *SIL {
	s := new(SIL)
	// store the name of the table in the returned sil file
	s.View.Name = name
	return s
}

// Marshal creates the SIL structure from the information in the SIL type
func (s *SIL) Marshal() (data []byte, err error) {
	// Create the Header insert
	data = append(data, s.Header.insert()...)

	data = append(data, s.Header.row()...)

	data = append(data, s.View.bytes()...)

	return data, nil
}

func endLine() []byte {
	return []byte(fmt.Sprintf(";%s", crlf))
}
