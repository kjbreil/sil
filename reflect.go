package sil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type row struct {
	elems []elem
}

type elem struct {
	name *string
	data *string
}

func (r *row) make(rowType interface{}, include bool) {
	values, fields := reflect.ValueOf(rowType), reflect.TypeOf(rowType)
	// loop over the fields
	for i := 0; i < fields.NumField(); i++ {
		val, name, ptr := value(values.Field(i), fields.Field(i))
		switch {
		case include && *ptr == true && *val == "":
			var v string
			r.elems = append(r.elems, elem{
				name: name,
				data: &v,
			})
		case !include && *ptr == true && *val == "":
			continue
		default:
			r.elems = append(r.elems, elem{
				name: name,
				data: val,
			})
		}
	}
}

func rowBytes(rowType interface{}) []byte {
	values, fields := reflect.ValueOf(rowType), reflect.TypeOf(rowType)
	// string array
	var sa []string
	// loop over the fields
	for i := 0; i < fields.NumField(); i++ {
		val, _, _ := value(values.Field(i), fields.Field(i))
		sa = append(sa, *val)
	}
	// join the strings with commas and put it in ()
	s := fmt.Sprintf("(%s)", strings.Join(sa, ","))
	return []byte(s)
}

func value(v reflect.Value, f reflect.StructField) (*string, *string, *bool) {
	silTag := f.Tag.Get("sil")
	if silTag == "" {
		return nil, nil, nil
	}

	// get the name of the object
	name := f.Name

	// get default tag
	dt := defaultTag(&f)

	// return bytes depending on kind
	bytes, pointer := kind(&v, &dt)
	return &bytes, &name, &pointer
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
func kind(v *reflect.Value, dt *string) (value string, pointer bool) {
	// has default tag boolean
	hd := len(*dt) != 0

	switch v.Kind() {
	case reflect.Ptr: // pointer so return kind of Elem()
		nv := v.Elem()
		s, _ := kind(&nv, dt)
		return s, true
	case reflect.Int:
		switch {
		case v.Int() == 0 && hd: // empty INT with default
			return *dt, false
		default:
			return strconv.Itoa(int(v.Int())), false
		}
	case reflect.String:
		switch {
		case v.Len() == 0 && hd: // empty string with default
			return fmt.Sprintf("'%s'", *dt), false
		case v.Len() == 0: // without default
			return "", false
		default:
			return fmt.Sprintf("'%s'", v.String()), false
		}
	default: // not defined above gets a blank entry
		return "", false
	}
}
