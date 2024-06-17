package silread

import (
	"fmt"
	"io"
	"reflect"
)

func UnmarshalReader(r io.Reader, data any) error {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return fmt.Errorf("data needs to be a pointer to a slice")
	}
	// check if data is a slice
	if reflect.TypeOf(data).Elem().Kind() != reflect.Slice {
		return fmt.Errorf("data is not a slice")
	}

	// make a new parser with the reader
	p := newParser(r)

	p.decode()

	return nil
}

func UnmarshalReaderChan(r io.Reader, dataChan any) error {
	if reflect.TypeOf(dataChan).Kind() != reflect.Chan {
		return fmt.Errorf("data needs to be a channel")
	}

	// make a new parser with the reader
	p := newParser(r)

	go p.decodeChan(dataChan)

	return nil
}
