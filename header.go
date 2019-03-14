package sil

// check and fix a header with default tag
func (h *Header) check() error {
	fields, values := fieldValue(h)

	_, err := forFields(fields, values)
	if err != nil {
		return err
	}
	return nil
}

// AddRplDCT Creates and returns the DCT information
// Needs to be replaced
func (s *SIL) AddRplDCT() {
	s.Header.check()
}
