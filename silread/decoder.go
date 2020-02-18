package silread

import (
	"fmt"
	"strings"
)

type decoder struct {
	p         parsed
	err       []error
	fcodes    []string
	tableName string
	view      bool // has reached the view data so reading data from now on
	done      bool // set when done with sil processing
	header    []string
	data      [][]string
}

func (prsd parsed) decode(s int) (*decoder, int) {
	// make a new decoder, put the parsed into it
	var d decoder
	d.p = prsd

	// set i to the value passed as s
	i := s

	for {
		ni := d.identifyLine(i)
		// if the new i matches the old i break out since processing failed
		if ni == i {
			break
		}

		if ni > len(d.p)-1 {
			break
		}

		i = ni
	}

	return &d, i
}

// itendifyLine identifys and works on the line returning the i of the next line
func (d *decoder) identifyLine(s int) int {
	// done returns the same s that as passed, breaking the processing
	if d.done {
		return s
	}

	// view has been reached, reading data
	if d.view {
		var lineData []string
		lineData, s = d.readDataLine(s, len(d.fcodes))
		d.data = append(d.data, lineData)

		return s
	}

	// switch over the tokens first for faster matching for line time
	switch d.p[s].tok {
	// End of Line so Advance one
	case CRLF:
		return s + 1
	case OPEN:
		d.readInsertLine(s)
	}

	// detect line type based on the value
	switch d.p[s].val {
	case "INSERT":
		return d.checkInsert(s)
		// CREATE tells the type of view for the SIL file
	case "CREATE":
		return d.checkCreate(s)
	}

	return s
}

func (d *decoder) readDataLine(s int, columns int) ([]string, int) {
	var lineData []string

	// read the first semicolin
	if d.p[s].tok != OPEN {
		d.err = append(d.err, fmt.Errorf("data does not start with ("))
	}
	s++
	// the number of columns should equal the number of fcodes

	for i := 0; i < columns; i++ {
		var data string
		data, s = d.readData(s)

		if d.p[s].tok != COMMA && i != columns-1 && data != "" {
			d.err = append(d.err, fmt.Errorf("data does not end with ,"))
		} else if d.p[s].tok == COMMA && data != "" {
			s++
		}

		lineData = append(lineData, data)
	}

	if d.p[s].tok == CLOSE {
		s++
	} else {
		d.err = append(d.err, fmt.Errorf("data does not end with )"))
	}

	// end of data grabbing
	if d.p[s].tok == SEMICOLON {
		// this set to false might not be needed...
		d.view = false
		if d.tableName != "" {
			d.done = true
		}
		// this might cause an error unless there is a crlf after the semicolin...
		s++
		s++

		return lineData, s
	}

	// endline
	if d.p[s].tok == CRLF {
		s++
	} else {
		d.err = append(d.err, fmt.Errorf("no endline at end of data"))
	}

	return lineData, s
}

func (d *decoder) readData(s int) (string, int) {
	var single bool

	var data string

	// if there is a single quote advance one and set single to be true
	if d.p[s].tok == SINGLE {
		single = true
		s++
	}

	// the data
	switch {
	case d.p[s].tok == COMMA:
		data = ""
		s++

		return data, s
	case d.p[s].tok == CLOSE:
		return "", s
	case d.p[s].tok == SINGLE && single:
		s++
		// unless the next token is a close then add another to s because there is another entry, if its the last it shouldn't double out
		if d.p[s+1].tok != CLOSE {
			s++
		}
		return "", s
	case d.p[s].tok != IDENT:
		d.err = append(d.err, fmt.Errorf("data is of another token type"))
		s++
	default:
		// if the next token is whitespace add it
		for {
			if d.p[s].tok == SINGLE || d.p[s].tok == COMMA || d.p[s].tok == CLOSE {
				break
			}

			data += d.p[s].val
			s++
		}
	}

	// check if its a single quote check if it should be closing here and error if we shouldn't be closing
	if d.p[s].tok == SINGLE {
		if single {
			s++
		} else {
			d.err = append(d.err, fmt.Errorf("data ends with ' but did not start with one"))
		}
	}

	return data, s
}

func (d *decoder) readInsertLine(s int) int {
	// if the token is a single quote
	if d.p[s+1].tok == SINGLE {
		return s
	}

	return s
}

func (d *decoder) checkCreate(s int) int {
	// just trying to skip the table dct line
	if d.p.getAction(s) == "DCT" {
		return d.p.nextLine(s)
	}

	d.tableName = d.p.getTable(s)

	// TODO: Check validity of CREATE statement

	fs := s + 10
	// loop until a whitespace record is found
forStart:
	for {
		d.fcodes = append(d.fcodes, d.p[fs].val)
		fs++
		switch d.p[fs].tok {
		case WS:
			break forStart
		case COMMA:
			fs++
		default:
			d.err = append(d.err, fmt.Errorf("f code parsing error"))
			break forStart
		}
	}

	if d.p[fs+4].tok != SEMICOLON {
		d.err = append(d.err, fmt.Errorf("no semicolin at end of CREATE"))
		return s
	}

	if d.p[fs+5].tok == CRLF {
		return fs + 6
	}

	return s
}

func (d *decoder) checkInsert(s int) int {
	name := d.p.getTable(s)
	if name == "HEADER" {
		// TODO: should validate
		// get the heard information
		if d.p.isInsert(s, "HEADER") {
			// header row found so skip to next CRLF + 1
			s = d.p.nextLine(s)
			// since there was a header row there should be a single insert row, not doing much validation on it since LOC
			// doesn't - just needs to be enclosed by () with a ; at the end
			e := d.p.nextCRLF(s)
			if d.p[s+2].val == "HC" {
				return d.p.nextCRLF(s)
			}
			// TODO: Properly announce which token is wrong rather than current error
			if d.p[s].tok != OPEN || d.p[e-2].tok != CLOSE || d.p[e-1].tok != SEMICOLON {
				d.err = append(d.err, fmt.Errorf("row for HEADER invalid, got %s%s%s want \"();\"", d.p[s].val, d.p[e-2].val, d.p[e-1].val))
				// since there was an error return s
				return s
			}

			d.header, s = d.readDataLine(s, 22)

			return s
		}
	}

	if d.p.isInsert(s, name) {
		// the insert has been read and validated, time to read the data
		d.view = true
		return d.p.nextLine(s)
	}

	d.err = append(d.err, fmt.Errorf("table type for INSERT does not match CREATE"))

	return s
}

// nextCRLF returns the the i of the next CRLF
func (prsd parsed) nextCRLF(s int) int {
	for i := s; i <= len(prsd); i++ {
		if prsd[i].tok == CRLF {
			return i
		}
	}

	return s
}

// nextLine returns the the i of the start of the next line
func (prsd parsed) nextLine(s int) int {
	return prsd.nextCRLF(s) + 1
}

// isInsert checks if a insert statement is valid, dct is the "table" to expect
func (prsd parsed) isInsert(s int, table string) bool {
	// generic switch, if something fails the statement is not valid
	switch {
	case prsd[s].val != "INSERT":
		return false
	case prsd[s+2].val != "INTO":
		return false
	case !strings.EqualFold(prsd.getTable(s), table):
		return false
	// case prsd[s+4].val != table:
	// 	return false
	case prsd[s+6].val != "VALUES":
		return false
	}

	return true
}

// isInsert checks if a insert statement is valid, dct is the "table" to expect
func (prsd parsed) getTable(s int) string {
	strgs := strings.SplitAfter(prsd[s+4].val, "_")
	switch strgs[1] {
	case "DCT":
		return strgs[0][0 : len(strgs[0])-1]
	case "CHG":
		return strgs[0][0 : len(strgs[0])-1]
	case "RSP":
		return strgs[0][0 : len(strgs[0])-1]
	}

	return "ERROR"
}

func (prsd parsed) getAction(s int) (string, error) {
	strgs := strings.SplitAfter(prsd[s+4].val, "_")
	if len(strgs) != 2 {
		return "", fmt.Errorf("table did not match table _ action naming: %s", prsd[s+4].val)
	}
	return strgs[1]
}
