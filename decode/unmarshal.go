package decode

import (
	"bytes"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal SIL bytes into a interface{}
func Unmarshal(b []byte, v interface{}) {
	// open a reader using the bytes as the start
	// this can be improved to read directly from a file
	r := bytes.NewReader(b)

	// make a new parser with the reader
	p := newParser(r)
	// prsd is the parsed file token parts
	prsd := p.parse()

	// just displaying the parts for now
	// for _, ep := range *prsd {
	// 	fmt.Println(ep)
	// }

	d := prsd.decode()
	if len(d.err) > 0 {
		log.Fatalf("could not decode the parsed sil file: %v\n", d.err)
	}

	// fmt.Println(d.data)

	// fieldMap isn't really a map in the go sense but is where the Fcode is in the type
	var fieldMap []int

	for _, ef := range d.fcodes {
		fieldIndex := findFieldIndex(ef, v)
		fieldMap = append(fieldMap, fieldIndex)
	}

	for i := range d.data {

		values := reflect.ValueOf(v).Elem()

		for c := range d.data[i] {
			if values.Field(fieldMap[c]).CanSet() {
				switch values.Field(fieldMap[c]).Type().Name() {
				case "string":
					values.Field(fieldMap[c]).SetString(d.data[i][c])
				case "int":
					dataInt, err := strconv.ParseInt(d.data[i][c], 10, 64)
					if err != nil {
						log.Panicln(err)
					}
					values.Field(fieldMap[c]).SetInt(dataInt)
				}
			}
		}

		indirect := reflect.Indirect(reflect.ValueOf(v))
		newIndirect := reflect.New(indirect.Type())
		// set the value of newIndirect to the value of indirect which is inturn the value of v
		newIndirect.Elem().Set(reflect.ValueOf(indirect.Interface()))
		// set data to the elem() of the newIndirect (so direct)
		data := newIndirect.Elem().Interface()

		d.s.View.Data = append(d.s.View.Data, data)

	}
	err := d.s.Write("test.sil", false, false)
	if err != nil {
		log.Println(err)
	}
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
func getSilTag(f *reflect.StructField) string {
	silTag := strings.Split(f.Tag.Get("sil"), ",")

	switch silTag[0] {
	case "":
		log.Panicln("does not contain a sil tag")
	}
	return silTag[0]
}
