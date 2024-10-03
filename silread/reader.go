package silread

import (
	"fmt"
	"github.com/kjbreil/loc-macro/pkg/script"
	"io"
	"reflect"
)

type Reader struct {
	r io.ReadSeekCloser
	Stats

	p      *parser
	macros script.Script
}

func NewReader(rs io.ReadSeekCloser) (*Reader, error) {
	r := &Reader{
		r: rs,
	}
	var err error

	// initialize the parser with the reader
	r.p = newParser(r.r)

	// get the stats from the file
	err = r.newStatsFromReader()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func NewReaderOnly(rs io.ReadSeekCloser, validTables []string) (*Reader, error) {
	r := &Reader{
		r: rs,
	}
	var err error

	// initialize the parser with the reader
	r.p = newParser(r.r)

	// get the stats from the file
	err = r.newStatsFromReaderOnly(validTables)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Reader) ClearReader() {
	r.r = nil
}

func UnmarshalReader(r io.Reader, data any) error {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return fmt.Errorf("data needs to be a pointer to a slice")
	}
	// check if data is a slice
	if reflect.TypeOf(data).Elem().Kind() != reflect.Slice {
		return fmt.Errorf("data is not a slice")
	}

	// make a new parser with the reader
	dataType := reflect.TypeOf(data).Elem().Elem()

	// make a channel of the type for datatype
	dataChan := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, dataType), 100)

	err := UnmarshalReaderChan(r, dataChan.Interface())
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

func (r *Reader) UnmarshalReaderChan(dataChan any) error {
	if reflect.TypeOf(dataChan).Kind() != reflect.Chan {
		return fmt.Errorf("data needs to be a channel")
	}
	_, err := r.r.Seek(0, 0)
	if err != nil {
		return err
	}

	// make a new parser with the reader
	p := newParser(r.r)

	go p.decodeChan(dataChan)

	return nil
}

func (r *Reader) Type() io.ReadSeeker {
	return r.r
}

func (r *Reader) Close() error {
	return r.r.Close()
}
