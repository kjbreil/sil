package sil

import (
	"testing"
)

// OBJ will probably not work for an actual SIL file, this is for testing
// creation from a type
type OBJ struct {
	UPC     string `sil:"F01"`
	DEFAULT int    `sil:"F16" default:"6"`
	POINTER *int   `sil:"F17"`
}

func TestMake(t *testing.T) {
	m := make(Multi)

	// m["OBJ"] = Make("OBJ", loc.OBJ{})
	m.Make("OBJ", OBJ{})

	m.AppendView("OBJ", OBJ{
		UPC: "9902",
	})

	m.AppendView("OBJ", OBJ{
		UPC: "8888",
	})
	n := 1
	m.AppendView("OBJ", OBJ{
		UPC:     "0000000009087",
		POINTER: &n,
	})
	m.AppendView("OBJ", OBJ{
		UPC: "9999",
	})

	// m["PRICE"] = Make("PRICE", loc.PRICE{})
	// m["PRICE"].View.Data = append(m["PRICE"].View.Data, loc.PRICE{
	// 	F01: "0000000009902",
	// })

	b, err := m.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatalf(string(b))
}
