package sil

import (
	"fmt"
	"strings"
)

// Check and fix a header
func (h *Header) Check() error {
	// for the error, which gets returned for any defaults inserted
	var defaulted []string

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
				defaulted = append(defaulted, field.Name)
			}
		}
	}
	if len(defaulted) != 0 {
		return fmt.Errorf("defaults inserted: %s", strings.Join(defaulted, ","))
	}
	return nil
}

// AddRplDCT Creates and returns the DCT information
// Needs to be replaced
func (s *SIL) AddRplDCT() {

	err := s.ViewHeader.Check()
	if err != nil {
		fmt.Println(err)
	}
}
