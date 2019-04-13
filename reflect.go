package sil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// row is the makeup of each row with an array of elem
type row struct {
	elems []elem
}

// elem is the parts of the row containg the name of the field and the data
// held within
type elem struct {
	name *string
	data *string
}

func (r *row) make(rowType interface{}, include bool) {
	values, fields := reflect.ValueOf(rowType), reflect.TypeOf(rowType)
	// loop over the fields
	for i := 0; i < fields.NumField(); i++ {
		val, name, ptr, _ := value(values.Field(i), fields.Field(i))
		switch {
		case include && *ptr && *val == "":
			var v string
			r.elems = append(r.elems, elem{
				name: name,
				data: &v,
			})
		case !include && *ptr && *val == "":
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
		val, _, _, _ := value(values.Field(i), fields.Field(i))
		sa = append(sa, *val)
	}
	// join the strings with commas and put it in ()
	s := fmt.Sprintf("(%s)", strings.Join(sa, ","))
	return []byte(s)
}

func value(v reflect.Value, f reflect.StructField) (*string, *string, *bool, error) {
	// get the silTag
	_, err := getSilTag(&f)
	if err != nil {
		return nil, nil, nil, err
	}
	// get the name of the object
	name := f.Name

	// get default tag
	dt := defaultTag(&f)

	// return bytes depending on kind
	bytes, pointer := kind(&v, &dt)
	switch name {
	case "F01":
		bytes = fmt.Sprintf("'%013v'", bytes[1:len(bytes)-1])
	}
	return &bytes, &name, &pointer, nil
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
		return reflectInt(v, dt, &hd), false
	case reflect.String:
		return reflectString(v, dt, &hd), false
	default: // not defined above gets a blank entry
		return "", false
	}
}

func reflectString(v *reflect.Value, dt *string, hd *bool) string {
	switch {
	case v.Len() == 0 && *hd: // empty string with default
		return fmt.Sprintf("'%s'", *dt)
	case v.Len() == 0: // without default
		return ""

	default:
		return fmt.Sprintf("'%s'", v.String())
	}
}

func reflectInt(v *reflect.Value, dt *string, hd *bool) string {
	switch {
	case v.Int() == 0 && *hd: // empty INT with default
		return *dt
	default:
		return strconv.Itoa(int(v.Int()))
	}
}
