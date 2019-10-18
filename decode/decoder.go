package decode

import (
	"fmt"
	"strings"
)

func (prsd parsed) decode() error {
	var i int

	switch {
	case !prsd.isInsert(i, prsd.nextCRLF(i), "HEADER_DCT"):
		return fmt.Errorf("header insert statement malformed")
	}

	fmt.Println(prsd.getTable(i, prsd.nextCRLF(i)))

	return nil
}

// string returns the string of the data between s and e
func (prsd parsed) nextCRLF(s int) int {

	for i := s; i <= len(prsd); i++ {
		if prsd[i].tok == CRLF {
			return i
		}
	}

	return s
}

// string returns the string of the data between s and e
func (prsd parsed) string(s, e int) string {

	var strgs []string
	for i := s; i <= e; i++ {
		strgs = append(strgs, prsd[i].val)
	}

	return strings.Join(strgs, "")
}

// isInsert checks if a insert statement is valid, dct is the "table" to expect
func (prsd parsed) isInsert(s, e int, table string) bool {
	// generic switch, if something fails the statement is not valid
	switch {
	case prsd[s].val != "INSERT":
		return false
	case prsd[s+2].val != "INTO":
		return false
	// case !strings.Contains(prsd[s+4].val, "_DCT"):
	// 	return false
	case prsd[s+4].val != table:
		return false
	case prsd[s+6].val != "VALUES":
		return false
	}
	return true
}

// isInsert checks if a insert statement is valid, dct is the "table" to expect
func (prsd parsed) getTable(s, e int) string {
	strgs := strings.SplitAfter(prsd[s+4].val, "_")
	if strgs[1] == "DCT" {
		return strgs[0][0 : len(strgs[0])-1]
	}
	return "ERROR"
}

// getValues returns an array of string representing the values in a open/close
func (prsd parsed) getValues(s, e int) []string {

	return []string{}
}
