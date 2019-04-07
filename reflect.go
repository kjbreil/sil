package sil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func row(rowType interface{}) []byte {
	values, fields := reflect.ValueOf(rowType), reflect.TypeOf(rowType)
	// string array
	var sa []string
	// loop over the fields
	for i := 0; i < fields.NumField(); i++ {
		val := value(values.Field(i), fields.Field(i))
		sa = append(sa, *val)
	}
	// join the strings with commas and put it in ()
	s := fmt.Sprintf("(%s)", strings.Join(sa, ","))
	return []byte(s)
}

func value(v reflect.Value, f reflect.StructField) *string {
	silTag := f.Tag.Get("sil")
	if silTag == "" {
		return nil
	}
	// get default tag
	dt := defaultTag(&f)
	if dt != "" {
		switch silTag {
		case "INTEGER":
		default:
			dt = fmt.Sprintf("'%s'", dt)
		}
		return &dt

	}
	// return be depending on kind
	b := kind(&v)
	return &b
}

func defaultTag(f *reflect.StructField) string {
	// get the default tag
	def := f.Tag.Get("default")
	// switch on default tag for special functions
	switch def {
	case "":
		return ""
	case "NOW":
		return JulianNow()
	default:
		return def
	}
}

func kind(v *reflect.Value, dt *string) string {
	switch v.Kind() {
	case reflect.Int:
		return strconv.Itoa(int(v.Int()))
	case reflect.String:
		if v.Len() == 0 {
			return v.String()
		}
		return fmt.Sprintf("'%s'", v.String())
	default: // not defined above gets a '', no data comes through
		return "''"
	}
}
