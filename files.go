package sil

import (
	"fmt"
	"io/ioutil"
)

// Write writes a SIL to a file
func (s *SIL) Write(filename string) error {
	// create the bytes of the SIL file
	d, err := s.Bytes()
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %v", err)
	}

	err = ioutil.WriteFile(filename, d, 0777)
	// return the error details
	if err != nil {
		return fmt.Errorf("ioutil error: %v", err)
	}
	return nil
}
