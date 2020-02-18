package silread

import (
	"io/ioutil"
	"testing"

	"github.com/kjbreil/sil/tables"
)

func TestUnmarshal(t *testing.T) {
	var obj tables.OBJ

	b, _ := ioutil.ReadFile("./examples/single.sil")

	s, err := Unmarshal(b, &obj)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Marshal(true)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUnmarshalHeaders(t *testing.T) {
	var cll tables.CLL

	b, _ := ioutil.ReadFile("./examples/with_header.sil")

	s, err := Unmarshal(b, &cll)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Marshal(true)
	if err != nil {
		t.Fatal(err)
	}
}
