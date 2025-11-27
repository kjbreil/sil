package sil

import (
	"reflect"
	"sync"
)

// typeCache caches reflection metadata for struct types
type typeCache struct {
	mu    sync.RWMutex
	cache map[reflect.Type]*typeMeta
}

// typeMeta holds cached metadata for a single type
type typeMeta struct {
	numFields  int
	fieldMetas []fieldMeta
}

// fieldMeta holds cached metadata for a single field
type fieldMeta struct {
	name     string // SIL tag name (e.g., "F01")
	hasZeroPad bool
	defaultTag string
	isPointer  bool
	fieldIndex int
}

// globalTypeCache is the global cache for type metadata
var globalTypeCache = &typeCache{
	cache: make(map[reflect.Type]*typeMeta),
}

// getTypeMeta returns cached metadata for a type, computing it if necessary
func (tc *typeCache) getTypeMeta(t reflect.Type) *typeMeta {
	tc.mu.RLock()
	meta, ok := tc.cache[t]
	tc.mu.RUnlock()
	if ok {
		return meta
	}

	// Compute metadata
	meta = tc.computeTypeMeta(t)

	tc.mu.Lock()
	tc.cache[t] = meta
	tc.mu.Unlock()

	return meta
}

// computeTypeMeta computes metadata for a type
func (tc *typeCache) computeTypeMeta(t reflect.Type) *typeMeta {
	meta := &typeMeta{
		numFields:  t.NumField(),
		fieldMetas: make([]fieldMeta, t.NumField()),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fm := fieldMeta{
			fieldIndex: i,
		}

		// Get SIL tag
		silTagStr := field.Tag.Get("sil")
		if silTagStr != "" {
			parts := splitTag(silTagStr)
			fm.name = parts[0]
			for _, opt := range parts[1:] {
				if opt == "zeropad" {
					fm.hasZeroPad = true
				}
			}
		}

		// Get default tag
		fm.defaultTag = field.Tag.Get("default")

		// Check if pointer
		fm.isPointer = field.Type.Kind() == reflect.Ptr

		meta.fieldMetas[i] = fm
	}

	return meta
}

// splitTag splits a tag string by comma
func splitTag(tag string) []string {
	var parts []string
	start := 0
	for i := 0; i < len(tag); i++ {
		if tag[i] == ',' {
			parts = append(parts, tag[start:i])
			start = i + 1
		}
	}
	parts = append(parts, tag[start:])
	return parts
}

// rowWithCache is a cached version of row.make
type rowWithCache struct {
	elems []elem
}

// makeWithCache creates a row using cached type metadata
func (r *rowWithCache) makeWithCache(rowType interface{}, include bool) error {
	t := reflect.TypeOf(rowType)
	v := reflect.ValueOf(rowType)

	meta := globalTypeCache.getTypeMeta(t)

	for i := 0; i < meta.numFields; i++ {
		fm := &meta.fieldMetas[i]
		fieldValue := v.Field(i)

		// Get the actual value
		val, isNilPtr, err := getFieldValue(fieldValue, fm)
		if err != nil {
			return err
		}

		// Handle required vs optional fields
		if !fm.isPointer && val == "" && fm.defaultTag == "" {
			return newRequiredFieldError(fm.name, t.Name())
		}

		switch {
		case !fm.isPointer && val == "":
			// Required field without value - already handled above
			continue
		case include && fm.isPointer && isNilPtr:
			// Include nil pointers as empty values
			var empty string
			name := fm.name
			r.elems = append(r.elems, elem{
				name: &name,
				data: &empty,
			})
		case !include && fm.isPointer && isNilPtr:
			// Skip nil pointers when not including
			continue
		default:
			name := fm.name
			r.elems = append(r.elems, elem{
				name: &name,
				data: &val,
			})
		}
	}

	return nil
}

// getFieldValue extracts the string value from a reflect.Value
func getFieldValue(v reflect.Value, fm *fieldMeta) (string, bool, error) {
	isNil := false

	// Handle pointers
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "", true, nil
		}
		v = v.Elem()
		isNil = false
	}

	// Get value based on kind
	var val string
	switch v.Kind() {
	case reflect.String:
		s := v.String()
		if s == "" {
			if fm.defaultTag != "" {
				val = getDefaultValue(fm.defaultTag)
			}
		} else {
			val = escapeString(s)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := v.Int()
		if i == 0 && fm.defaultTag != "" {
			val = fm.defaultTag
		} else {
			val = formatInt(i)
		}
	default:
		val = ""
	}

	// Apply zeropad if needed
	if fm.hasZeroPad && val != "" {
		padded, err := applyZeroPad(val)
		if err != nil {
			return "", isNil, err
		}
		val = padded
	}

	return val, isNil, nil
}

// getDefaultValue processes special default values
func getDefaultValue(def string) string {
	switch def {
	case "NOW":
		return "'" + JulianNow() + "'"
	case "NOWTIME":
		return "'" + JulianTimeNow() + "'"
	default:
		return "'" + def + "'"
	}
}

// escapeString escapes single quotes and wraps in quotes
func escapeString(s string) string {
	// Escape single quotes by doubling them
	escaped := make([]byte, 0, len(s)+2)
	escaped = append(escaped, '\'')
	for i := 0; i < len(s); i++ {
		if s[i] == '\'' {
			escaped = append(escaped, '\'', '\'')
		} else {
			escaped = append(escaped, s[i])
		}
	}
	escaped = append(escaped, '\'')
	return string(escaped)
}

// formatInt formats an integer
func formatInt(i int64) string {
	// Simple int to string conversion
	if i == 0 {
		return "0"
	}

	neg := false
	if i < 0 {
		neg = true
		i = -i
	}

	var buf [20]byte
	pos := len(buf)
	for i > 0 {
		pos--
		buf[pos] = byte('0' + i%10)
		i /= 10
	}

	if neg {
		pos--
		buf[pos] = '-'
	}

	return string(buf[pos:])
}

// applyZeroPad applies zero padding to a value
func applyZeroPad(val string) (string, error) {
	// Remove quotes if present
	if len(val) >= 2 && val[0] == '\'' && val[len(val)-1] == '\'' {
		inner := val[1 : len(val)-1]
		if len(inner) > 13 {
			return "", newZeroPadError(inner)
		}
		// Pad to 13 characters
		padded := make([]byte, 15) // quotes + 13 chars
		padded[0] = '\''
		padded[14] = '\''
		start := 1 + (13 - len(inner))
		for i := 1; i < start; i++ {
			padded[i] = '0'
		}
		copy(padded[start:], inner)
		return string(padded), nil
	}
	return val, nil
}

// newRequiredFieldError creates an error for missing required field
func newRequiredFieldError(fieldName, typeName string) error {
	return &requiredFieldError{fieldName: fieldName, typeName: typeName}
}

type requiredFieldError struct {
	fieldName string
	typeName  string
}

func (e *requiredFieldError) Error() string {
	return "the element " + e.fieldName + " does not contain any data and is required for table " + e.typeName
}

// newZeroPadError creates an error for zeropad overflow
func newZeroPadError(value string) error {
	return &zeroPadError{value: value}
}

type zeroPadError struct {
	value string
}

func (e *zeroPadError) Error() string {
	return "padded field contains more than 13 characters " + e.value
}

// splitWithCache uses cached reflection for better performance
func splitWithCache(rows []interface{}, include bool) (map[string]section, error) {
	if len(rows) == 0 {
		return make(map[string]section), nil
	}

	var ssec section

	for i := range rows {
		var r rowWithCache
		err := r.makeWithCache(rows[i], include)
		if err != nil {
			return nil, err
		}
		// Convert rowWithCache to row
		standardRow := row{elems: r.elems}
		ssec = append(ssec, standardRow)
	}

	secs := make(map[string]section)
	for i := range ssec {
		var key string
		for x := range ssec[i].elems {
			key = key + *ssec[i].elems[x].name
		}
		secs[key] = append(secs[key], ssec[i])
	}

	return secs, nil
}
