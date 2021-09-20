package sil

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/kjbreil/sil/tables"
)

// OBJ will probably not work for an actual SIL file, this is for testing
// creation from a type
// type OBJ struct {
// 	UPC     string `sil:"F01"`
// 	DEFAULT int    `sil:"F16" default:"6"`
// 	POINTER *int   `sil:"F17"`
// }

func TestMake(t *testing.T) {
	m := make(Multi)

	// m["OBJ"] = Make("OBJ", loc.OBJ{})
	m.Make("OBJ", tables.OBJ{})

	m.AppendView("OBJ", tables.OBJ{
		UPCCode: "9902",
	})

	m.AppendView("OBJ", tables.OBJ{
		UPCCode: "8888",
	})
	m.AppendView("OBJ", tables.OBJ{
		UPCCode: "0000000009087",
	})
	m.AppendView("OBJ", tables.OBJ{
		UPCCode: "9999",
	})

	m.Make("PRICE", tables.PRICE{})
	ap := "31.50"
	m.AppendView("PRICE", tables.PRICE{
		UPCCode: "9087",
		Price:   &ap,
	})

	b, err := m.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	//
	err = ioutil.WriteFile(fmt.Sprintf("%d.sil", time.Now().Nanosecond()), b, 0666)
	// return the error details
	if err != nil {
		t.Fatalf("ioutil error: %v", err)
	}

	t.Fatalf(string(b))
}

func TestSingle(t *testing.T) {
	// First test is to make sure you get an error when missing a required non defaulted field
	var s SIL
	s.View.Name = "OBJ"
	s.View.Data = append(s.View.Data, tables.OBJ{
		RecordStatus: 1,
	})
	_, err := s.Marshal(true)
	if err == nil {
		t.Fatalf("sil marshaling did not error when missing a required field")
	}
}

func TestDoubleSingleQuote(t *testing.T) {
	// First test is to make sure you get an error when missing a required non defaulted field

	u := "960000000062"
	ld := "test's"
	var s SIL
	s.View.Name = "OBJ"
	s.View.Data = append(s.View.Data, tables.OBJ{
		UPCCode:         u,
		LongDescription: &ld,
	})
	_, err := s.Marshal(false)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Write("test.sil", false, false)
	if err != nil {
		t.Fatal(err)
	}
}
