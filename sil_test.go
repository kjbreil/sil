package sil

import "testing"

func TestMakeSIL(t *testing.T) {

	s := New()

	s.CreateDCT()
	s.AddRplDCT()
	s.TableCLK()
	u := User{
		Number: 40,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.addUser(u)
	u = User{
		Number: 41,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.addUser(u)
	u = User{
		Number: 42,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.addUser(u)
	u = User{
		Number: 43,
		First:  "",
		Last:   "",
		Short:  "Some Person",
		Level:  2,
	}
	s.View.addUser(u)

	s.Write("./out.sil")

}
