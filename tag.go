package sil

import (
	"fmt"
	"reflect"
)

func silTag(f *reflect.StructField) error {
	silTag := f.Tag.Get("sil")
	switch silTag {
	case "":
		return fmt.Errorf("does not contain a sil tag")
	}
	return nil
}
