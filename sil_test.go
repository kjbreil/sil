package sil

import "testing"

func TestMakeSIL(t *testing.T) {

	s := MakeCLK()

	u := User{
		Number: 40,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)
	u = User{
		Number: 41,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)
	u = User{
		Number: 42,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)
	u = User{
		Number: 43,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.AddUser(u)

	s.Write("./out.sil")

}
