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
	// make the sil multi object, make is used to create the map
	m := make(sil.Multi)

	// #nosec
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return &m, err
	}

	r := bytes.NewReader(b)

	// prsd is the parsed file token parts
	prsd := newParser(r).parse()

	f := 0

	for {
		var d *decoder

		d, f = prsd.decode(f)
		if len(d.err) > 0 {
			err := fmt.Errorf("could not decode the parsed sil file: %v", d.err)
			return &m, err
		}

		if d.tableName == "" {
			break
		}

		// get the object
		tbl := tables[d.tableName]

		var fieldMap []int

		for _, ef := range d.fcodes {
			fieldIndex := findFieldIndex(ef, tbl)
			if fieldIndex == -1 {
				err := fmt.Errorf("field %s does not exist in type definition", ef)
				return &m, err
			}

			fieldMap = append(fieldMap, fieldIndex)
		}

		s, err := d.SIL(tbl, fieldMap)
		if err != nil {
			log.Println(err)
		}

		_, ok := m[d.tableName]
		if !ok {
			m.Make(d.tableName, tbl)
		}

		for _, ed := range s.View.Data {
			m.AppendView(d.tableName, ed)
		}
	}

	return &m, nil
}
