package sil

import (
	"fmt"
	"reflect"
	"strings"
)

type silTag struct {
	field *reflect.StructField
}

func getSilTag(f *reflect.StructField) (*silTag, error) {

	st := silTag{
		field: f,
	}

	err := st.get()
	if err != nil {
		return nil, err
	}

	return &st, nil
}

func (st *silTag) get() error {
	silTag := strings.Split(st.field.Tag.Get("sil"), ",")

	switch silTag[0] {
	case "":
		return fmt.Errorf("does not contain a sil tag")
	}

	return nil
}

func silArguments(args []string) {

}
