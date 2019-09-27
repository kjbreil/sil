package decode

import (
	"io/ioutil"
	"testing"

	"github.com/locug/sil/loc"
)

func TestUnmarshal(t *testing.T) {

	var obj loc.ObjTab
	// b := []byte("INSERT INTO HEADER_DCT VALUES")
	// Unmarshal(b, &obj)
	//
	// fmt.Println("CODE:", obj.UPCCode)

	b, _ := ioutil.ReadFile("./751943000.sil")

	Unmarshal(b, &obj)

	t.Fail()
}
