// Package silread reads sil files into SIL or Multi objects from the sil module
package silread

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
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

	dataType := reflect.TypeOf(data).Elem().Elem()

	// make a channel of the type for datatype
	dataChan := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, dataType), 100)

	err = UnmarshalReaderChan(r, dataChan.Interface())
	if err != nil {
		return err
	}
	viewDataSlice := reflect.MakeSlice(reflect.SliceOf(dataType), 0, 0)

	for {
		v, ok := dataChan.Recv()
		if !ok {
			break
		}
		viewDataSlice = reflect.Append(viewDataSlice, v)
	}

	viewDataValue := reflect.Indirect(reflect.ValueOf(data))

	viewDataValue.Set(viewDataSlice)

	return
}

func unmarshalValues(input []string, result reflect.Value, fieldMap []int) (err error) {
	// if the fieldmap is 0 length then no data was read, probably empty lines
	if len(fieldMap) == 0 {
		err = fmt.Errorf("fieldMap is empty")
		return
	}

	if result.Kind() == reflect.Ptr {
		result = result.Elem()
	}

	for c := range input {
		if fieldMap[c] == -1 {
			continue
		}
		if !result.Field(fieldMap[c]).CanSet() {
			err = fmt.Errorf("cannot set field @%d named %s", 1, result.Field(fieldMap[c]).Type().Name())
			return
		}

		if input[c] == "" {
			continue
		}

		if result.Field(fieldMap[c]).CanSet() {
			err = unmarshalValue(result.Field(fieldMap[c]), input[c])
			if err != nil {
				return err
			}
		}
	}

	return
}

func unmarshalValue(value reflect.Value, input string) error {
	var err error
	if !value.CanSet() {
		panic(fmt.Sprintf("cannot set"))
	}
	kind := value.Kind()

	if kind == reflect.Ptr {
		if !value.Elem().IsValid() || value.Elem().IsNil() {
			var t reflect.Type
			t = value.Type().Elem()
			value.Set(reflect.New(t))
			kind = t.Kind()
			value = value.Elem()
		} else {
			kind = value.Elem().Kind()
		}
	}

	switch kind {
	case reflect.String:
		value.SetString(input)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var dataInt int64
		if len(input) == 0 {
			dataInt = 0
		} else {
			dataInt, err = strconv.ParseInt(input, 10, 64)
		}
		if err != nil {
			return fmt.Errorf("conversion of data type int did not convert from %s err: %v", input, err)
		}
		value.SetInt(dataInt)
	case reflect.Struct:
		switch value.Interface().(type) {
		case time.Time:
			var err error
			var t time.Time
			if len(input) == 0 {
				return nil
			}
			if len(input) == 7 {
				t, err = time.Parse("2006002", input[:7])
			} else {
				t, err = time.Parse("2006002 15:04:05", input[:16])
			}
			if err == nil {
				value.Set(reflect.ValueOf(t))
			}
		default:
			return fmt.Errorf("unhandled type for unmarshalValue: %v", value.Type())
		}
	default:
		return fmt.Errorf("unhandled kind for unmarshalValue: %v", kind)
	}

	return nil
}

func findFieldIndex(fcode string, v interface{}) int {
	tp := reflect.TypeOf(v)
	// walk down the first pointer
	// TODO: These should return errors
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	// look into the slice
	if tp.Kind() == reflect.Slice {
		tp = tp.Elem()
	}
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
