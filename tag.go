package sil

import (
	"fmt"
	"reflect"
	"strings"
)

type silTag struct {
	field *reflect.StructField
}

// getSilTag takes the StructField and returns a silTag pointer along with a pointer to a zero-pad int lenght of the zero pad
func getSilTag(f *reflect.StructField) (*silTag, bool, error) {

	st := silTag{
		field: f,
	}

	options, err := st.get()
	if err != nil {
		return nil, false, err
	}

	if options != nil {
		return &st, true, nil
	}

	return &st, false, nil
}

func (st *silTag) get() (*[]string, error) {
	silTag := strings.Split(st.field.Tag.Get("sil"), ",")

	switch silTag[0] {
	case "":
		return nil, fmt.Errorf("does not contain a sil tag")
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
			// _, ok := strconv.Atoi(silTag[i])
			// if ok != nil {
			// 	options = append(options, silTag[i])
			// }
		}
		if len(options) > 0 {
			return &options, nil
		}
	}

	return nil, nil
}

func silArguments(args []string) {

}
