package sil

import (
	"testing"
)

// OBJ will probably not work for an actual SIL file, this is for testing
// creation from a type
type OBJ struct {
	F01 string `sil:"CHAR(13)"`
	F16 int    `sil:"INTEGER"`
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
		F16: 17,
		F17: &n,
	})
	s.View.Data = append(s.View.Data, OBJ{
		F01: "0000000009902",
		F16: 17,
	})

	ss, err := s.Optional()
	if err != nil {
		t.Fatalf("failed to get optional: %v", err)
	}

	for _, es := range ss {
		b, err := es.Bytes()
		if err != nil {
			t.Fatalf("failed to convert to bytes with: %v", err)
		}
		t.Log(string(b))
	}

	t.Fail()
}

// func TestHeader(t *testing.T) {
// 	// create a header
// 	var h Header
// 	// run the check of the header
// 	err := h.check()
// 	if err != nil {
// 		t.Fatalf("failed to check header: %v", err)
// 	}
// 	// make the header into a row
// 	s, err := MakeRow(h)
// 	if err != nil {
// 		t.Fatalf("failed to make a row: %v", err)
// 	}

// 	// set the reference header to what we should see from defaults
// 	reference := "('HM','00000001','MANUAL','PAL',,," + JulianNow() + ",0000," + JulianNow() + ",0000,,'ADDRPL','ADDRPL FROM GO',,,,,,,,,)"

// 	if s != reference {
// 		t.Fatalf("Created header row did not match expected header row\nWant: %s\nHave: %s", reference, s)

// 	}

// }
