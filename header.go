package sil

import "reflect"

// Check and fix a header
func (h *Header) Check() error {

	fields := reflect.TypeOf(h)
	values := reflect.ValueOf(h)

	// check if the interface is a pointer and then get the elements that it points
	// to - fixes panic: reflect: NumField of non-struct type
	if fields.Kind() == reflect.Ptr && fields.Elem().Kind() == reflect.Struct {
		fields = fields.Elem()
		values = values.Elem()
	}

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		// if there is no sil tag skip
		if field.Tag.Get("sil") == "" {
			continue
		}
		// get the default tag
		tag := field.Tag.Get("default")
		// if there is no sil tag skip
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
