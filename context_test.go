package sil

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestSplitConcurrent_ContextCancellation tests that splitConcurrent respects context cancellation
func TestSplitConcurrent_ContextCancellation(t *testing.T) {
	// Create a large dataset to ensure processing takes some time
	const numRows = 1000
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := splitConcurrent(ctx, rows, false)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestSplitConcurrent_ContextTimeout tests that splitConcurrent respects context timeout
func TestSplitConcurrent_ContextTimeout(t *testing.T) {
	// Create a large dataset
	const numRows = 10000
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	// Create a context with an extremely short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	// Give it a moment to expire
	time.Sleep(1 * time.Millisecond)

	_, err := splitConcurrent(ctx, rows, false)
	if err == nil {
		t.Fatal("expected error from timed out context")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected context.DeadlineExceeded error, got: %v", err)
	}
}

// TestSplitWithContext_Cancellation tests the sequential path with context cancellation
func TestSplitWithContext_Cancellation(t *testing.T) {
	// Small dataset to trigger sequential processing
	const numRows = 50
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := splitWithContext(ctx, rows, false)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestMarshalConcurrent_ContextCancellation tests that MarshalConcurrent respects context cancellation
func TestMarshalConcurrent_ContextCancellation(t *testing.T) {
	m := make(Multi)
	// Create multiple tables
	for i := 0; i < 5; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		for j := 0; j < 100; j++ {
			m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
		}
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := m.MarshalConcurrent(ctx)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestMarshalWithContext_Cancellation tests that MarshalWithContext respects context cancellation
func TestMarshalWithContext_Cancellation(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"
	for i := 0; i < 100; i++ {
		s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "row"})
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := s.MarshalWithContext(ctx)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestWriteContext_Cancellation tests that WriteContext respects context cancellation
func TestWriteContext_Cancellation(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"
	s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "row"})

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tempDir := t.TempDir()
	err := s.WriteContext(ctx, tempDir+"/test.sil", false, false)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestMultiWriteContext_Cancellation tests that Multi.WriteContext respects context cancellation
func TestMultiWriteContext_Cancellation(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", ConcurrentTestStruct{})
	m.AppendView("TEST", ConcurrentTestStruct{ID: "row"})

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tempDir := t.TempDir()
	err := m.WriteContext(ctx, tempDir+"/test.sil", false)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestWriteConcurrentContext_Cancellation tests that WriteConcurrentContext respects context cancellation
func TestWriteConcurrentContext_Cancellation(t *testing.T) {
	m := make(Multi)
	for i := 0; i < 3; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tempDir := t.TempDir()
	err := m.WriteConcurrentContext(ctx, tempDir+"/test.sil", false)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestWriteSeparateConcurrentContext_Cancellation tests that WriteSeparateConcurrentContext respects context cancellation
func TestWriteSeparateConcurrentContext_Cancellation(t *testing.T) {
	m := make(Multi)
	for i := 0; i < 3; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tempDir := t.TempDir()
	results := m.WriteSeparateConcurrentContext(ctx, func(name string) string {
		return tempDir + "/" + name + ".sil"
	}, false)

	// Should have at least one error result
	hasError := false
	for _, result := range results {
		if result.Err != nil {
			hasError = true
			if !errors.Is(result.Err, context.Canceled) {
				t.Errorf("expected context.Canceled error, got: %v", result.Err)
			}
		}
	}
	if !hasError {
		t.Error("expected at least one error from cancelled context")
	}
}

// TestMarshalWithOptsContext_Cancellation tests that MarshalWithOptsContext respects context cancellation
func TestMarshalWithOptsContext_Cancellation(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"
	for i := 0; i < 200; i++ {
		s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "row"})
	}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := s.MarshalWithOptsContext(ctx, MarshalOptions{Concurrent: true})
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestContextCancellation_MidOperation tests cancellation during processing
func TestContextCancellation_MidOperation(t *testing.T) {
	// Create a large dataset
	const numRows = 5000
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after a short delay
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	// This should either complete successfully or return a context error
	_, err := splitConcurrent(ctx, rows, false)
	if err != nil {
		if !errors.Is(err, context.Canceled) {
			t.Errorf("unexpected error type: %v", err)
		}
	}
	// It's okay if it succeeds (fast machine) or fails with context.Canceled
}

// TestMarshalWithContext_ValidContext tests that MarshalWithContext works with valid context
func TestMarshalWithContext_ValidContext(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"
	s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "001"})

	ctx := context.Background()
	data, err := s.MarshalWithContext(ctx)
	if err != nil {
		t.Fatalf("MarshalWithContext with valid context returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

// TestMarshalContextInclude tests the MarshalContextInclude method
func TestMarshalContextInclude(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"
	s.View.Data = append(s.View.Data, ConcurrentTestStruct{ID: "001"})

	ctx := context.Background()
	data, err := s.MarshalContextInclude(ctx, false)
	if err != nil {
		t.Fatalf("MarshalContextInclude returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

// TestWriteFileContext_ValidContext tests writeFileContext with valid context
func TestWriteFileContext_ValidContext(t *testing.T) {
	ctx := context.Background()
	tempDir := t.TempDir()
	filename := tempDir + "/test.sil"
	data := []byte("test data")

	err := writeFileContext(ctx, filename, data, false)
	if err != nil {
		t.Fatalf("writeFileContext with valid context returned error: %v", err)
	}
}

// TestWriteFileContext_Cancellation tests that writeFileContext respects context cancellation
func TestWriteFileContext_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tempDir := t.TempDir()
	filename := tempDir + "/test.sil"
	data := []byte("test data")

	err := writeFileContext(ctx, filename, data, false)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestJSONContext_Cancellation tests that JSONContext respects context cancellation
func TestJSONContext_Cancellation(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tempDir := t.TempDir()
	err := s.JSONContext(ctx, tempDir+"/test.json")
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestJSONContext_ValidContext tests that JSONContext works with valid context
func TestJSONContext_ValidContext(t *testing.T) {
	var s SIL
	s.View.Name = "TEST"

	ctx := context.Background()
	tempDir := t.TempDir()
	err := s.JSONContext(ctx, tempDir+"/test.json")
	if err != nil {
		t.Fatalf("JSONContext with valid context returned error: %v", err)
	}
}

// TestMultiMarshalWithContext_Cancellation tests that Multi.MarshalWithContext respects context cancellation
func TestMultiMarshalWithContext_Cancellation(t *testing.T) {
	m := make(Multi)
	for i := 0; i < 3; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, ConcurrentTestStruct{})
		m.AppendView(tableName, ConcurrentTestStruct{ID: "row"})
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := m.MarshalWithContext(ctx)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled error, got: %v", err)
	}
}

// TestMultiMarshalWithContext_ValidContext tests that Multi.MarshalWithContext works with valid context
func TestMultiMarshalWithContext_ValidContext(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", ConcurrentTestStruct{})
	m.AppendView("TEST", ConcurrentTestStruct{ID: "row"})

	ctx := context.Background()
	data, err := m.MarshalWithContext(ctx)
	if err != nil {
		t.Fatalf("MarshalWithContext with valid context returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

// TestContextDeadlineExceeded tests handling of context deadline exceeded
func TestContextDeadlineExceeded(t *testing.T) {
	const numRows = 1000
	rows := make([]interface{}, numRows)
	for i := range rows {
		rows[i] = ConcurrentTestStruct{ID: "id"}
	}

	// Create a context with an expired deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-1*time.Hour))
	defer cancel()

	_, err := splitConcurrent(ctx, rows, false)
	if err == nil {
		t.Fatal("expected error from expired deadline context")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected context.DeadlineExceeded error, got: %v", err)
	}
}
