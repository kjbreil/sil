package sil

import "testing"

func TestMakeSIL(t *testing.T) {

	s := New()

	s.CreateDCT()
	s.AddRplDCT()
	s.TableCLK()

	s.Write("./out.sil")

}
