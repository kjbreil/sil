package sil

import (
	"reflect"
	"strings"
	"testing"
)

// CacheTestStruct for testing cached reflection
type CacheTestStruct struct {
	ID      string  `sil:"F01"`
	Name    string  `sil:"F02" default:"DEFAULT"`
	Value   int     `sil:"F03" default:"42"`
	Opt     *string `sil:"F04"`
	ZeroPad string  `sil:"F05,zeropad"`
}

func TestTypeCacheGetMeta(t *testing.T) {
	typ := reflect.TypeOf(CacheTestStruct{})
	meta := globalTypeCache.getTypeMeta(typ)

	if meta.numFields != 5 {
		t.Fatalf("expected 5 fields, got %d", meta.numFields)
	}

	// Verify field metadata
	tests := []struct {
		index      int
		name       string
		hasZeroPad bool
		defaultTag string
		isPointer  bool
	}{
		{0, "F01", false, "", false},
		{1, "F02", false, "DEFAULT", false},
		{2, "F03", false, "42", false},
		{3, "F04", false, "", true},
		{4, "F05", true, "", false},
	}

	for _, tt := range tests {
		fm := &meta.fieldMetas[tt.index]
		if fm.name != tt.name {
			t.Errorf("field %d: expected name %s, got %s", tt.index, tt.name, fm.name)
		}
		if fm.hasZeroPad != tt.hasZeroPad {
			t.Errorf("field %d: expected hasZeroPad %v, got %v", tt.index, tt.hasZeroPad, fm.hasZeroPad)
		}
		if fm.defaultTag != tt.defaultTag {
			t.Errorf("field %d: expected defaultTag %s, got %s", tt.index, tt.defaultTag, fm.defaultTag)
		}
		if fm.isPointer != tt.isPointer {
			t.Errorf("field %d: expected isPointer %v, got %v", tt.index, tt.isPointer, fm.isPointer)
		}
	}
}

func TestTypeCacheConcurrentAccess(t *testing.T) {
	typ := reflect.TypeOf(CacheTestStruct{})

	// Access from multiple goroutines
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			meta := globalTypeCache.getTypeMeta(typ)
			if meta.numFields != 5 {
				t.Errorf("expected 5 fields, got %d", meta.numFields)
			}
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestRowWithCache_MakeWithCache(t *testing.T) {
	opt := "optional"
	ts := CacheTestStruct{
		ID:      "001",
		Name:    "test",
		Value:   100,
		Opt:     &opt,
		ZeroPad: "123",
	}

	var r rowWithCache
	err := r.makeWithCache(ts, false)
	if err != nil {
		t.Fatalf("makeWithCache returned error: %v", err)
	}

	if len(r.elems) != 5 {
		t.Fatalf("expected 5 elements, got %d", len(r.elems))
	}
}

func TestRowWithCache_DefaultValues(t *testing.T) {
	ts := CacheTestStruct{
		ID:      "001",
		ZeroPad: "123",
	}

	var r rowWithCache
	err := r.makeWithCache(ts, false)
	if err != nil {
		t.Fatalf("makeWithCache returned error: %v", err)
	}

	// Check defaults are applied
	for _, elem := range r.elems {
		if *elem.name == "F02" {
			if !strings.Contains(*elem.data, "DEFAULT") {
				t.Errorf("expected default value for F02, got %s", *elem.data)
			}
		}
		if *elem.name == "F03" {
			if *elem.data != "42" {
				t.Errorf("expected default 42 for F03, got %s", *elem.data)
			}
		}
	}
}

func TestRowWithCache_NilPointersExcluded(t *testing.T) {
	ts := CacheTestStruct{
		ID:      "001",
		ZeroPad: "123",
	}

	var r rowWithCache
	err := r.makeWithCache(ts, false)
	if err != nil {
		t.Fatalf("makeWithCache returned error: %v", err)
	}

	// Nil pointer (F04) should not be included
	for _, elem := range r.elems {
		if *elem.name == "F04" {
			t.Error("nil pointer F04 should not be included when include=false")
		}
	}
}

func TestRowWithCache_NilPointersIncluded(t *testing.T) {
	ts := CacheTestStruct{
		ID:      "001",
		ZeroPad: "123",
	}

	var r rowWithCache
	err := r.makeWithCache(ts, true)
	if err != nil {
		t.Fatalf("makeWithCache returned error: %v", err)
	}

	// Find F04 which should be included as empty
	found := false
	for _, elem := range r.elems {
		if *elem.name == "F04" {
			found = true
			if *elem.data != "" {
				t.Errorf("expected empty value for nil pointer, got %s", *elem.data)
			}
		}
	}
	if !found {
		t.Error("F04 should be included when include=true")
	}
}

func TestRowWithCache_ZeroPad(t *testing.T) {
	ts := CacheTestStruct{
		ID:      "001",
		ZeroPad: "123",
	}

	var r rowWithCache
	err := r.makeWithCache(ts, false)
	if err != nil {
		t.Fatalf("makeWithCache returned error: %v", err)
	}

	for _, elem := range r.elems {
		if *elem.name == "F05" {
			if *elem.data != "'0000000000123'" {
				t.Errorf("expected zero-padded value, got %s", *elem.data)
			}
		}
	}
}

func TestRowWithCache_MissingRequired(t *testing.T) {
	ts := CacheTestStruct{
		// Missing ID
		ZeroPad: "123",
	}

	var r rowWithCache
	err := r.makeWithCache(ts, false)
	if err == nil {
		t.Fatal("expected error for missing required field")
	}
}

func TestSplitWithCache(t *testing.T) {
	rows := []interface{}{
		CacheTestStruct{ID: "001", ZeroPad: "1"},
		CacheTestStruct{ID: "002", ZeroPad: "2"},
		CacheTestStruct{ID: "003", ZeroPad: "3"},
	}

	secs, err := splitWithCache(rows, false)
	if err != nil {
		t.Fatalf("splitWithCache returned error: %v", err)
	}

	if len(secs) != 1 {
		t.Fatalf("expected 1 section, got %d", len(secs))
	}

	for _, sec := range secs {
		if len(sec) != 3 {
			t.Fatalf("expected 3 rows, got %d", len(sec))
		}
	}
}

func TestEscapeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "'hello'"},
		{"test's", "'test''s'"},
		{"", "''"},
		{"a'b'c", "'a''b''c'"},
	}

	for _, tt := range tests {
		result := escapeString(tt.input)
		if result != tt.expected {
			t.Errorf("escapeString(%q): expected %s, got %s", tt.input, tt.expected, result)
		}
	}
}

func TestFormatInt(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "0"},
		{42, "42"},
		{-42, "-42"},
		{12345, "12345"},
		{-12345, "-12345"},
	}

	for _, tt := range tests {
		result := formatInt(tt.input)
		if result != tt.expected {
			t.Errorf("formatInt(%d): expected %s, got %s", tt.input, tt.expected, result)
		}
	}
}

func TestApplyZeroPad(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"'123'", "'0000000000123'", false},
		{"'1'", "'0000000000001'", false},
		{"'1234567890123'", "'1234567890123'", false},
		{"'12345678901234'", "", true}, // Too long
	}

	for _, tt := range tests {
		result, err := applyZeroPad(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("applyZeroPad(%s): expected error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("applyZeroPad(%s): unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("applyZeroPad(%s): expected %s, got %s", tt.input, tt.expected, result)
			}
		}
	}
}

// Benchmarks comparing cached vs non-cached reflection
func BenchmarkRowMake(b *testing.B) {
	ts := CacheTestStruct{
		ID:      "001",
		Name:    "test",
		Value:   100,
		ZeroPad: "123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var r row
		_ = r.make(ts, false)
	}
}

func BenchmarkRowWithCacheMake(b *testing.B) {
	ts := CacheTestStruct{
		ID:      "001",
		Name:    "test",
		Value:   100,
		ZeroPad: "123",
	}

	// Warm up cache
	var r rowWithCache
	_ = r.makeWithCache(ts, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var r rowWithCache
		_ = r.makeWithCache(ts, false)
	}
}

func BenchmarkSplitOriginal(b *testing.B) {
	rows := make([]interface{}, 100)
	for i := range rows {
		rows[i] = CacheTestStruct{
			ID:      "id",
			Name:    "name",
			Value:   i,
			ZeroPad: "123",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplitWithCache(b *testing.B) {
	rows := make([]interface{}, 100)
	for i := range rows {
		rows[i] = CacheTestStruct{
			ID:      "id",
			Name:    "name",
			Value:   i,
			ZeroPad: "123",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = splitWithCache(rows, false)
	}
}

func BenchmarkSplitOriginal_1000(b *testing.B) {
	rows := make([]interface{}, 1000)
	for i := range rows {
		rows[i] = CacheTestStruct{
			ID:      "id",
			Name:    "name",
			Value:   i,
			ZeroPad: "123",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = split(rows, false)
	}
}

func BenchmarkSplitWithCache_1000(b *testing.B) {
	rows := make([]interface{}, 1000)
	for i := range rows {
		rows[i] = CacheTestStruct{
			ID:      "id",
			Name:    "name",
			Value:   i,
			ZeroPad: "123",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = splitWithCache(rows, false)
	}
}
