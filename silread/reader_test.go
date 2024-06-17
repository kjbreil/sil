package silread

import (
	"fmt"
	"github.com/kjbreil/sil/tables"
	"os"
	"testing"
)

func TestUnmarshalReader(t *testing.T) {
	var dss []*tables.PRICE

	f, err := os.Open("./examples/Price_Load.sql")
	if err != nil {
		t.Fatal(err)
	}
	err = UnmarshalReader(f, &dss)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUnmarshalReaderChan(t *testing.T) {
	priceChan := make(chan tables.PRICE, 100)

	f, err := os.Open("./examples/Price_Load.sql")
	if err != nil {
		t.Fatal(err)
	}
	err = UnmarshalReaderChan(f, priceChan)
	if err != nil {
		t.Fatal(err)
	}

	for d := range priceChan {
		fmt.Println(d)
	}

}
