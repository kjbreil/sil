package silread

import (
	"io/ioutil"
	"testing"

	"github.com/kjbreil/sil/loc"
)

func TestUnmarshal(t *testing.T) {

	var obj loc.ObjTab

	b, _ := ioutil.ReadFile("./examples/single.sil")

	Unmarshal(b, &obj)

	t.Fail()
}
