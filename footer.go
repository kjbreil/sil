package sil

// Footer is an array of strings to be apended to end of file
type Footer []string

// Append adds a line to the bottom of the SIL file
func (s *SIL) Append(str string) {
	s.Footer = append(s.Footer, str)
}
