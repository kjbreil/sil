package sil

func (s *SIL) checkHeader() error {
	return nil
}

// AddRplDCT Creates and returns the DCT information
// Needs to be replaced
func (s *SIL) AddRplDCT() {
	s.ViewHeader.F902 = f902 // Batch identifier
	s.ViewHeader.F903 = f903 // Batch creator
	s.ViewHeader.F901 = "HM" // Batch type
	s.ViewHeader.F910 = f910
	s.ViewHeader.F904 = f904 // Batch destination
	s.ViewHeader.F909 = f909
	s.ViewHeader.F912 = "ADDRPL"
	s.ViewHeader.F913 = "ADDRPL CHANGED OPERATORS"
}
