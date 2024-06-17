package silread

import (
	"testing"
)

func TestGetStats(t *testing.T) {
	s, err := GetStats("./examples/dss.sql")
	if err != nil {
		t.Fatal(err)
	}
	if s.DataLines != 4 {
		t.Fatalf("expected 4 data lines, got %d", s.DataLines)
	}
}
