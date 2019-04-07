package sil

// View is the data of the SIL file
// Name is the table name
// Data is an array of interfaces
type View struct {
	Name string
	Data []interface{}
}

func (v *View) bytes() (b []byte) {
	for i := range v.Data {
		row := row(v.Data[i])
		b = append(b, row...)
		b = append(b, []byte(crlf)...)
	}
	// remove the last CRLF, 2 bytes
	b = b[:len(b)-2]
	// append the endline code (; + crlf)
	b = append(b, endLine()...)
	return
}
