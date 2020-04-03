package sil

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/text/encoding/charmap"
)

// SIL is the structure of a SIL file
type SIL struct {
	Header Header
	View   View
	Footer Footer

	prefix int

	TableType interface{}
}

// Some Constants
const (
	crlf = "\r\n"
)

// Make makes a sil file of the definition (as struct) passed
func Make(name string, definition interface{}) *SIL {
	s := new(SIL)
	// store the name of the table in the returned sil file
	s.View.Name = name
	s.View.Required = true
	rand.Seed(time.Now().UnixNano())
	s.prefix = rand.Intn(100)
	return s
}

// Marshal creates the SIL structure from the information in the SIL type
func (s *SIL) Marshal(include bool) (data []byte, err error) {
	// check to make sure the view.Name has been set
	if s.View.Name == "" {
		return data, fmt.Errorf("view name not set")
	}
	// get the multiple sections
	secs, err := split(s.View.Data, include)
	if err != nil {
		return []byte{}, err
	}
	for _, sec := range secs {
		// Create the Header insert
		s.Header.Identifier = batchNum(s.prefix)
		if s.View.Action != "LOAD" {
			data = append(data, s.Header.insert()...)
			data = append(data, s.Header.row()...)
		}
		data = append(data, sec.create(&s.View)...)

		data = append(data, []byte(crlf)...)
	}

	data = append(data, s.Footer.bytes()...)

	data, err = charmap.Windows1252.NewEncoder().Bytes(data)
	if err != nil {
		return []byte{}, fmt.Errorf("conversion to 1252 failed: %v", err)
	}

	return data, nil
}

func endLine() []byte {
	return []byte(fmt.Sprintf(";%s%s", crlf, crlf))
}

// batchNum returns a batch number based on the current time
// pass a prefix that will group the batches together
func batchNum(prefix int) string {
	t := time.Now()
	rand.Seed(t.UnixNano())
	return fmt.Sprintf("%02d%06d", prefix, rand.Intn(1000000))
}
