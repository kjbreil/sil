package sil

import "strings"

// Footer is an array of strings to be apended to end of file
type Footer []string

// Append adds a line to the bottom of the SIL file
func (s *SIL) Append(str string) {
	s.Footer = append(s.Footer, str)
	s.Footer = append(s.Footer, crlf)
}

func (f *Footer) bytes() []byte {
	return []byte(strings.Join(*f, ""))
}
