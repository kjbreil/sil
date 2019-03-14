package sil

import (
	"testing"
)

func TestHeader(t *testing.T) {
	var h Header

	h.check()

	s, err := MakeRow(h)
	if err != nil {
		t.Fatalf("i had an errur: %v", err)
	}

	if s != "('HM','00000001','MANUAL','PAL',,,"+JulianNow()+",0000,"+JulianNow()+",0000,,'ADDRPL','ADDRPL FROM GO',,,,,,,,,)" {
		t.Fail()

	}

}
