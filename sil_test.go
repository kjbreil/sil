package sil

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	var h Header
	h.F912 = "ADDRPL"
	// h.F913 = "ADDRPL CHANGED OPERATORS"

	h.F907 = JulianNow()
	h.F908 = "0000"
	h.F909 = JulianNow()
	h.F910 = "0000"

	err := h.Check()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	s := MakeRow(h)

	if s != "('HM','00000001','MANUAL','PAL',,,"+JulianNow()+",0000,"+JulianNow()+",0000,,'ADDRPL','ADDRPL FROM GO',,,,,,,,,)" {
		t.Fail()

	}

}
