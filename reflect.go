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

// elem is the parts of the row containing the name of the field and the data
// held within
type elem struct {
	name *string
	data *string
}

func (r *row) make(rowType interface{}, include bool) error {
	values, fields := reflect.ValueOf(rowType), reflect.TypeOf(rowType)
	// loop over the fields
	for i := 0; i < fields.NumField(); i++ {
		// get the value, name and if its a pointer
		val, name, ptr, err := value(values.Field(i), fields.Field(i))
		if err != nil {
			return err
		}
		switch {
		case !*ptr && *val == "": // panic if its a required field without any data
			return fmt.Errorf("the element %s does not contain any data and is required for table %s", *name, fields.Name())
		case include && *ptr && *val == "":
			var v string
			r.elems = append(r.elems, elem{
				name: name,
				data: &v,
			})
		case !include && *ptr && *val == "": // if we are not including pointers and it is a pointer and does not have a value contiue (skip)
			continue
		default:
			r.elems = append(r.elems, elem{
				name: name,
				data: val,
			})
		}
	}
	return nil
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
	t, pad, err := getSilTag(&f)
	if err != nil {
		return nil, nil, nil, err
	}

	// get default tag
	dt := defaultTag(&f)

	// return bytes depending on kind
	bytes, pointer := kind(&v, &dt)

	if pad && len(bytes) > 0 {
		bytes = fmt.Sprintf("'%013v'", bytes[1:len(bytes)-1])
	}

	if pad && len(bytes) > 13 {
		return nil, nil, nil, fmt.Errorf("padded field contains more than 13 characters %s", string(bytes))
	}

	return &bytes, &t.name, &pointer, nil
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
	case "NOWTIME":
		return JulianTimeNow()
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

// test
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
