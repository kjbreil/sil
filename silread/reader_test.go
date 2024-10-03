package silread

import (
	"github.com/kjbreil/sil/tables"
	"os"
	"testing"
)

func TestUnmarshalReader(t *testing.T) {
	var dss []tables.PRICE

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
	var count int
	for range priceChan {
		count++
		// time.Sleep(time.Millisecond * 1)
	}
}

func TestNewReader(t *testing.T) {
	r, err := NewReader("./examples/price_load.sql")
	if err != nil {
		t.Fatal(err)
	}

	if r.DataLines != 4 {
		t.Fatalf("expected 4 data lines, got %d", r.DataLines)
	}
}
