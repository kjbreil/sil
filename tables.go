package sil

import (
	"reflect"
	"strings"
)

// MakeTable makes a table after passing the name of the table and the type interface
// should improve to include the name in the type itself and validation that
// the type passed is valid to be passed into this interface (i.e. confirm that
// the fields all follow the correct naming convention for SIL)
// this first looks to a sil tag on the type to assign the SQL type otherwise tries to predict
func (s *SIL) MakeTable(name string, tableType interface{}) {
	s.Table.Name = name

	fields := reflect.TypeOf(tableType)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		tag := field.Tag.Get("sil")
		typeName := field.Type.Name()
		var fieldType string

		if tag != "" {
			fieldType = strings.ToUpper(tag)

		} else {
			switch typeName {
			case "int":
				fieldType = "INTEGER"
			case "string":
				fieldType = "CHAR(30)"
			default:
				fieldType = typeName
			}
		}

		// fieldType := "INTEGER"

		column := Column{
			Name: field.Name,
			Type: fieldType,
		}

		s.Table.Columns = append(s.Table.Columns, column)

	}

}
