package sil

import (
	"fmt"
	"reflect"
	"strings"
)

// MakeTable makes a table after passing the name of the table and the type interface
// should improve to include the name in the type itself and validation that
// the type passed is valid to be passed into this interface (i.e. confirm that
// the fields all follow the correct naming convention for SIL)
// this first looks to a sil tag on the type to assign the SQL type otherwise tries to predict
func (s *SIL) MakeTable(tableType interface{}) {
	// reflect the tableType to get the fields
	fields := reflect.TypeOf(tableType)

	// loop over the fields
	for i := 0; i < fields.NumField(); i++ {
		// get the single field
		field := fields.Field(i)
		// get the tag is it exists
		tag := field.Tag.Get("sil")
		typeName := field.Type.Name()
		var fieldType string

		// if the tag is not blank then force uppercase on tag just in case
		// otherwise try and predict, if that fails it uses the GO type name which will
		// not work so should error. I still haven't decided how to handle pointers or
		// bad references in the type definitions
		if tag != "" {
			fieldType = strings.ToUpper(tag)

		} else {
			switch typeName {
			case "int":
				fieldType = sqlInt
			case "string":
				fieldType = "CHAR(30)"
			default:
				fieldType = typeName
			}
		}

		column := Column{
			Name: field.Name,
			Type: fieldType,
		}

		s.Table.Columns = append(s.Table.Columns, column)
	}
}

// MakeRow makes a row of data to be inserted based on the type definition and data in the object
func MakeRow(rowData interface{}) string {
	// create members to hold the strings to join later
	var members []string

	fields := reflect.TypeOf(rowData)
	values := reflect.ValueOf(rowData)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		// get the SIL tag from the object
		tag := field.Tag.Get("sil")
		// if there is no sil tag skip
		if tag == "" {
			continue
		}

		// INTEGERS need to be insterted without single quotes, all others with single quotes
		if tag == sqlInt || value.Len() == 0 {
			members = append(members, fmt.Sprintf("%v", value))
		} else {
			members = append(members, fmt.Sprintf("'%v'", value))
		}
	}

	return (fmt.Sprintf("(%s)", strings.Join(members, ",")))

}
