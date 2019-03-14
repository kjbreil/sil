package sil

import (
	"testing"
)

func TestHeader(t *testing.T) {
	var h Header
	h.F902 = f902 // Batch identifier
	h.F903 = f903 // Batch creator
	h.F901 = "HM" // Batch type
	h.F910 = f910
	h.F904 = f904 // Batch destination
	h.F909 = f909
	h.F912 = "ADDRPL"
	h.F913 = "ADDRPL CHANGED OPERATORS"

	h.F907 = JulianNow()
	h.F908 = "0000"
	h.F909 = JulianNow()
	h.F910 = "0000"

	s := MakeRow(h)

	if s != "('HM','00000001','MANUAL','PAL',,,2019072,0000,2019072,0000,,'ADDRPL','ADDRPL CHANGED OPERATORS',,,,,,,,,)" {
		t.Fail()

	}

}
