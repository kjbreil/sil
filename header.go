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
