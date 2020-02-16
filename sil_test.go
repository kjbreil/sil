package sil

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/kjbreil/go-loc/loc"
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
	m.Make("OBJ", loc.ObjTab{})

	m.AppendView("OBJ", loc.ObjTab{
		UPCCode: "9902",
	})

	m.AppendView("OBJ", loc.ObjTab{
		UPCCode: "8888",
	})
	n := 1
	m.AppendView("OBJ", loc.ObjTab{
		UPCCode:      "0000000009087",
		CategoryCode: &n,
	})
	m.AppendView("OBJ", loc.ObjTab{
		UPCCode: "9999",
	})

	m.Make("PRICE", loc.PriceTab{})
	ap := "31.50"
	m.AppendView("PRICE", loc.PriceTab{
		UPCCode:     "9087",
		ActivePrice: &ap,
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
	s.View.Data = append(s.View.Data, loc.ObjTab{
		RecordStatus: 1,
	})
	_, err := s.Marshal(true)
	if err == nil {
		t.Fatalf("sil marshaling did not error when missing a required field")
	}
}
