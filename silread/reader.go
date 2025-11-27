package silread

import (
	"context"
	"fmt"
	"io"
	"reflect"

	"github.com/kjbreil/loc-macro/pkg/script"
)

type Reader struct {
	r io.ReadSeekCloser
	Stats

	p      *parser
	macros script.Script
}

// NewReader creates a new Reader from the given ReadSeekCloser
// Uses context.Background() for backwards compatibility
func NewReader(rs io.ReadSeekCloser) (*Reader, error) {
	return NewReaderContext(context.Background(), rs)
}

// NewReaderContext creates a new Reader from the given ReadSeekCloser with context support
// The context can be used to cancel the operation
func NewReaderContext(ctx context.Context, rs io.ReadSeekCloser) (*Reader, error) {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r := &Reader{
		r: rs,
	}
	var err error

	// initialize the parser with the reader
	r.p = newParser(r.r)

	// get the stats from the file
	err = r.newStatsFromReaderWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// NewReaderOnly creates a new Reader from the given ReadSeekCloser, only allowing specific tables
// Uses context.Background() for backwards compatibility
func NewReaderOnly(rs io.ReadSeekCloser, validTables []string) (*Reader, error) {
	return NewReaderOnlyContext(context.Background(), rs, validTables)
}

// NewReaderOnlyContext creates a new Reader from the given ReadSeekCloser, only allowing specific tables
// The context can be used to cancel the operation
func NewReaderOnlyContext(ctx context.Context, rs io.ReadSeekCloser, validTables []string) (*Reader, error) {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r := &Reader{
		r: rs,
	}
	var err error

	// initialize the parser with the reader
	r.p = newParser(r.r)

	// get the stats from the file
	err = r.newStatsFromReaderOnlyWithContext(ctx, validTables)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Reader) ClearReader() {
	r.r = nil
}

// UnmarshalReader unmarshals data from a reader into a slice
// Uses context.Background() for backwards compatibility
func UnmarshalReader(r io.Reader, data any) error {
	return UnmarshalReaderContext(context.Background(), r, data)
}

// UnmarshalReaderContext unmarshals data from a reader into a slice with context support
// The context can be used to cancel the operation
func UnmarshalReaderContext(ctx context.Context, r io.Reader, data any) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

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

	err := UnmarshalReaderChanContext(ctx, r, dataChan.Interface())
	if err != nil {
		return err
	}
	viewDataSlice := reflect.MakeSlice(reflect.SliceOf(dataType), 0, 0)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
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

// UnmarshalReaderChan unmarshals data from a reader into a channel
// Uses context.Background() for backwards compatibility
func UnmarshalReaderChan(r io.Reader, dataChan any) error {
	return UnmarshalReaderChanContext(context.Background(), r, dataChan)
}

// UnmarshalReaderChanContext unmarshals data from a reader into a channel with context support
// The context can be used to cancel the operation
func UnmarshalReaderChanContext(ctx context.Context, r io.Reader, dataChan any) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

	if reflect.TypeOf(dataChan).Kind() != reflect.Chan {
		return fmt.Errorf("data needs to be a channel")
	}

	// make a new parser with the reader
	p := newParser(r)

	go p.decodeChanWithContext(ctx, dataChan)

	return nil
}

// UnmarshalReaderChan unmarshals data from the Reader into a channel
// Uses context.Background() for backwards compatibility
func (r *Reader) UnmarshalReaderChan(dataChan any) error {
	return r.UnmarshalReaderChanContext(context.Background(), dataChan)
}

// UnmarshalReaderChanContext unmarshals data from the Reader into a channel with context support
// The context can be used to cancel the operation
func (r *Reader) UnmarshalReaderChanContext(ctx context.Context, dataChan any) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

	if reflect.TypeOf(dataChan).Kind() != reflect.Chan {
		return fmt.Errorf("data needs to be a channel")
	}
	_, err := r.r.Seek(0, 0)
	if err != nil {
		return err
	}

	// make a new parser with the reader
	p := newParser(r.r)

	go p.decodeChanWithContext(ctx, dataChan)

	return nil
}

func (r *Reader) Type() io.ReadSeeker {
	return r.r
}

func (r *Reader) Close() error {
	return r.r.Close()
}
