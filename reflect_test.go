package sil

import (
	"fmt"
	"testing"

	"github.com/kjbreil/sil/loc"
)

func TestReflect(t *testing.T) {
	s := Make("CFG", loc.CFG{})

	var c loc.CFG
	c.F1000 = "001901"
	c.F1056 = "001"
	c.F2846 = "GROC_LANE"
	c.F253 = JulianNow()
	c.F940 = 999
	c.F941 = 999
	c.F1001 = 1
	c.F1264 = JulianNow()

	keys := []string{"KEY1", "KEY2"}
	for range keys {
		c.F2845 = "SMS"
		c.F2847 = "VALUE"

		s.View.Data = append(s.View.Data, c)

	}

	// for _, section := range ini.Sections() {
	// 	// declare the CFG table type
	// 	var c loc.CFG

	// 	// make the key to fill
	// 	var k Key

	// 	k.filename = filename
	// 	k.section = section.Name()
	// 	for _, ele := range section.Keys() {
	// 		k.key = ele.Name()
	// 		c.F2845 = k.String()
	// 		c.F2847 = ele.Value()

	// 		s.View.Data = append(s.View.Data, c)

	// 	}
	// }
	// should read error return
	fmt.Println(s)

	s.Write("out.sil")
	t.Fail()
}
