package sil

import (
	"context"
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
// The context can be used to cancel the operation
func splitConcurrent(ctx context.Context, rows []interface{}, include bool) (map[string]section, error) {
	if len(rows) == 0 {
		return make(map[string]section), nil
	}

	// Check context before starting
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// For small datasets, use sequential processing (overhead not worth it)
	if len(rows) < 100 {
		return splitWithContext(ctx, rows, include)
	}

	numWorkers := runtime.GOMAXPROCS(0)
	if numWorkers > len(rows) {
		numWorkers = len(rows)
	}

	// Channel for work items (row indices)
	workChan := make(chan int, len(rows))
	// Channel for results
	resultChan := make(chan rowResult, len(rows))
	// Channel to signal workers to stop
	doneChan := make(chan struct{})

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-doneChan:
					return
				case idx, ok := <-workChan:
					if !ok {
						return
					}
					var r row
					err := r.make(rows[idx], include)
					select {
					case resultChan <- rowResult{
						index: idx,
						row:   r,
						err:   err,
					}:
					case <-doneChan:
						return
					}
				}
			}
		}()
	}

	// Send work (with context cancellation support)
	go func() {
		defer close(workChan)
		for i := range rows {
			select {
			case <-ctx.Done():
				return
			case workChan <- i:
			}
		}
	}()

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results in order
	processedRows := make([]row, len(rows))
	var firstErr error
	collected := 0

	for {
		select {
		case <-ctx.Done():
			close(doneChan)
			return nil, ctx.Err()
		case result, ok := <-resultChan:
			if !ok {
				// Channel closed, all results collected
				goto buildSections
			}
			if result.err != nil && firstErr == nil {
				firstErr = result.err
			}
			processedRows[result.index] = result.row
			collected++
		}
	}

buildSections:
	if firstErr != nil {
		return nil, firstErr
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

// splitWithContext is a context-aware version of split for small datasets
func splitWithContext(ctx context.Context, rows []interface{}, include bool) (map[string]section, error) {
	var ssec section

	// take every row and reflect it
	for i := range rows {
		// Check context periodically (every row for small datasets)
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		var r row
		err := r.make(rows[i], include)
		if err != nil {
			return nil, err
		}
		ssec = append(ssec, r)
	}

	secs := make(map[string]section)

	for i := range ssec {
		// make the name of the section for the map based on the fields
		var key string
		for x := range ssec[i].elems {
			key = key + *ssec[i].elems[x].name
		}

		secs[key] = append(secs[key], ssec[i])
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
// The context can be used to cancel the operation
func (m Multi) MarshalConcurrent(ctx context.Context) (data []byte, err error) {
	if len(m) == 0 {
		return data, nil
	}

	// Check context before starting
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// For single table, use sequential
	if len(m) == 1 {
		return m.MarshalWithContext(ctx)
	}

	// Assign prefix to all SILs (must be done sequentially first)
	prefix := generatePrefix()
	for _, s := range m {
		s.prefix = prefix
	}

	// Create channel for results
	resultChan := make(chan marshalResult, len(m))
	// Channel to signal workers to stop
	doneChan := make(chan struct{})

	// Start goroutine for each table
	var wg sync.WaitGroup
	for name, s := range m {
		wg.Add(1)
		go func(name string, s *SIL) {
			defer wg.Done()
			// Check context before marshaling
			select {
			case <-doneChan:
				return
			default:
			}
			b, err := s.MarshalWithContext(ctx)
			select {
			case resultChan <- marshalResult{
				name: name,
				data: b,
				err:  err,
			}:
			case <-doneChan:
				return
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
	var firstErr error
	for {
		select {
		case <-ctx.Done():
			close(doneChan)
			return nil, ctx.Err()
		case result, ok := <-resultChan:
			if !ok {
				// Channel closed, all results collected
				if firstErr != nil {
					return nil, firstErr
				}
				return allData, nil
			}
			if result.err != nil && firstErr == nil {
				firstErr = result.err
			}
			if result.err == nil {
				allData = append(allData, result.data...)
			}
		}
	}
}

// MarshalWithOptions provides options for marshaling
type MarshalOptions struct {
	Concurrent bool
	Include    bool
}

// MarshalWithOpts marshals with specified options
// Uses context.Background() for backwards compatibility
func (s *SIL) MarshalWithOpts(opts MarshalOptions) ([]byte, error) {
	return s.MarshalWithOptsContext(context.Background(), opts)
}

// MarshalWithOptsContext marshals with specified options and context support
func (s *SIL) MarshalWithOptsContext(ctx context.Context, opts MarshalOptions) ([]byte, error) {
	if opts.Concurrent {
		return s.marshalConcurrent(ctx, opts.Include)
	}
	return s.MarshalWithContext(ctx)
}

// marshalConcurrent uses concurrent processing for large datasets
func (s *SIL) marshalConcurrent(ctx context.Context, include bool) (data []byte, err error) {
	if s.View.Name == "" {
		return data, errViewNameNotSet
	}

	// Check context before starting
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if !include {
		include = s.Include
	}

	// Use concurrent split for large datasets
	secs, err := splitConcurrent(ctx, s.View.Data, include)
	if err != nil {
		return []byte{}, err
	}

	// Check context after split
	if err := ctx.Err(); err != nil {
		return nil, err
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
// Uses context.Background() for backwards compatibility
func (m Multi) WriteSeparateConcurrent(filenameFunc func(tableName string) string, archive bool) []WriteResult {
	return m.WriteSeparateConcurrentContext(context.Background(), filenameFunc, archive)
}

// WriteSeparateConcurrentContext writes each table in Multi to separate files concurrently
// The filenames are generated using the provided pattern function
// The context can be used to cancel the operation
func (m Multi) WriteSeparateConcurrentContext(ctx context.Context, filenameFunc func(tableName string) string, archive bool) []WriteResult {
	if len(m) == 0 {
		return nil
	}

	// Check context before starting
	if err := ctx.Err(); err != nil {
		return []WriteResult{{Err: err}}
	}

	results := make(chan WriteResult, len(m))
	// Channel to signal workers to stop
	doneChan := make(chan struct{})
	var wg sync.WaitGroup

	for name, s := range m {
		wg.Add(1)
		go func(name string, s *SIL) {
			defer wg.Done()
			// Check if we should stop
			select {
			case <-doneChan:
				return
			default:
			}
			filename := filenameFunc(name)
			err := s.WriteContext(ctx, filename, false, archive)
			select {
			case results <- WriteResult{
				Name:     name,
				Filename: filename,
				Err:      err,
			}:
			case <-doneChan:
				return
			}
		}(name, s)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allResults []WriteResult
	for {
		select {
		case <-ctx.Done():
			close(doneChan)
			allResults = append(allResults, WriteResult{Err: ctx.Err()})
			return allResults
		case result, ok := <-results:
			if !ok {
				return allResults
			}
			allResults = append(allResults, result)
		}
	}
}

// WriteConcurrent writes Multi to a single file using concurrent marshaling
// Uses context.Background() for backwards compatibility
func (m *Multi) WriteConcurrent(filename string, archive bool) error {
	return m.WriteConcurrentContext(context.Background(), filename, archive)
}

// WriteConcurrentContext writes Multi to a single file using concurrent marshaling
// The context can be used to cancel the operation
func (m *Multi) WriteConcurrentContext(ctx context.Context, filename string, archive bool) error {
	d, err := m.MarshalConcurrent(ctx)
	if err != nil {
		return fmt.Errorf("sil bytes conversion error: %w", err)
	}

	return writeFileContext(ctx, filename, d, archive)
}

// writeFile is a helper function for writing data to a file with optional archive bit manipulation
// Uses context.Background() for backwards compatibility
func writeFile(filename string, data []byte, archive bool) error {
	return writeFileContext(context.Background(), filename, data, archive)
}

// writeFileContext is a helper function for writing data to a file with optional archive bit manipulation
// The context can be used to cancel the operation
func writeFileContext(ctx context.Context, filename string, data []byte, archive bool) error {
	// Check context before starting
	if err := ctx.Err(); err != nil {
		return err
	}

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

	// Check context before writing
	if err := ctx.Err(); err != nil {
		f.Close()
		return err
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
