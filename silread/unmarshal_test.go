package silread

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/kjbreil/sil/tables"
)

// func TestUnmarshal(t *testing.T) {
// 	var obj tables.OBJ
//
// 	b, _ := ioutil.ReadFile("./examples/single.sil")
//
// 	s, err := Unmarshal(b, &obj)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	_, err = s.Marshal(true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// }
//
// func TestUnmarshalHeaders(t *testing.T) {
// 	var cll tables.CLL
//
// 	b, _ := ioutil.ReadFile("./examples/with_header.sil")
//
// 	s, err := Unmarshal(b, &cll)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	_, err = s.Marshal(true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestUnmarshalLoad(t *testing.T) {
	var dss []*tables.DSS

	b, _ := ioutil.ReadFile("./examples/dss.sql")

	err := Unmarshal(b, &dss)
	if err != nil {
		t.Fatal(err)
	}
	if len(dss) != 4 {
		t.Fatalf("expected dss length 4, got %d", len(dss))
	}
}

func TestUnmarshalLocLoad(t *testing.T) {
	var dss []*tables.DSS

	b, _ := os.ReadFile("./examples/Loc_Load.sql")

	err := Unmarshal(b, &dss)
	if err != nil {
		t.Fatal(err)
	}
}
