package sil

import (
	"context"
	"fmt"
	"strings"
)

// section is a view
type section []row

// split processes rows sequentially. It uses context.Background() internally.
// For context-aware splitting, use splitWithContext from concurrent.go
func split(rows []interface{}, include bool) (map[string]section, error) {
	return splitWithContext(context.Background(), rows, include)
}

// create makes the sil structure for each section
func (sec section) create(view *View) (data []byte) {
	// get the name array of the first section, all sections "should" match
	na, _ := sec[0].array()
	// join the names together with ,
	names := strings.Join(na, ",")
	// #nosec - ignore sql injection possibility error
	d := []byte(fmt.Sprintf("CREATE VIEW %s AS SELECT %s FROM %s_DCT;%s%s", view.action(), names, view.Name, crlf, crlf))

	// #nosec  - ignore sql injection possibility error
	d = append(d, []byte(fmt.Sprintf("INSERT INTO %s VALUES%s", view.action(), crlf))...)

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
