package decode

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
	data      [][]string
}

func (prsd parsed) decode() *decoder {

	// make a new decoder, put the parsed into it
	var d decoder
	d.p = prsd

	var i int

	for {
		ni := d.identifyLine(i)
		// if the new i matches the old i break out since processing failed
		if ni == i {
			// if d.err != nil {
			// 	return nil
			// }
			break
		}
		if ni > len(d.p)-1 {
			break
		}
		i = ni
	}

	// fmt.Println("Parser got here")

	return &d
}

// itendifyLine identifys and works on the line returning the i of the next line
func (d *decoder) identifyLine(s int) int {
	// view has been reached, reading data
	if d.view {
		return d.readDataLine(s)
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

func (d *decoder) readDataLine(s int) int {
	var lineData []string

	// read the first semicolin
	if d.p[s].tok != 5 {
		d.err = append(d.err, fmt.Errorf("data does not start with ("))
	}
	s++
	// the number of columns should equal the number of fcodes
	columns := len(d.fcodes)

	for i := 0; i < columns; i++ {
		var data string
		data, s = d.readData(s)

		if d.p[s].tok != 4 && i != columns-1 {
			d.err = append(d.err, fmt.Errorf("data does not end with ,"))
		} else if d.p[s].tok == 4 {
			s++
		}
		lineData = append(lineData, data)
	}

	if d.p[s].tok == 6 {
		s++
	} else {
		d.err = append(d.err, fmt.Errorf("data does not end with )"))
	}

	// end of data grabbing
	if d.p[s].tok == 7 {
		d.view = false
		s++
	}

	// endline
	if d.p[s].tok == 9 {
		s++
	} else {
		d.err = append(d.err, fmt.Errorf("no endline at end of data"))
	}

	d.data = append(d.data, lineData)

	// fmt.Println("the data: ", lineData, d.p[s].tok)

	return s
}

func (d *decoder) readData(s int) (string, int) {
	var single bool

	var data string

	// if there is a single quote advance one and set single to be true
	if d.p[s].tok == 8 {
		single = true
		s++
		// d.err = append(d.err, fmt.Errorf("data does not start with '"))
	}

	// the data
	if d.p[s].tok != 3 {
		d.err = append(d.err, fmt.Errorf("data is of another token type"))
	} else {
		data = d.p[s].val
	}
	s++
	// temp for now expect a ' but need to conditionally look for and skip
	if d.p[s].tok == 8 {
		if single {
			s++
		} else {
			d.err = append(d.err, fmt.Errorf("data ends with ' but did not start with one"))
		}
	}

	return data, s
}

func (d *decoder) readInsertLine(s int) int {

	fmt.Println(d.p[s+1])

	// switch to make
	switch {
	case d.p[s+1].tok == SINGLE:
		return s
	}

	return s
}

func (d *decoder) checkCreate(s int) int {
	name := d.p.getTable(s)
	switch name {
	case "OBJ":
		d.tableName = "OBJ"
	default:
		d.err = append(d.err, fmt.Errorf("table type %s not reconized yet", name))
		return s
	}

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
		// we don't care about the HEADER_DCT information so skip those
		// Still validate them if they exist

		if d.p.isInsert(s, d.p.nextCRLF(s), "HEADER_DCT") {
			// header row found so skip to next CRLF + 1
			s = d.p.nextLine(s)
			// since there was a header row there should be a single insert row, not doing much validation on it since LOC
			// doesn't - just needs to be enclosed by () with a ; at the end
			e := d.p.nextCRLF(s)
			// TODO: Properly announce which token is wrong rather than current error
			if d.p[s].tok != OPEN || d.p[e-2].tok != CLOSE || d.p[e-1].tok != SEMICOLON {
				d.err = append(d.err, fmt.Errorf("row for HEADER invalid, got %s%s%s want \"();\"", d.p[s].val, d.p[e-2].val, d.p[e-1].val))
				// since there was an error return s
				return s
			}
			return e + 1
		}
	}

	if d.p.isInsert(s, d.p.nextCRLF(s), fmt.Sprintf("%s_CHG", d.tableName)) {
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
	case !strings.Contains(prsd[s+4].val, table):
		return false
	case prsd[s+4].val != table:
		return false
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

	}

	return "ERROR"
}

// getValues returns an array of string representing the values in a open/close
func (prsd parsed) getValues(s, e int) []string {

	return []string{}
}
