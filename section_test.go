package sil

import (
	"strings"
	"testing"
)

// TestStruct is a simple test struct for testing section functionality
type TestStruct struct {
	Field1 string `sil:"F01"`
	Field2 int    `sil:"F02" default:"10"`
	Field3 string `sil:"F03" default:"DEFAULT"`
}

// TestStructOptional has optional pointer fields
type TestStructOptional struct {
	Field1 string  `sil:"F01"`
	Field2 *string `sil:"F02"`
	Field3 *int    `sil:"F03"`
}

func TestSplit_SingleRow(t *testing.T) {
	rows := []interface{}{
		TestStruct{Field1: "value1"},
	}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	if len(secs) != 1 {
		t.Fatalf("expected 1 section, got %d", len(secs))
	}

	// Check that section contains one row
	for _, sec := range secs {
		if len(sec) != 1 {
			t.Fatalf("expected section to have 1 row, got %d", len(sec))
		}
	}
}

func TestSplit_MultipleRowsSameFields(t *testing.T) {
	rows := []interface{}{
		TestStruct{Field1: "value1"},
		TestStruct{Field1: "value2"},
		TestStruct{Field1: "value3"},
	}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	if len(secs) != 1 {
		t.Fatalf("expected 1 section (same fields), got %d", len(secs))
	}

	// Check that section contains all rows
	for _, sec := range secs {
		if len(sec) != 3 {
			t.Fatalf("expected section to have 3 rows, got %d", len(sec))
		}
	}
}

func TestSplit_DifferentFieldCompositions(t *testing.T) {
	opt1 := "optional1"
	opt2 := 42

	rows := []interface{}{
		TestStructOptional{Field1: "value1"},              // Only Field1
		TestStructOptional{Field1: "value2", Field2: &opt1}, // Field1 + Field2
		TestStructOptional{Field1: "value3", Field3: &opt2}, // Field1 + Field3
	}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	if len(secs) != 3 {
		t.Fatalf("expected 3 sections (different field compositions), got %d", len(secs))
	}

	// Each section should have exactly one row
	for key, sec := range secs {
		if len(sec) != 1 {
			t.Fatalf("section %s: expected 1 row, got %d", key, len(sec))
		}
	}
}

func TestSplit_IncludeOptionalFields(t *testing.T) {
	rows := []interface{}{
		TestStructOptional{Field1: "value1"},
		TestStructOptional{Field1: "value2"},
	}

	// With include=true, nil pointers should be included as empty values
	secs, err := split(rows, true)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	if len(secs) != 1 {
		t.Fatalf("expected 1 section when including optional fields, got %d", len(secs))
	}
}

func TestSplit_EmptyRows(t *testing.T) {
	rows := []interface{}{}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	if len(secs) != 0 {
		t.Fatalf("expected 0 sections for empty input, got %d", len(secs))
	}
}

func TestSplit_MissingRequiredField(t *testing.T) {
	rows := []interface{}{
		TestStruct{}, // Field1 is required (non-pointer, no default)
	}

	_, err := split(rows, false)
	if err == nil {
		t.Fatal("expected error for missing required field, got nil")
	}
}

func TestSplit_LargeDataset(t *testing.T) {
	// Create a large dataset for performance baseline
	const numRows = 1000
	rows := make([]interface{}, numRows)
	for i := 0; i < numRows; i++ {
		rows[i] = TestStruct{Field1: "value"}
	}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	// All rows have same fields, should be one section
	if len(secs) != 1 {
		t.Fatalf("expected 1 section, got %d", len(secs))
	}

	for _, sec := range secs {
		if len(sec) != numRows {
			t.Fatalf("expected %d rows in section, got %d", numRows, len(sec))
		}
	}
}

func TestSectionCreate(t *testing.T) {
	rows := []interface{}{
		TestStruct{Field1: "value1"},
		TestStruct{Field1: "value2"},
	}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	view := &View{
		Name:     "TEST",
		Required: true,
	}

	for _, sec := range secs {
		data := sec.create(view)
		output := string(data)

		// Verify CREATE VIEW statement
		if !strings.Contains(output, "CREATE VIEW TEST_CHG AS SELECT") {
			t.Error("expected CREATE VIEW statement in output")
		}

		// Verify INSERT INTO statement
		if !strings.Contains(output, "INSERT INTO TEST_CHG VALUES") {
			t.Error("expected INSERT INTO statement in output")
		}

		// Verify FROM clause references the DCT
		if !strings.Contains(output, "FROM TEST_DCT") {
			t.Error("expected FROM TEST_DCT in output")
		}

		// Verify ending semicolon
		if !strings.HasSuffix(output, ";\r\n\r\n") {
			t.Error("expected output to end with semicolon and CRLF")
		}
	}
}

func TestSectionCreate_WithAction(t *testing.T) {
	rows := []interface{}{
		TestStruct{Field1: "value1"},
	}

	secs, err := split(rows, false)
	if err != nil {
		t.Fatalf("split returned error: %v", err)
	}

	view := &View{
		Name:   "TEST",
		Action: "LOAD",
	}

	for _, sec := range secs {
		data := sec.create(view)
		output := string(data)

		// Verify action is used
		if !strings.Contains(output, "TEST_LOAD") {
			t.Error("expected TEST_LOAD in output when action is LOAD")
		}
	}
}

func TestRowArray(t *testing.T) {
	r := row{
		elems: []elem{
			{name: ptr("F01"), data: ptr("'value1'")},
			{name: ptr("F02"), data: ptr("10")},
			{name: ptr("F03"), data: ptr("'DEFAULT'")},
		},
	}

	names, values := r.array()

	if len(names) != 3 {
		t.Fatalf("expected 3 names, got %d", len(names))
	}
	if len(values) != 3 {
		t.Fatalf("expected 3 values, got %d", len(values))
	}

	expectedNames := []string{"F01", "F02", "F03"}
	expectedValues := []string{"'value1'", "10", "'DEFAULT'"}

	for i := range names {
		if names[i] != expectedNames[i] {
			t.Errorf("name[%d]: expected %s, got %s", i, expectedNames[i], names[i])
		}
		if values[i] != expectedValues[i] {
			t.Errorf("value[%d]: expected %s, got %s", i, expectedValues[i], values[i])
		}
	}
}

func TestRowArray_Empty(t *testing.T) {
	r := row{
		elems: []elem{},
	}

	names, values := r.array()

	if len(names) != 0 {
		t.Fatalf("expected 0 names, got %d", len(names))
	}
	if len(values) != 0 {
		t.Fatalf("expected 0 values, got %d", len(values))
	}
}

// Helper function to create string pointer
func ptr(s string) *string {
	return &s
}
