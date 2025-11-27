package sil

import (
	"context"
	"strings"
	"sync"
	"testing"
)

// ConcurrentTestStruct for testing concurrent operations
type ConcurrentTestStruct struct {
	ID   string `sil:"F01"`
	Name string `sil:"F02" default:"DEFAULT"`
}

// ConcurrentTestStructOptional with optional fields
type ConcurrentTestStructOptional struct {
	ID    string  `sil:"F01"`
	Value *string `sil:"F02"`
}

func TestSplitConcurrent_SmallDataset(t *testing.T) {
	// Small datasets should use sequential processing
	rows := make([]interface{}, 50)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	secs, err := splitConcurrent(context.Background(), rows, false)
	if err != nil {
		t.Fatalf("splitConcurrent returned error: %v", err)
	}

	if len(secs) != 1 {
		t.Fatalf("expected 1 section, got %d", len(secs))
	}

	for _, sec := range secs {
		if len(sec) != 50 {
			t.Fatalf("expected 50 rows, got %d", len(sec))
		}
	}
}

func TestSplitConcurrent_LargeDataset(t *testing.T) {
	// Large dataset should use concurrent processing
	const numRows = 500
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	secs, err := splitConcurrent(context.Background(), rows, false)
	if err != nil {
		t.Fatalf("splitConcurrent returned error: %v", err)
	}

	if len(secs) != 1 {
		t.Fatalf("expected 1 section, got %d", len(secs))
	}

	for _, sec := range secs {
		if len(sec) != numRows {
			t.Fatalf("expected %d rows, got %d", numRows, len(sec))
		}
	}
}

func TestSplitConcurrent_MixedFieldCompositions(t *testing.T) {
	const numRows = 500
	rows := make([]interface{}, numRows)
	for i := range rows {
		if i%2 == 0 {
			rows[i] = ConcurrentTestStructOptional{ID: "id"}
		} else {
			val := "value"
			rows[i] = ConcurrentTestStructOptional{ID: "id", Value: &val}
		}
	}

	secs, err := splitConcurrent(context.Background(), rows, false)
	if err != nil {
		t.Fatalf("splitConcurrent returned error: %v", err)
	}

	// Should have 2 sections (different field compositions)
	if len(secs) != 2 {
		t.Fatalf("expected 2 sections, got %d", len(secs))
	}

	totalRows := 0
	for _, sec := range secs {
		totalRows += len(sec)
	}
	if totalRows != numRows {
		t.Fatalf("expected %d total rows, got %d", numRows, totalRows)
	}
}

func TestSplitConcurrent_Empty(t *testing.T) {
	rows := []interface{}{}

	secs, err := splitConcurrent(context.Background(), rows, false)
	if err != nil {
		t.Fatalf("splitConcurrent returned error: %v", err)
	}

	if len(secs) != 0 {
		t.Fatalf("expected 0 sections, got %d", len(secs))
	}
}

func TestSplitConcurrent_Error(t *testing.T) {
	rows := []interface{}{
		RequiredFieldStruct{}, // Missing required field
	}

	_, err := splitConcurrent(context.Background(), rows, false)
	if err == nil {
		t.Fatal("expected error for missing required field")
	}
}

func TestSplitConcurrent_PreservesOrder(t *testing.T) {
	// Test that order is preserved for rows with same field composition
	const numRows = 200
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: string(rune('A' + (i % 26)))}
	}

	// Run multiple times to catch race conditions
	for run := 0; run < 10; run++ {
		secs, err := splitConcurrent(context.Background(), rows, false)
		if err != nil {
			t.Fatalf("run %d: splitConcurrent returned error: %v", run, err)
		}

		// Verify all rows are present
		for _, sec := range secs {
			if len(sec) != numRows {
				t.Fatalf("run %d: expected %d rows, got %d", run, numRows, len(sec))
			}

			// Verify order is preserved
			for i, r := range sec {
				expected := string(rune('A' + (i % 26)))
				// Extract the ID value from the row
				for _, elem := range r.elems {
					if *elem.name == "F01" {
						// Remove quotes from value
						actual := strings.Trim(*elem.data, "'")
						if actual != expected {
							t.Errorf("run %d: row %d: expected ID %s, got %s", run, i, expected, actual)
						}
						break
					}
				}
			}
		}
	}
}

func TestMultiMarshalConcurrent_Empty(t *testing.T) {
	m := make(Multi)

	data, err := m.MarshalConcurrent(context.Background())
	if err != nil {
		t.Fatalf("MarshalConcurrent returned error: %v", err)
	}

	if len(data) != 0 {
		t.Errorf("expected empty data, got %d bytes", len(data))
	}
}

func TestMultiMarshalConcurrent_SingleTable(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", ConcurrentTestStruct{})
	m.AppendView("TEST", ConcurrentTestStruct{ID: "001"})

	data, err := m.MarshalConcurrent(context.Background())
	if err != nil {
		t.Fatalf("MarshalConcurrent returned error: %v", err)
	}

	output := string(data)
	if !strings.Contains(output, "TEST") {
		t.Error("expected TEST in output")
	}
}

func TestMultiMarshalConcurrent_MultipleTables(t *testing.T) {
	m := make(Multi)

	// Create multiple tables
	for i := 0; i < 5; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		for j := 0; j < 10; j++ {
			m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
		}
	}

	data, err := m.MarshalConcurrent(context.Background())
	if err != nil {
		t.Fatalf("MarshalConcurrent returned error: %v", err)
	}

	output := string(data)

	// Verify all tables are present
	for i := 0; i < 5; i++ {
		tableName := string(rune('A' + i))
		if !strings.Contains(output, tableName) {
			t.Errorf("expected table %s in output", tableName)
		}
	}

	// Should have 5 CREATE VIEW statements
	count := strings.Count(output, "CREATE VIEW")
	if count != 5 {
		t.Errorf("expected 5 CREATE VIEW statements, got %d", count)
	}
}

func TestMultiMarshalConcurrent_Error(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", RequiredFieldStruct{})
	m.AppendView("TEST", RequiredFieldStruct{}) // Missing required field

	_, err := m.MarshalConcurrent(context.Background())
	if err == nil {
		t.Fatal("expected error for missing required field")
	}
}

func TestMarshalWithOpts_Sequential(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"
	s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "001"})

	opts := MarshalOptions{
		Concurrent: false,
		Include:    false,
	}

	data, err := s.MarshalWithOpts(opts)
	if err != nil {
		t.Fatalf("MarshalWithOpts returned error: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

func TestMarshalWithOpts_Concurrent(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"

	// Add enough rows to trigger concurrent processing
	for i := 0; i < 200; i++ {
		s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "row"})
	}

	opts := MarshalOptions{
		Concurrent: true,
		Include:    false,
	}

	data, err := s.MarshalWithOpts(opts)
	if err != nil {
		t.Fatalf("MarshalWithOpts returned error: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty data")
	}

	output := string(data)
	if !strings.Contains(output, "TEST") {
		t.Error("expected TEST in output")
	}
}

// TestConcurrentSafety verifies no race conditions
func TestConcurrentSafety(t *testing.T) {
	const numRows = 500
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	// Run multiple concurrent calls
	var wg sync.WaitGroup
	errors := make(chan error, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := splitConcurrent(context.Background(), rows, false)
			if err != nil {
				errors <- err
			}
		}()
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		t.Errorf("concurrent execution error: %v", err)
	}
}

// Benchmarks for comparing sequential vs concurrent
func BenchmarkSplit_100Rows(b *testing.B) {
	rows := make([]interface{}, 100)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplitConcurrent_100Rows(b *testing.B) {
	ctx := context.Background()
	rows := make([]interface{}, 100)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = splitConcurrent(ctx, rows, false)
	}
}

func BenchmarkSplit_500Rows(b *testing.B) {
	rows := make([]interface{}, 500)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplitConcurrent_500Rows(b *testing.B) {
	ctx := context.Background()
	rows := make([]interface{}, 500)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = splitConcurrent(ctx, rows, false)
	}
}

func BenchmarkSplit_1000Rows(b *testing.B) {
	rows := make([]interface{}, 1000)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplitConcurrent_1000Rows(b *testing.B) {
	ctx := context.Background()
	rows := make([]interface{}, 1000)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = splitConcurrent(ctx, rows, false)
	}
}

func BenchmarkSplit_5000Rows(b *testing.B) {
	rows := make([]interface{}, 5000)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplitConcurrent_5000Rows(b *testing.B) {
	ctx := context.Background()
	rows := make([]interface{}, 5000)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = splitConcurrent(ctx, rows, false)
	}
}

func BenchmarkMultiMarshal_5Tables(b *testing.B) {
	m := make(Multi)
	for i := 0; i < 5; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		for j := 0; j < 100; j++ {
			m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Marshal()
	}
}

func BenchmarkMultiMarshalConcurrent_5Tables(b *testing.B) {
	ctx := context.Background()
	m := make(Multi)
	for i := 0; i < 5; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		for j := 0; j < 100; j++ {
			m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.MarshalConcurrent(ctx)
	}
}

func BenchmarkMultiMarshal_10Tables(b *testing.B) {
	m := make(Multi)
	for i := 0; i < 10; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		for j := 0; j < 100; j++ {
			m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Marshal()
	}
}

func BenchmarkMultiMarshalConcurrent_10Tables(b *testing.B) {
	ctx := context.Background()
	m := make(Multi)
	for i := 0; i < 10; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		for j := 0; j < 100; j++ {
			m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.MarshalConcurrent(ctx)
	}
}

// Tests for file write functions
func TestWriteSeparateConcurrent_Empty(t *testing.T) {
	m := make(Multi)

	results := m.WriteSeparateConcurrent(func(name string) string {
		return name + ".sil"
	}, false)

	if len(results) != 0 {
		t.Errorf("expected no results for empty Multi, got %d", len(results))
	}
}

func TestWriteSeparateConcurrent_MultipleTables(t *testing.T) {
	m := make(Multi)
	for i := 0; i < 3; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
	}

	// Use temp dir for file writing
	tempDir := t.TempDir()
	results := m.WriteSeparateConcurrent(func(name string) string {
		return tempDir + "/" + name + ".sil"
	}, false)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	// Verify all writes succeeded
	for _, result := range results {
		if result.Err != nil {
			t.Errorf("write failed for %s: %v", result.Name, result.Err)
		}
	}
}

func TestWriteConcurrent(t *testing.T) {
	m := make(Multi)
	for i := 0; i < 3; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
	}

	tempDir := t.TempDir()
	filename := tempDir + "/test.sil"

	err := m.WriteConcurrent(filename, false)
	if err != nil {
		t.Fatalf("WriteConcurrent returned error: %v", err)
	}
}

func TestWriteFile(t *testing.T) {
	tempDir := t.TempDir()
	filename := tempDir + "/test.sil"
	data := []byte("test data")

	err := writeFile(filename, data, false)
	if err != nil {
		t.Fatalf("writeFile returned error: %v", err)
	}
}
