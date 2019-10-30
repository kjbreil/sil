package sil

import (
	"fmt"
	"strings"
)

// section is a view
type section []row

// spit needs to be reworked it currently will combine two parts as the same because its based on number of
// elements instead of what elements are contained.
func split(rows []interface{}, include bool) (map[string]section, error) {
	var ssec section

	// take every row and reflect it
	for i := range rows {
		var r row
		// TODO: Handle this error but split needs to return an error then
		err := r.make(rows[i], include)
		if err != nil {
			return nil, err
		}
		ssec = append(ssec, r)
	}

	secs := make(map[string]section)

	for i := range ssec {
		// make the name of the section for the map based on the fields
		var key string
		for x := range ssec[i].elems {
			key = key + *ssec[i].elems[x].name
		}

		secs[key] = append(secs[key], ssec[i])
	}

	return secs, nil

}

// create makes the sil structure for each section
func (sec section) create(table string) (data []byte) {
	// get the name array of the first section, all sections "should" match
	na, _ := sec[0].array()
	// join the names together with ,
	names := strings.Join(na, ",")
	// #nosec - ignore sql injection possibility error
	d := []byte(fmt.Sprintf("CREATE VIEW %s_CHG AS SELECT %s FROM %s_DCT;%s%s", table, names, table, crlf, crlf))

	// #nosec  - ignore sql injection possibility error
	d = append(d, []byte(fmt.Sprintf("INSERT INTO %s_CHG VALUES%s", table, crlf))...)

	// create each line of the batch
	for _, r := range sec {
		_, sa := r.array()
		d = append(d, []byte(fmt.Sprintf("(%s)%s", strings.Join(sa, ","), crlf))...)
	}
	// remove the last CRLF, 2 bytes
	d = d[:len(d)-2]
	// append the endline code (; + crlf)
	d = append(d, endLine()...)
	return d
}

func (r *row) array() (na []string, sa []string) {
	for _, e := range r.elems {
		na = append(na, *e.name)
		sa = append(sa, *e.data)
	}
	return
}
