// Package silread reads sil files into SIL or Multi objects from the sil module
package silread

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/kjbreil/sil"
)

// Unmarshal SIL bytes into a interface{}
func Unmarshal(b []byte, v interface{}) (s sil.SIL, err error) {
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
		fieldIndex := findFieldIndex(ef, v)
		if fieldIndex == -1 {
			err = fmt.Errorf("field %s does not exist in type definition", ef)
			return
		}

		fieldMap = append(fieldMap, fieldIndex)
	}

	return d.SIL(v, fieldMap)
}

func (d *decoder) SIL(v interface{}, fieldMap []int) (s sil.SIL, err error) {
	// if the fieldmap is 0 length then no data was read, probably empty lines
	if len(fieldMap) == 0 {
		err = fmt.Errorf("fieldMap was 0 when passed to SIL")
		return
	}

	s.TableType = d.tableName
	s.View.Name = d.tableName

	for i := range d.data {
		values := reflect.ValueOf(v)
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

		indirect := reflect.Indirect(reflect.ValueOf(v))
		newIndirect := reflect.New(indirect.Type())
		// set the value of newIndirect to the value of indirect which is inturn the value of v
		newIndirect.Elem().Set(reflect.ValueOf(indirect.Interface()))
		// set data to the elem() of the newIndirect (so direct)
		data := newIndirect.Elem().Interface()

		s.View.Data = append(s.View.Data, data)
	}

	// get ints for ending and active date/times
	endingDateInt, err := stringToInt(d.header[6])
	if err != nil {
		return
	}

	endingTimeInt, err := stringToInt(d.header[7])
	if err != nil {
		return
	}

	activeDateInt, err := stringToInt(d.header[8])

	if err != nil {
		return
	}

	activeTimeInt, err := stringToInt(d.header[9])

	if err != nil {
		return
	}

	s.Header = sil.Header{
		Type:              d.header[0],
		Identifier:        d.header[1],
		Creator:           d.header[2],
		Destination:       d.header[3],
		AuditFile:         d.header[4],
		ResponseFile:      d.header[5],
		EndingDate:        endingDateInt,
		EndingTime:        endingTimeInt,
		ActiveDate:        activeDateInt,
		ActiveTime:        activeTimeInt,
		PurgeDate:         d.header[10],
		ActionType:        d.header[11],
		Description:       d.header[12],
		UserOneState:      d.header[13],
		MaximumErrorCount: d.header[14],
		FileVersion:       d.header[15],
		CreatorVersion:    d.header[16],
		PrimaryKey:        d.header[17],
		SpecificCommand:   d.header[18],
		TagType:           d.header[19],
		ExecutionPriority: d.header[20],
		LongDescription:   "",
	}

	return
}

func findFieldIndex(fcode string, v interface{}) int {
	tp := reflect.TypeOf(v)
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
