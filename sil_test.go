package sil

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/locug/sil/loc"
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

	// m["PRICE"] = Make("PRICE", loc.PRICE{})
	// m["PRICE"].View.Data = append(m["PRICE"].View.Data, loc.PRICE{
	// 	F01: "0000000009902",
	// })

	b, err := m.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%d.sil", time.Now().Nanosecond()), b, 0666)
	// return the error details
	if err != nil {
		t.Fatalf("ioutil error: %v", err)
	}

	t.Fatalf(string(b))
}
