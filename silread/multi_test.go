package silread

import (
	"log"
	"testing"

	"github.com/kjbreil/sil/loc"
)

func TestMulti(t *testing.T) {
	filename := "./examples/single.sil"

	tables := make(map[string]interface{})

	var obj loc.ObjTab

	tables["OBJ"] = &obj

	m, err := Multi(filename, tables)
	if err != nil {
		t.Fatal(err)
	}

	b, err := m.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	log.Println(string(b))

	t.Fail()
}
