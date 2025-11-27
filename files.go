package sil

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// Write writes a SIL to a file
// Uses context.Background() for backwards compatibility
func (s *SIL) Write(filename string, include bool, archive bool) error {
	return s.WriteContext(context.Background(), filename, include, archive)
}

// WriteContext writes a SIL to a file with context support
// The context can be used to cancel the operation
func (s *SIL) WriteContext(ctx context.Context, filename string, include bool, archive bool) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

	// create the bytes of the SIL file
	d, err := s.MarshalWithContext(ctx)
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %v", err)
	}

	// Check context after marshaling
	if err := ctx.Err(); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %s with err: %v", filename, err)
	}
	// if we are manipulating the archive bit first SET archive bit
	// This really might not be needed but needs to be tested over UNC paths
	// the file might need to be closed first to allow the system to set the bit properly
	if archive {
		err = setArchive(filename)
		if err != nil {
			f.Close()
			return fmt.Errorf("error trying to set archive bit for %s with err: %v", filename, err)
		}
	}

	// Check context before writing
	if err := ctx.Err(); err != nil {
		f.Close()
		return err
	}

	// Write the bytes to the file, this returns an int of bytes written
	i, err := f.Write(d)

	if err != nil {
		f.Close()
		return fmt.Errorf("could not write bytes to file %s with err: %w", filename, err)
	} else if i != len(d) {
		f.Close()
		return fmt.Errorf("number of bytes writte to %s did not match length of sil bytes", filename)
	}
	// Close the file, otherwise the archive bit cannot be unset
	err = f.Close()
	if err != nil {
		return fmt.Errorf("error closing the file %s", filename)
	}
	// if we are manipulating the archive bit unset the archive bit
	if archive {
		err = unsetArchive(filename)
		if err != nil {
			return fmt.Errorf("error trying to set archive bit for %s with err: %w", filename, err)
		}
	}

	return nil
}

// Write writes a Multi to a file
// Uses context.Background() for backwards compatibility
func (m *Multi) Write(filename string, archive bool) error {
	return m.WriteContext(context.Background(), filename, archive)
}

// WriteContext writes a Multi to a file with context support
// The context can be used to cancel the operation
func (m *Multi) WriteContext(ctx context.Context, filename string, archive bool) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

	// create the bytes of the SIL file
	d, err := m.MarshalWithContext(ctx)
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %w", err)
	}

	// Check context after marshaling
	if err := ctx.Err(); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %s with err: %w", filename, err)
	}
	// if we are manipulating the archive bit first SET archive bit
	// This really might not be needed but needs to be tested over UNC paths
	// the file might need to be closed first to allow the system to set the bit properly
	if archive {
		err = setArchive(filename)
		if err != nil {
			f.Close()
			return fmt.Errorf("error trying to set archive bit for %s with err: %w", filename, err)
		}
	}

	// Check context before writing
	if err := ctx.Err(); err != nil {
		f.Close()
		return err
	}

	// Write the bytes to the file, this returns an int of bytes written
	i, err := f.Write(d)

	if err != nil {
		f.Close()
		return fmt.Errorf("could not write bytes to file %s with err: %w", filename, err)
	} else if i != len(d) {
		f.Close()
		return fmt.Errorf("number of bytes writte to %s did not match length of sil bytes", filename)
	}
	// Close the file, otherwise the archive bit cannot be unset
	err = f.Close()
	if err != nil {
		return fmt.Errorf("error closing the file %s", filename)
	}
	// if we are manipulating the archive bit unset the archive bit
	if archive {
		err = unsetArchive(filename)
		if err != nil {
			return fmt.Errorf("error trying to set archive bit for %s with err: %w", filename, err)
		}
	}

	return nil
}

// JSON Creates a JSON file of the SIL file
// Uses context.Background() for backwards compatibility
func (s *SIL) JSON(filename string) error {
	return s.JSONContext(context.Background(), filename)
}

// JSONContext creates a JSON file of the SIL file with context support
// The context can be used to cancel the operation
func (s *SIL) JSONContext(ctx context.Context, filename string) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Check context before encoding
	if err := ctx.Err(); err != nil {
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
