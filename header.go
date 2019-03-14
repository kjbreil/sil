package sil

// check and fix a header with default tag
func (h *Header) check() {
	fields, values := fieldValue(h)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		// if there is no sil tag skip
		if field.Tag.Get("sil") == "" {
			continue
		}
		// get the default tag
		tag := field.Tag.Get("default")
		// if there is no default tag skip
		if tag == "" {
			continue
		}

		// if the value is not there insert default
		if value.Len() == 0 {
			switch value.Type().Name() {
			case "string":
				value.SetString(tag)
			}
		}
	}
}

// AddRplDCT Creates and returns the DCT information
// Needs to be replaced
func (s *SIL) AddRplDCT() {
	s.Header.check()
}
