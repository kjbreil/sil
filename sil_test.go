package sil

import (
	"fmt"
	"testing"
)

// OBJ will probably not work for an actual SIL file, this is for testing
// creation from a type
type OBJ struct {
	F01 string `sil:"CHAR(13)"`
	F16 int    `sil:"INTEGER" default:"666"`
	F17 *int   `sil:"INTEGER"`
}

func TestMake(t *testing.T) {
	s := Make("OBJ", OBJ{})

	s.View.Data = append(s.View.Data, OBJ{
		F01: "0000000009902",
		F16: 17,
	})
	n := 1
	s.View.Data = append(s.View.Data, OBJ{
		F01: "0000000009087",
		F17: &n,
	})
	s.View.Data = append(s.View.Data, OBJ{
		F01: "0000000009902",
		F16: 17,
	})

	str, err := s.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Putting out")
	fmt.Println(string(str))

	// ss, err := s.Optional()
	// if err != nil {
	// 	t.Fatalf("failed to get optional: %v", err)
	// }

	// for _, es := range ss {
	// 	b, err := es.Bytes()
	// 	if err != nil {
	// 		t.Fatalf("failed to convert to bytes with: %v", err)
	// 	}
	// 	t.Log(string(b))
	// }

	t.Fail()
}
