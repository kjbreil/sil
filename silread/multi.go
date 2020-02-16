package silread

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kjbreil/sil"
)

// Multi reads from sil files that contain multiple batches/tables points to a
// single SIL file and given a map of string to functions the string is the
// first part of the SIL view table. for example a OBJ_CHG (CHG is hard coded
// right now) would have a string OBJ_CHG
func Multi(filename string, tables map[string]interface{}) (*sil.Multi, error) {

	b, _ := ioutil.ReadFile("./examples/single.sil")

	m := make(sil.Multi)

	// var robj loc.ObjTab

	obj := tables["OBJ"]

	// open a reader using the bytes as the start
	// this can be improved to read directly from a file
	r := bytes.NewReader(b)

	// make a new parser with the reader
	p := newParser(r)
	// prsd is the parsed file token parts
	prsd := p.parse()

	d, _ := prsd.decode(0)
	if len(d.err) > 0 {
		err := fmt.Errorf("could not decode the parsed sil file: %v", d.err)
		return &m, err
	}

	var fieldMap []int

	for _, ef := range d.fcodes {
		fieldIndex := findFieldIndex(ef, obj)
		if fieldIndex == -1 {
			err := fmt.Errorf("field %s does not exist in type definition", ef)
			return &m, err
		}

		fieldMap = append(fieldMap, fieldIndex)
	}

	s, err := d.SIL(obj, fieldMap)
	if err != nil {
		log.Println(err)
	}

	m.Make("OBJ", obj)
	m["OBJ"] = &s

	// m.Make("TEST", robj)
	// m["TEST"] = &s

	// m.Make("TEST", &tobj)
	// m["TEST"] = &s

	return &m, nil
}
