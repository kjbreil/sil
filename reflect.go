package sil

import (
	"fmt"
	"reflect"
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

func forFields(fields reflect.Type, values reflect.Value) (members []string) {
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		sil := field.Tag.Get("sil")
		// if there is no sil tag skip
		if sil == "" {
			continue
		}

		members = append(members, silTag(sil, value.String()))

		// get the default tag
		def := field.Tag.Get("default")
		// if there is no default tag skip
		if def == "" {
			continue
		}

		// if the value is not there insert default
		if value.Len() == 0 {
			switch value.Type().Name() {
			case "string":
				value.SetString(def)
			}
		}
	}
	return
}

func silTag(tag string, value string) string {

	// INTEGERS need to be insterted without single quotes, all others with single quotes
	if tag == sqlInt || value == "" {
		return fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf("'%v'", value)
}
