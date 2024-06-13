// Package silread reads sil files into SIL or Multi objects from the sil module
package silread

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal SIL bytes into a interface{}
func Unmarshal(b []byte, data any) (err error) {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		err = fmt.Errorf("data needs to be a pointer to a slice")
		return
	}
	// check if data is a slice
	if reflect.TypeOf(data).Elem().Kind() != reflect.Slice {
		err = fmt.Errorf("data is not a slice")
		return
	}

	// open a reader using the bytes as the start
	// this can be improved to read directly from a file
	r := bytes.NewReader(b)

	// make a new parser with the reader
	p := newParser(r)
	// prsd is the parsed file token parts
	prsd := p.parse()

	d, _ := prsd.decode(0)
	if len(d.err) > 0 {
		err = fmt.Errorf("could not decode the parsed sil file: %v", d.err)
		return
	}

	// fieldMap isn't really a map in the go sense but is where the Fcode is in the type
	var fieldMap []int

	for _, ef := range d.fcodes {
		fieldIndex := findFieldIndex(ef, data)
		if fieldIndex == -1 {
			err = fmt.Errorf("field %s does not exist in type definition", ef)
			return
		}

		fieldMap = append(fieldMap, fieldIndex)
	}

	err = d.unmarshal(data, fieldMap)
	if err != nil {
		return err
	}

	return
}

func (d *decoder) unmarshal(data any, fieldMap []int) (err error) {
	// if the fieldmap is 0 length then no data was read, probably empty lines
	if len(fieldMap) == 0 {
		err = fmt.Errorf("fieldMap is empty")
		return
	}

	var isPointerValue bool
	tableType := reflect.TypeOf(data).Elem().Elem()
	if tableType.Kind() == reflect.Ptr {
		tableType = tableType.Elem()
		isPointerValue = true
	}

	viewDataSlice := reflect.MakeSlice(reflect.TypeOf(data).Elem(), 0, len(d.data))

	for i := range d.data {
		// create then new values of TableType to insert the data into
		entry := reflect.New(tableType)
		values := entry
		if values.Kind() == reflect.Ptr {
			values = values.Elem()
		}

		for c := range d.data[i] {
			if !values.Field(fieldMap[c]).CanSet() {
				err = fmt.Errorf("cannot set field @%d named %s", i, values.Field(fieldMap[c]).Type().Name())
				return
			}

			if values.Field(fieldMap[c]).CanSet() {
				switch values.Field(fieldMap[c]).Type().Name() {
				case "string":
					values.Field(fieldMap[c]).SetString(d.data[i][c])
				case "int":
					var dataInt int64
					// for when the data is empty
					if len(d.data[i][c]) == 0 {
						dataInt = 0
					} else {
						dataInt, err = strconv.ParseInt(d.data[i][c], 10, 64)
					}
					if err != nil {
						err = fmt.Errorf("conversion of data type int did not convert from %s err: %v", d.data[i][c], err)
						return
					}

					values.Field(fieldMap[c]).SetInt(dataInt)
				default:
					// probably a pointer
					switch values.Field(fieldMap[c]).Type().Elem().Name() {
					case "string":
						data := d.data[i][c]
						values.Field(fieldMap[c]).Set(reflect.ValueOf(&data))
					case "int":
						var dataInt int
						if len(d.data[i][c]) == 0 {
							dataInt = 0
						} else {
							dataInt, err = strconv.Atoi(d.data[i][c])
						}
						if err != nil {
							err = fmt.Errorf("conversion of data type int did not convert from %s err: %v", d.data[i][c], err)
							return
						}

						values.Field(fieldMap[c]).Set(reflect.ValueOf(&dataInt))
					}
				}
			}
		}
		if isPointerValue {
			viewDataSlice = reflect.Append(viewDataSlice, entry)

		} else {

			viewDataSlice = reflect.Append(viewDataSlice, values)
		}
	}

	viewDataValue := reflect.Indirect(reflect.ValueOf(data))

	viewDataValue.Set(viewDataSlice)
	fmt.Println(viewDataSlice.Interface())
	fmt.Println(data)
	fmt.Println("line")

	return
}

func findFieldIndex(fcode string, v interface{}) int {
	tp := reflect.TypeOf(v)
	// walk down the first pointer
	// TODO: These should return errors
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	// look into the slice
	tp = tp.Elem()
	// if its a pointer go into the pointer
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		tag := getSilTag(&field)

		if tag == fcode {
			return i
		}
	}

	return -1
}

// getSilTag takes the StructField and returns a silTag pointer along with bool for padding
// needs to actuall reference the sil module function (which needs to be exported)
func getSilTag(f *reflect.StructField) string {
	silTag := strings.Split(f.Tag.Get("sil"), ",")

	switch silTag[0] {
	case "":
		log.Panicln("does not contain a sil tag")
	}

	return silTag[0]
}

func stringToInt(is string) (int, error) {

	if len(is) == 0 {
		return 0, nil
	}

	endingDateInt, err := strconv.Atoi(is)
	if err != nil {
		err = fmt.Errorf("header ending date did not convert to int %s", is)
		return 0, err
	}

	return endingDateInt, nil
}
