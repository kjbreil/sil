package sil

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

// rowResult holds the result of processing a single row
type rowResult struct {
	index int
	row   row
	err   error
}

// splitConcurrent processes rows concurrently using a worker pool
// It maintains order by using indexed results
func splitConcurrent(rows []interface{}, include bool) (map[string]section, error) {
	if len(rows) == 0 {
		return make(map[string]section), nil
	}

	// For small datasets, use sequential processing (overhead not worth it)
	if len(rows) < 100 {
		return split(rows, include)
	}

	numWorkers := runtime.GOMAXPROCS(0)
	if numWorkers > len(rows) {
		numWorkers = len(rows)
	}

	// Channel for work items (row indices)
	workChan := make(chan int, len(rows))
	// Channel for results
	resultChan := make(chan rowResult, len(rows))

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range workChan {
				var r row
				err := r.make(rows[idx], include)
				resultChan <- rowResult{
					index: idx,
					row:   r,
					err:   err,
				}
			}
		}()
	}

	// Send work
	for i := range rows {
		workChan <- i
	}
	close(workChan)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results in order
	processedRows := make([]row, len(rows))
	for result := range resultChan {
		if result.err != nil {
			return nil, result.err
		}
		processedRows[result.index] = result.row
	}

	// Build sections from processed rows
	secs := make(map[string]section)
	for i := range processedRows {
		var key string
		for x := range processedRows[i].elems {
			key = key + *processedRows[i].elems[x].name
		}
		secs[key] = append(secs[key], processedRows[i])
	}

	return secs, nil
}

// marshalResult holds the result of marshaling a single SIL
type marshalResult struct {
	name string
	data []byte
	err  error
}

// MarshalConcurrent creates the SIL structure from the information in the Multi
// using concurrent processing for multiple tables
func (m Multi) MarshalConcurrent() (data []byte, err error) {
	if len(m) == 0 {
		return data, nil
	}

	// For single table, use sequential
	if len(m) == 1 {
		return m.Marshal()
	}

	// Assign prefix to all SILs (must be done sequentially first)
	prefix := generatePrefix()
	for _, s := range m {
		s.prefix = prefix
	}

	// Create channel for results
	resultChan := make(chan marshalResult, len(m))

	// Start goroutine for each table
	var wg sync.WaitGroup
	for name, s := range m {
		wg.Add(1)
		go func(name string, s *SIL) {
			defer wg.Done()
			b, err := s.Marshal(false)
			resultChan <- marshalResult{
				name: name,
				data: b,
				err:  err,
			}
		}(name, s)
	}

	// Wait and close channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results - order doesn't matter for Multi since
	// each batch has its own header
	var allData []byte
	for result := range resultChan {
		if result.err != nil {
			return nil, result.err
		}
		allData = append(allData, result.data...)
	}

	return allData, nil
}

// MarshalWithOptions provides options for marshaling
type MarshalOptions struct {
	Concurrent bool
	Include    bool
}

// MarshalWithOpts marshals with specified options
func (s *SIL) MarshalWithOpts(opts MarshalOptions) ([]byte, error) {
	if opts.Concurrent {
		return s.marshalConcurrent(opts.Include)
	}
	return s.Marshal(opts.Include)
}

// marshalConcurrent uses concurrent processing for large datasets
func (s *SIL) marshalConcurrent(include bool) (data []byte, err error) {
	if s.View.Name == "" {
		return data, errViewNameNotSet
	}

	if !include {
		include = s.Include
	}

	// Use concurrent split for large datasets
	secs, err := splitConcurrent(s.View.Data, include)
	if err != nil {
		return []byte{}, err
	}

	for _, sec := range secs {
		if s.Header.Identifier == "" {
			s.Header.Identifier = batchNum(s.prefix)
		}
		if s.View.Action != "LOAD" {
			data = append(data, s.Header.insert()...)
			data = append(data, s.Header.row()...)
		}
		data = append(data, sec.create(&s.View)...)
		data = append(data, []byte(crlf)...)
	}

	data = append(data, s.Footer.bytes()...)

	data, err = encodeWindows1252(data)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// generatePrefix generates a random prefix for batch grouping
func generatePrefix() int {
	// Using crypto/rand would be better but keeping compatibility
	// with existing behavior that uses math/rand
	return randIntn(100)
}

// randIntn is a wrapper for random number generation
// This allows for easier testing and potential future improvements
var randIntn = func(n int) int {
	// Use the existing rand from math/rand that's seeded in Make/New
	return internalRandIntn(n)
}

// WriteResult holds the result of writing a file
type WriteResult struct {
	Name     string
	Filename string
	Err      error
}

// WriteSeparateConcurrent writes each table in Multi to separate files concurrently
// The filenames are generated using the provided pattern function
func (m Multi) WriteSeparateConcurrent(filenameFunc func(tableName string) string, archive bool) []WriteResult {
	if len(m) == 0 {
		return nil
	}

	results := make(chan WriteResult, len(m))
	var wg sync.WaitGroup

	for name, s := range m {
		wg.Add(1)
		go func(name string, s *SIL) {
			defer wg.Done()
			filename := filenameFunc(name)
			err := s.Write(filename, false, archive)
			results <- WriteResult{
				Name:     name,
				Filename: filename,
				Err:      err,
			}
		}(name, s)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allResults []WriteResult
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}

// WriteConcurrent writes Multi to a single file using concurrent marshaling
func (m *Multi) WriteConcurrent(filename string, archive bool) error {
	d, err := m.MarshalConcurrent()
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %w", err)
	}

	return writeFile(filename, d, archive)
}

// writeFile is a helper function for writing data to a file with optional archive bit manipulation
func writeFile(filename string, data []byte, archive bool) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %s with err: %w", filename, err)
	}

	if archive {
		err = setArchive(filename)
		if err != nil {
			f.Close()
			return fmt.Errorf("error trying to set archive bit for %s with err: %w", filename, err)
		}
	}

	i, err := f.Write(data)
	if err != nil {
		f.Close()
		return fmt.Errorf("could not write bytes to file %s with err: %w", filename, err)
	}
	if i != len(data) {
		f.Close()
		return fmt.Errorf("number of bytes written to %s did not match length of sil bytes", filename)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("error closing the file %s", filename)
	}

	if archive {
		err = unsetArchive(filename)
		if err != nil {
			return fmt.Errorf("error trying to unset archive bit for %s with err: %w", filename, err)
		}
	}

	return nil
}
