package sil

// check and fix a header with default tag
func (h *Header) check() {
	fields, values := fieldValue(h)

	forFields(fields, values)
}

// AddRplDCT Creates and returns the DCT information
// Needs to be replaced
func (s *SIL) AddRplDCT() {
	s.Header.check()
}
