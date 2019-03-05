package sil

import "testing"

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
