package sil

import (
	"fmt"
	"strings"
)

// section is a view
type section []row

func split(rows []interface{}) map[int]section {
	var ssec section

	for i := range rows {
		var r row
		r.make(rows[i], false)
		ssec = append(ssec, r)
	}

	secs := make(map[int]section)

	for i := range ssec {
		secs[len(ssec[i].elems)] = append(secs[len(ssec[i].elems)], ssec[i])
	}

	return secs

}

func (sec section) create(table string) (data []byte) {
	na, sa := sec[0].array()

	names := strings.Join(na, ",")
	// #nosec
	d := []byte(fmt.Sprintf("CREATE VIEW %s_CHG AS SELECT %s FROM %s_DCT%s", table, names, table, crlf))

	// #nosec
	d = append(d, []byte(fmt.Sprintf("INSERT INTO %s_CHG VALUES%s", table, crlf))...)

	for _, r := range sec {
		r.array()
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
