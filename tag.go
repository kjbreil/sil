package sil

import (
	"fmt"
	"reflect"
	"strings"
)

type silTag struct {
	name    string
	options []string
	field   *reflect.StructField
}

// getSilTag takes the StructField and returns a silTag pointer along with bool for padding
func getSilTag(f *reflect.StructField) (*silTag, bool, error) {

	st := silTag{
		field: f,
	}
	// get assigns the name and options to the sil tag
	err := st.get()

	if err != nil {
		return nil, false, err
	}

	if len(st.options) > 0 {
		return &st, true, nil
	}

	return &st, false, nil
}

// get returns the sil tag and then an array of any other arguments for sil
func (st *silTag) get() error {
	silTag := strings.Split(st.field.Tag.Get("sil"), ",")

	switch silTag[0] {
	case "":
		return fmt.Errorf("does not contain a sil tag")
	}

	if len(silTag) > 1 {
		var options []string
		// options := silTag[1:]
		for i := range silTag {
			if i == 0 {
				continue
			}

			if silTag[i] == "zeropad" {
				options = append(options, silTag[i])
			}
		}
		if len(options) > 0 {
			st.options = options
		}
	}

	st.name = silTag[0]

	return nil
}
