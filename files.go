package sil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Write writes a SIL to a file
func (s *SIL) Write(filename string, include bool) error {
	// create the bytes of the SIL file
	d, err := s.Marshal(include)
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %v", err)
	}

	err = ioutil.WriteFile(filename, d, 0666)
	// return the error details
	if err != nil {
		return fmt.Errorf("ioutil error: %v", err)
	}
	return nil
}

// JSON Creates a JSON file of the SIL file
func (s *SIL) JSON(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
