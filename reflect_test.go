package sil

import (
	"reflect"
	"strings"
	"testing"
)

// ReflectTestStruct for testing reflection functionality (without zeropad for ease of testing)
type ReflectTestStruct struct {
	StringField  string  `sil:"F01"`
	IntField     int     `sil:"F02"`
	DefaultField string  `sil:"F03" default:"MYDEFAULT"`
	IntDefault   int     `sil:"F04" default:"42"`
	PtrString    *string `sil:"F05"`
	PtrInt       *int    `sil:"F06"`
}

// ReflectTestStructWithZeroPad includes zeropad field for specific tests
type ReflectTestStructWithZeroPad struct {
	StringField string `sil:"F01"`
	ZeroPad     string `sil:"F07,zeropad"`
}

type RequiredFieldStruct struct {
	Required string `sil:"F01"`
}

func TestRowMake_AllFieldsPopulated(t *testing.T) {
	ptrStr := "pointer string"
	ptrInt := 100
	ts := ReflectTestStruct{
		StringField:  "test",
		IntField:     123,
		DefaultField: "custom",
		IntDefault:   99,
		PtrString:    &ptrStr,
		PtrInt:       &ptrInt,
	}

	var r row
	err := r.make(ts, false)
	if err != nil {
		t.Fatalf("row.make returned error: %v", err)
	}

	if len(r.elems) != 6 {
		t.Fatalf("expected 6 elements, got %d", len(r.elems))
	}

	// Verify some key values
	found := false
	for _, elem := range r.elems {
		if *elem.name == "F01" && strings.Contains(*elem.data, "test") {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected F01 element with 'test' value")
	}
}

func TestRowMake_DefaultValues(t *testing.T) {
	ts := ReflectTestStruct{
		StringField: "test",
	}

	var r row
	err := r.make(ts, false)
	if err != nil {
		t.Fatalf("row.make returned error: %v", err)
	}

	// Check that default values are used
	for _, elem := range r.elems {
		if *elem.name == "F03" {
			if !strings.Contains(*elem.data, "MYDEFAULT") {
				t.Errorf("expected default value MYDEFAULT for F03, got %s", *elem.data)
			}
		}
		if *elem.name == "F04" {
			if *elem.data != "42" {
				t.Errorf("expected default value 42 for F04, got %s", *elem.data)
			}
		}
	}
}

func TestRowMake_NilPointersExcluded(t *testing.T) {
	ts := ReflectTestStruct{
		StringField: "test",
	}

	var r row
	err := r.make(ts, false)
	if err != nil {
		t.Fatalf("row.make returned error: %v", err)
	}

	// Nil pointers should not be included when include=false
	for _, elem := range r.elems {
		if *elem.name == "F05" || *elem.name == "F06" {
			t.Errorf("nil pointer field %s should not be included when include=false", *elem.name)
		}
	}
}

func TestRowMake_NilPointersIncluded(t *testing.T) {
	ts := ReflectTestStruct{
		StringField: "test",
	}

	var r row
	err := r.make(ts, true)
	if err != nil {
		t.Fatalf("row.make returned error: %v", err)
	}

	// With include=true, all fields including nil pointers should be present
	fieldNames := make(map[string]bool)
	for _, elem := range r.elems {
		fieldNames[*elem.name] = true
	}

	expectedFields := []string{"F01", "F02", "F03", "F04", "F05", "F06"}
	for _, f := range expectedFields {
		if !fieldNames[f] {
			t.Errorf("expected field %s to be included", f)
		}
	}
}

func TestRowMake_MissingRequiredField(t *testing.T) {
	// Empty required field should cause error
	ts := RequiredFieldStruct{}

	var r row
	err := r.make(ts, false)
	if err == nil {
		t.Fatal("expected error for missing required field")
	}

	if !strings.Contains(err.Error(), "does not contain any data") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRowBytes(t *testing.T) {
	ts := ReflectTestStruct{
		StringField: "test",
		IntField:    123,
	}

	b := rowBytes(ts)
	output := string(b)

	// Should be wrapped in parentheses
	if !strings.HasPrefix(output, "(") || !strings.HasSuffix(output, ")") {
		t.Error("expected output to be wrapped in parentheses")
	}

	// Should contain the values
	if !strings.Contains(output, "'test'") {
		t.Error("expected string value to be quoted")
	}
	if !strings.Contains(output, "123") {
		t.Error("expected int value")
	}
}

func TestValue_StringField(t *testing.T) {
	type TestStr struct {
		Str string `sil:"F01"`
	}

	ts := TestStr{Str: "hello"}
	v := reflect.ValueOf(ts)
	f := reflect.TypeOf(ts).Field(0)

	val, name, ptr, err := value(v.Field(0), f)
	if err != nil {
		t.Fatalf("value returned error: %v", err)
	}

	if *name != "F01" {
		t.Errorf("expected name F01, got %s", *name)
	}

	if *val != "'hello'" {
		t.Errorf("expected quoted value 'hello', got %s", *val)
	}

	if *ptr != false {
		t.Error("expected ptr to be false for non-pointer field")
	}
}

func TestValue_IntField(t *testing.T) {
	type TestInt struct {
		Num int `sil:"F02"`
	}

	ts := TestInt{Num: 42}
	v := reflect.ValueOf(ts)
	f := reflect.TypeOf(ts).Field(0)

	val, name, ptr, err := value(v.Field(0), f)
	if err != nil {
		t.Fatalf("value returned error: %v", err)
	}

	if *name != "F02" {
		t.Errorf("expected name F02, got %s", *name)
	}

	if *val != "42" {
		t.Errorf("expected value 42, got %s", *val)
	}

	if *ptr != false {
		t.Error("expected ptr to be false for non-pointer field")
	}
}

func TestValue_PointerField(t *testing.T) {
	type TestPtr struct {
		Ptr *string `sil:"F03"`
	}

	str := "pointer value"
	ts := TestPtr{Ptr: &str}
	v := reflect.ValueOf(ts)
	f := reflect.TypeOf(ts).Field(0)

	val, name, ptr, err := value(v.Field(0), f)
	if err != nil {
		t.Fatalf("value returned error: %v", err)
	}

	if *name != "F03" {
		t.Errorf("expected name F03, got %s", *name)
	}

	if !strings.Contains(*val, "pointer value") {
		t.Errorf("expected value containing 'pointer value', got %s", *val)
	}

	if *ptr != true {
		t.Error("expected ptr to be true for pointer field")
	}
}

func TestValue_DefaultTag(t *testing.T) {
	type TestDefault struct {
		Field string `sil:"F01" default:"DEFAULT_VALUE"`
	}

	ts := TestDefault{} // Empty, should use default
	v := reflect.ValueOf(ts)
	f := reflect.TypeOf(ts).Field(0)

	val, _, _, err := value(v.Field(0), f)
	if err != nil {
		t.Fatalf("value returned error: %v", err)
	}

	if !strings.Contains(*val, "DEFAULT_VALUE") {
		t.Errorf("expected default value, got %s", *val)
	}
}

func TestDefaultTag_NOW(t *testing.T) {
	type TestNow struct {
		Field string `sil:"F01" default:"NOW"`
	}

	f := reflect.TypeOf(TestNow{}).Field(0)
	result := defaultTag(&f)

	// NOW should return julian date format (YYYYDDD)
	if len(result) != 7 {
		t.Errorf("expected julian date length 7, got %d: %s", len(result), result)
	}
}

func TestDefaultTag_NOWTIME(t *testing.T) {
	type TestNowTime struct {
		Field string `sil:"F01" default:"NOWTIME"`
	}

	f := reflect.TypeOf(TestNowTime{}).Field(0)
	result := defaultTag(&f)

	// NOWTIME should return julian date with time (YYYYDDD HH:MM:SS)
	if !strings.Contains(result, ":") {
		t.Errorf("expected time format with colons, got %s", result)
	}
}

func TestDefaultTag_CustomValue(t *testing.T) {
	type TestCustom struct {
		Field string `sil:"F01" default:"CUSTOM"`
	}

	f := reflect.TypeOf(TestCustom{}).Field(0)
	result := defaultTag(&f)

	if result != "CUSTOM" {
		t.Errorf("expected CUSTOM, got %s", result)
	}
}

func TestKind_String(t *testing.T) {
	s := "test string"
	v := reflect.ValueOf(s)
	dt := ""

	val, ptr := kind(&v, &dt)

	if val != "'test string'" {
		t.Errorf("expected quoted string, got %s", val)
	}
	if ptr != false {
		t.Error("expected ptr to be false")
	}
}

func TestKind_Int(t *testing.T) {
	i := 42
	v := reflect.ValueOf(i)
	dt := ""

	val, ptr := kind(&v, &dt)

	if val != "42" {
		t.Errorf("expected 42, got %s", val)
	}
	if ptr != false {
		t.Error("expected ptr to be false")
	}
}

func TestKind_Pointer(t *testing.T) {
	s := "pointer"
	p := &s
	v := reflect.ValueOf(p)
	dt := ""

	val, ptr := kind(&v, &dt)

	if val != "'pointer'" {
		t.Errorf("expected quoted pointer value, got %s", val)
	}
	if ptr != true {
		t.Error("expected ptr to be true for pointer")
	}
}

func TestReflectString_Empty(t *testing.T) {
	s := ""
	v := reflect.ValueOf(s)
	dt := ""
	hd := false

	result := reflectString(&v, &dt, &hd)

	if result != "" {
		t.Errorf("expected empty string, got %s", result)
	}
}

func TestReflectString_EmptyWithDefault(t *testing.T) {
	s := ""
	v := reflect.ValueOf(s)
	dt := "DEFAULT"
	hd := true

	result := reflectString(&v, &dt, &hd)

	if result != "'DEFAULT'" {
		t.Errorf("expected 'DEFAULT', got %s", result)
	}
}

func TestReflectString_WithValue(t *testing.T) {
	s := "value"
	v := reflect.ValueOf(s)
	dt := ""
	hd := false

	result := reflectString(&v, &dt, &hd)

	if result != "'value'" {
		t.Errorf("expected 'value', got %s", result)
	}
}

func TestReflectString_EscapeSingleQuotes(t *testing.T) {
	s := "test's value"
	v := reflect.ValueOf(s)
	dt := ""
	hd := false

	result := reflectString(&v, &dt, &hd)

	// Single quotes should be doubled
	if result != "'test''s value'" {
		t.Errorf("expected single quotes to be escaped, got %s", result)
	}
}

func TestReflectInt_Zero(t *testing.T) {
	i := 0
	v := reflect.ValueOf(i)
	dt := ""
	hd := false

	result := reflectInt(&v, &dt, &hd)

	if result != "0" {
		t.Errorf("expected 0, got %s", result)
	}
}

func TestReflectInt_ZeroWithDefault(t *testing.T) {
	i := 0
	v := reflect.ValueOf(i)
	dt := "99"
	hd := true

	result := reflectInt(&v, &dt, &hd)

	if result != "99" {
		t.Errorf("expected default 99, got %s", result)
	}
}

func TestReflectInt_WithValue(t *testing.T) {
	i := 123
	v := reflect.ValueOf(i)
	dt := "99"
	hd := true

	result := reflectInt(&v, &dt, &hd)

	if result != "123" {
		t.Errorf("expected 123, got %s", result)
	}
}

func TestValue_ZeroPadField(t *testing.T) {
	type TestZeroPad struct {
		UPC string `sil:"F01,zeropad"`
	}

	ts := TestZeroPad{UPC: "123"}
	v := reflect.ValueOf(ts)
	f := reflect.TypeOf(ts).Field(0)

	val, _, _, err := value(v.Field(0), f)
	if err != nil {
		t.Fatalf("value returned error: %v", err)
	}

	// Zeropad should pad to 13 characters with leading zeros
	if *val != "'0000000000123'" {
		t.Errorf("expected zero-padded value, got %s", *val)
	}
}

func TestValue_ZeroPadTooLong(t *testing.T) {
	type TestZeroPad struct {
		UPC string `sil:"F01,zeropad"`
	}

	ts := TestZeroPad{UPC: "12345678901234"} // 14 chars, exceeds 13
	v := reflect.ValueOf(ts)
	f := reflect.TypeOf(ts).Field(0)

	_, _, _, err := value(v.Field(0), f)
	if err == nil {
		t.Fatal("expected error for zeropad field exceeding 13 characters")
	}
}
