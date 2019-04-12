package sil

import (
	"fmt"
	"math/rand"
	"time"
)

// SIL is the structure of a SIL file
type SIL struct {
	Header Header
	View   View
	Footer Footer

	TableType interface{}
}

// Some Constants
const (
	crlf = "\r\n"
)

// Make makes a sil file of the definiton (as struct) passed
func Make(name string, definition interface{}) *SIL {
	s := new(SIL)
	// store the name of the table in the returned sil file
	s.View.Name = name
	s.View.Required = true
	return s
}

// Marshal creates the SIL structure from the information in the SIL type
func (s *SIL) Marshal() (data []byte, err error) {
	rand.Seed(time.Now().UnixNano())
	prefix := rand.Intn(100)
	// get the multiple sections
	secs := multi(s.View.Data)
	for _, sec := range secs {
		// Create the Header insert
		s.Header.F902 = batchNum(prefix)
		data = append(data, s.Header.insert()...)
		data = append(data, s.Header.row()...)

		data = append(data, sec.create(s.View.Name)...)
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
	return fmt.Sprintf("%02d%02d%02d%02d", prefix, t.Hour(), t.Minute(), t.Second())
}
