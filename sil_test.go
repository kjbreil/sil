package sil

import (
	"testing"
)

func TestMakeSIL(t *testing.T) {

	s := Make("CLK", CLK{})

	u := User{
		Number: 40,
		First:  "S1",
		Last:   "S2",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)
	u = User{
		Number: 41,
		First:  "S1",
		Last:   "S2",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)
	u = User{
		Number: 42,
		First:  "S1",
		Last:   "S2",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)
	u = User{
		Number: 43,
		First:  "S1",
		Last:   "S2",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)

	s.Write("./out.sil")

}

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

	if "('HM','00000001','MANUAL','PAL',,,2019072,0000,2019072,0000,,'ADDRPL','ADDRPL CHANGED OPERATORS',,,,,,,,,)" != s {
		t.Fail()

	}

}
