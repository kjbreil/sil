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

	d := prsd.decode()
	if len(d.err) > 0 {
		err = fmt.Errorf("could not decode the parsed sil file: %v\n", d.err)
		return
	}

	// fmt.Println(d.data)

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

	s.TableType = d.tableName
	s.View.Name = d.tableName

	for i := range d.data {

		values := reflect.ValueOf(v).Elem()

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
					dataInt, err = strconv.ParseInt(d.data[i][c], 10, 64)
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
						dataInt, err = strconv.Atoi(d.data[i][c])
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
	endingDateInt, err := strconv.Atoi(d.header[6])
	if err != nil {
		err = fmt.Errorf("header ending date did not convert to int %s", d.header[6])
		return
	}
	endingTimeInt, err := strconv.Atoi(d.header[7])
	if err != nil {
		err = fmt.Errorf("header ending time did not convert to int %s", d.header[7])
		return
	}
	activeDateInt, err := strconv.Atoi(d.header[8])
	if err != nil {
		err = fmt.Errorf("header active date did not convert to int %s", d.header[8])
		return
	}
	activeTimeInt, err := strconv.Atoi(d.header[9])
	if err != nil {
		err = fmt.Errorf("header active time did not convert to int %s", d.header[9])
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
	tp := reflect.TypeOf(v).Elem()
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
