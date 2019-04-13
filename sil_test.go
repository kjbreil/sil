package sil

import (
	"testing"

	"github.com/kjbreil/sil/loc"
)

// OBJ will probably not work for an actual SIL file, this is for testing
// creation from a type
// type OBJ struct {
// 	F01 string `sil:"CHAR(13)"`
// 	F16 int    `sil:"INTEGER" default:"666"`
// 	F17 *int   `sil:"INTEGER"`
// }

func TestMake(t *testing.T) {
	m := make(Multi)

	// m["OBJ"] = Make("OBJ", loc.OBJ{})
	m.Make("OBJ", loc.OBJ{})

	m.AppendView("OBJ", loc.OBJ{
		F01: "9902",
	})

	m.AppendView("OBJ", loc.OBJ{
		F01: "8888",
	})
	// n := 1
	// m.AppendView("OBJ", loc.OBJ{
	// 	F01: "0000000009087",
	// 	F17: &n,
	// })
	m.AppendView("OBJ", loc.OBJ{
		F01: "9999",
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
