package sil

import "reflect"

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
