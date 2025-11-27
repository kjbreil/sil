package sil

import (
	"strings"
	"testing"
)

// MultiTestTable is a simple struct for testing Multi functionality
type MultiTestTable struct {
	ID   string `sil:"F01"`
	Name string `sil:"F02" default:"DEFAULT_NAME"`
}

// MultiTestTable2 is another table for testing multiple table handling
type MultiTestTable2 struct {
	Code  string  `sil:"F01"`
	Value *string `sil:"F03"`
}

func TestMultiMake(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})

	if _, ok := m["TEST"]; !ok {
		t.Fatal("expected TEST table to exist in Multi")
	}

	if m["TEST"].View.Name != "TEST" {
		t.Errorf("expected view name TEST, got %s", m["TEST"].View.Name)
	}
}

func TestMultiMake_Multiple(t *testing.T) {
	m := make(Multi)
	m.Make("TABLE1", MultiTestTable{})
	m.Make("TABLE2", MultiTestTable2{})

	if len(m) != 2 {
		t.Fatalf("expected 2 tables, got %d", len(m))
	}

	if _, ok := m["TABLE1"]; !ok {
		t.Error("expected TABLE1 to exist")
	}
	if _, ok := m["TABLE2"]; !ok {
		t.Error("expected TABLE2 to exist")
	}
}

func TestMultiMake_Overwrite(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})
	m.AppendView("TEST", MultiTestTable{ID: "first"})

	// Overwrite with new Make
	m.Make("TEST", MultiTestTable{})

	// View data should be empty after overwrite
	if len(m["TEST"].View.Data) != 0 {
		t.Errorf("expected empty view data after overwrite, got %d items", len(m["TEST"].View.Data))
	}
}

func TestMultiAppendView(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})
	m.AppendView("TEST", MultiTestTable{ID: "001"})
	m.AppendView("TEST", MultiTestTable{ID: "002"})
	m.AppendView("TEST", MultiTestTable{ID: "003"})

	if len(m["TEST"].View.Data) != 3 {
		t.Fatalf("expected 3 items in view data, got %d", len(m["TEST"].View.Data))
	}
}

func TestMultiAppendView_DifferentTables(t *testing.T) {
	m := make(Multi)
	m.Make("TABLE1", MultiTestTable{})
	m.Make("TABLE2", MultiTestTable2{})

	m.AppendView("TABLE1", MultiTestTable{ID: "001"})
	m.AppendView("TABLE2", MultiTestTable2{Code: "A001"})

	if len(m["TABLE1"].View.Data) != 1 {
		t.Errorf("TABLE1 expected 1 item, got %d", len(m["TABLE1"].View.Data))
	}
	if len(m["TABLE2"].View.Data) != 1 {
		t.Errorf("TABLE2 expected 1 item, got %d", len(m["TABLE2"].View.Data))
	}
}

func TestMultiSetHeaders(t *testing.T) {
	m := make(Multi)
	m.Make("TABLE1", MultiTestTable{})
	m.Make("TABLE2", MultiTestTable2{})

	m.SetHeaders("BATCH_DESCRIPTION")

	for name, s := range m {
		if s.Header.Description != "BATCH_DESCRIPTION" {
			t.Errorf("table %s: expected header description BATCH_DESCRIPTION, got %s", name, s.Header.Description)
		}
	}
}

func TestMultiMarshal_Empty(t *testing.T) {
	m := make(Multi)

	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	if len(data) != 0 {
		t.Errorf("expected empty data for empty Multi, got %d bytes", len(data))
	}
}

func TestMultiMarshal_SingleTable(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})
	m.AppendView("TEST", MultiTestTable{ID: "001"})

	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	output := string(data)

	// Should contain CREATE VIEW statement
	if !strings.Contains(output, "CREATE VIEW") {
		t.Error("expected CREATE VIEW in output")
	}

	// Should contain INSERT INTO statement
	if !strings.Contains(output, "INSERT INTO") {
		t.Error("expected INSERT INTO in output")
	}

	// Should contain the table name
	if !strings.Contains(output, "TEST") {
		t.Error("expected TEST in output")
	}
}

func TestMultiMarshal_MultipleTables(t *testing.T) {
	m := make(Multi)
	m.Make("TABLE1", MultiTestTable{})
	m.Make("TABLE2", MultiTestTable{})

	m.AppendView("TABLE1", MultiTestTable{ID: "001"})
	m.AppendView("TABLE2", MultiTestTable{ID: "002"})

	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	output := string(data)

	// Should contain both tables
	if !strings.Contains(output, "TABLE1") {
		t.Error("expected TABLE1 in output")
	}
	if !strings.Contains(output, "TABLE2") {
		t.Error("expected TABLE2 in output")
	}

	// Should have multiple CREATE VIEW statements
	count := strings.Count(output, "CREATE VIEW")
	if count != 2 {
		t.Errorf("expected 2 CREATE VIEW statements, got %d", count)
	}
}

func TestMultiMarshal_LargeDataset(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})

	// Add many rows
	const numRows = 100
	for i := 0; i < numRows; i++ {
		m.AppendView("TEST", MultiTestTable{ID: "row"})
	}

	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

func TestMultiMarshal_WithOptionalFields(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable2{})

	val := "optional_value"
	m.AppendView("TEST", MultiTestTable2{Code: "001", Value: &val})
	m.AppendView("TEST", MultiTestTable2{Code: "002"}) // Value is nil

	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	output := string(data)

	// Should handle optional fields correctly
	if !strings.Contains(output, "optional_value") {
		t.Error("expected optional_value in output")
	}
}

func TestMultiMarshal_ConsistentBatchPrefix(t *testing.T) {
	m := make(Multi)
	m.Make("TABLE1", MultiTestTable{})
	m.Make("TABLE2", MultiTestTable{})

	m.AppendView("TABLE1", MultiTestTable{ID: "001"})
	m.AppendView("TABLE2", MultiTestTable{ID: "002"})

	// Marshal and verify all tables have same prefix assigned
	_, err := m.Marshal()
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	// All SILs in Multi should have same prefix after Marshal
	var prefix int
	first := true
	for _, s := range m {
		if first {
			prefix = s.prefix
			first = false
		} else if s.prefix != prefix {
			t.Errorf("expected all SILs to have same prefix %d, got %d", prefix, s.prefix)
		}
	}
}

func TestMultiMarshal_Error(t *testing.T) {
	m := make(Multi)
	m.Make("TEST", RequiredFieldStruct{})

	// Add row with missing required field
	m.AppendView("TEST", RequiredFieldStruct{})

	_, err := m.Marshal()
	if err == nil {
		t.Fatal("expected error for missing required field")
	}
}

// Benchmark tests for establishing baseline performance
func BenchmarkMultiMarshal_SmallDataset(b *testing.B) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})
	for i := 0; i < 10; i++ {
		m.AppendView("TEST", MultiTestTable{ID: "row"})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Marshal()
	}
}

func BenchmarkMultiMarshal_MediumDataset(b *testing.B) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})
	for i := 0; i < 100; i++ {
		m.AppendView("TEST", MultiTestTable{ID: "row"})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Marshal()
	}
}

func BenchmarkMultiMarshal_LargeDataset(b *testing.B) {
	m := make(Multi)
	m.Make("TEST", MultiTestTable{})
	for i := 0; i < 1000; i++ {
		m.AppendView("TEST", MultiTestTable{ID: "row"})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Marshal()
	}
}

func BenchmarkMultiMarshal_MultipleTables(b *testing.B) {
	m := make(Multi)
	for i := 0; i < 10; i++ {
		tableName := string(rune('A' + i))
		m.Make(tableName, MultiTestTable{})
		for j := 0; j < 100; j++ {
			m.AppendView(tableName, MultiTestTable{ID: "row"})
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Marshal()
	}
}

func BenchmarkSplit_LargeDataset(b *testing.B) {
	rows := make([]interface{}, 1000)
	for i := range rows {
		rows[i] = MultiTestTable{ID: "row"}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplit_MixedFieldCompositions(b *testing.B) {
	rows := make([]interface{}, 1000)
	for i := range rows {
		if i%2 == 0 {
			rows[i] = MultiTestTable2{Code: "code"}
		} else {
			val := "value"
			rows[i] = MultiTestTable2{Code: "code", Value: &val}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}
