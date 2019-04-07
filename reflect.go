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

	// return be depending on kind
	b := kind(&v, &dt)
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

// this needs all kinds of work. only works for int and string at the moment
func kind(v *reflect.Value, dt *string) string {
	// has default tag boolean
	hd := len(*dt) != 0

	switch v.Kind() {
	case reflect.Ptr: // pointer so return kind of Elem()
		nv := v.Elem()
		s := kind(&nv, dt)
		return s
	case reflect.Int:
		switch {
		case v.Int() == 0 && hd: // empty INT with default
			return *dt
		default:
			return strconv.Itoa(int(v.Int()))
		}
	case reflect.String:
		switch {
		case v.Len() == 0 && hd: // empty string with default
			return fmt.Sprintf("'%s'", *dt)
		case v.Len() == 0: // without default
			return ""
		default:
			return fmt.Sprintf("'%s'", v.String())
		}
	default: // not defined above gets a blank entry
		return ""
	}
}
