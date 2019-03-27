package sil

import (
	"fmt"
	"testing"

	"github.com/kjbreil/sil/loc"
)

func TestReflect(t *testing.T) {
	s := Make("CFG", loc.CFG{})

	c := loc.CFG{
		F1000: "001901",
		F1056: "001",
		F2846: "GROC_LANE",
		F253:  JulianNow(),
		F940:  999,
		F941:  999,
		F1001: 1,
		F1264: JulianNow(),
	}
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
