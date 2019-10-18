package sil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Write writes a SIL to a file
func (s *SIL) Write(filename string, include bool, archive bool) error {
	// create the bytes of the SIL file
	d, err := s.Marshal(include)
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %v", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %s with err: %v", filename, err)
	}
	// if we are manipulating the archive bit first SET archive bit
	if archive {
		err = setArchive(filename)
		if err != nil {
			return fmt.Errorf("error trying to set archive bit for %s with err: %v", filename, err)
		}
	}

	// Write the bytes to the file, this returns an int of bytes written
	i, err := f.Write(d)

	if err != nil {
		return fmt.Errorf("could not write bytes to file %s with err: %v", filename, err)
	} else if i != len(d) {
		return fmt.Errorf("number of bytes writte to %s did not match length of sil bytes", filename)
	}
	// if we are manipulating the archive bit unset the archive bit
	if archive {
		err = unsetArchive(filename)
		if err != nil {
			return fmt.Errorf("error trying to set archive bit for %s with err: %v", filename, err)
		}
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
