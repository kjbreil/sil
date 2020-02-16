package silread

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/kjbreil/sil/tables"
)

func TestUnmarshal(t *testing.T) {
	var obj tables.OBJ

	b, _ := ioutil.ReadFile("./examples/single.sil")

	s, err := Unmarshal(b, &obj)
	if err != nil {
		log.Println(err)
	}

	err = s.Write("test.sil", false, false)
	if err != nil {
		log.Println(err)
	}

	t.Fail()
}

// func TestFolder(t *testing.T) {

// 	silFiles, err := ioutil.ReadDir("./examples")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, eachFile := range silFiles {
// 		b, _ := ioutil.ReadFile("./examples/single.sil")
// 		Unmarshal(b, &obj)
// 	}
// }
