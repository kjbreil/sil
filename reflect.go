package sil

import (
	"fmt"
	"reflect"
	"strconv"
)

func fieldValue(tableType interface{}) (reflect.Type, reflect.Value) {
	fields := reflect.TypeOf(tableType)
	values := reflect.ValueOf(tableType)

	// check if the interface is a pointer and then get the elements that it points
	// to - fixes panic: reflect: NumField of non-struct type
	if fields.Kind() == reflect.Ptr && fields.Elem().Kind() == reflect.Struct {
		fields = fields.Elem()
		values = values.Elem()
	}
	return fields, values
}

func forFields(fields reflect.Type, values reflect.Value) (members []string, err error) {
	// label Fields because of nested loops
Fields:
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		s := field.Tag.Get("sil")
		// if there is no sil tag skip
		if s == "" {
			continue
		}

		members = append(members, silTag(s, value))

		// get the default tag
		def := field.Tag.Get("default")
		// switch on default tag for special functions
		switch def {
		case "":
			continue Fields
		case "NOW":
			def = JulianNow()
		}

		// if the value is not there insert default. Defaults can only be on
		// required fields
		if value.Len() == 0 && value.Kind() != reflect.Ptr {
			switch value.Type().Name() {
			case "int":
				// declare here is to prevent shadow error below
				var is int64
				is, err = strconv.ParseInt(def, 10, 64)
				// the default did not convert to int so freak the f out
				if err != nil {
					return members, fmt.Errorf("default tag not int: %v", err)
				}
				value.SetInt(is)
			case "string": // strings fall in here
				value.SetString(def)

			}
		}
	}
	return
}

func silTag(tag string, value reflect.Value) string {

	// INTEGERS need to be insterted without single quotes, all others with single quotes
	switch {
	// if the element is a pointer and is nil return a blank
	case value.Kind() == reflect.Ptr && value.IsNil():
		return ""
	// if the element is a pointer and has a value assign the "value" variable
	// to the element of contained in the pointer and fallthrough to next case
	case value.Kind() == reflect.Ptr:
		value = value.Elem()
		fallthrough
	case tag == sqlInt && value.Type().Name() == "int":
		return fmt.Sprintf("%d", value.Int())
	case value.String() == "" || tag == sqlInt:
		return fmt.Sprintf("%v", value.String())
	default:
		return fmt.Sprintf("'%v'", value.String())
	}
}
