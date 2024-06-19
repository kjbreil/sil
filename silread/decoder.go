package silread

import (
	"fmt"
	"strings"
)

type decoder struct {
	p            parsed
	err          []error
	fcodes       []string
	fieldMap     []int
	tableName    string
	view         bool // has reached the view data so reading data from now on
	headerInsert bool // has read header so next line is headerinfo
	done         bool // set when done with sil processing
	header       []string
	data         [][]string
}

func newDecoder(p *parser) *decoder {
	return &decoder{}
}

func (d *decoder) makeFieldMap(data any) {
	for _, ef := range d.fcodes {

		d.fieldMap = append(d.fieldMap, findFieldIndex(ef, data))
	}
}

// itendifyLine identifies and works on the line returning the i of the next line
func (d *decoder) identifyLine(s int) int {
	// done returns the same s that as passed, breaking the processing
	if d.done {
		return s
	}
	var err error

	// if headerInsert has been reached, read the header
	if d.headerInsert {
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

		d.header, s, err = readDataLine(d.p, s, 22)
		if err != nil {
			d.err = append(d.err, err)
		}
		d.headerInsert = false
		return s
	}

	// view has been reached, reading data
	if d.view {
		var lineData []string
		lineData, s, err = readDataLine(d.p, s, len(d.fcodes))
		if err != nil {
			d.err = append(d.err, err)
		}
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

	// don't want line so read till CRLF
	for d.p[s].tok != CRLF {
		s++
	}

	return s
}

type rowData []string

func (r rowData) want() string {
	var sb strings.Builder

	sb.WriteString("[]string{")

	for i := range r {
		sb.WriteString("\"")
		sb.WriteString(r[i])
		sb.WriteString("\"")
		if i != len(r)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")

	return sb.String()
}

func readDataLine(p parsed, s int, columns int) (rowData, int, error) {
	var lineData []string

	// read the first semicolin
	if p[s].tok != OPEN {
		return lineData, s, fmt.Errorf("data does not start with (")
	}
	s++
	// the number of columns should equal the number of fcodes

	for i := 0; i < columns; i++ {
		var data string
		data, s = readData(p, s)

		if p[s].tok != COMMA && i != columns-1 && data != "" {
			return lineData, s, fmt.Errorf("data does not end with,")
		} else if p[s].tok == COMMA && data != "" {
			s++
		}

		lineData = append(lineData, data)
	}

	if p[s].tok == CLOSE {
		s++
	} else {
		return lineData, s, fmt.Errorf("data does not end with )")
	}

	// end of data grabbing
	if p[s].tok == SEMICOLON {
		// if d.tableName != "" {
		// 	d.done = true
		// }
		// this might cause an error unless there is a crlf after the semicolin...
		s++
		s++

		return lineData, s, nil
	}

	// advance a comma if it exists (doesn't seem to be strictly needed)
	if p[s].tok == COMMA {
		s++
	}

	// endline
	if p[s].tok == CRLF {
		s++
	} else {
		return lineData, s, fmt.Errorf("no endline at end of data")
	}

	return lineData, s, nil
}

func readData(p parsed, s int) (string, int) {
	var single bool

	var data string

	var opens int

	for {
		if single {
			if p[s].tok == SINGLE {
				single = !single
				// check if the next is another single quote, add a single single quote instead of two (Because it was
				// escaped)
				if p[s+1].tok == SINGLE {
					s++
					single = !single
					data += "''"
				}
			} else {
				data += p[s].val
			}

		} else {
			switch p[s].tok {
			case SINGLE:
				single = !single
				// check if next is a single quote, and capture a double single quote (since we are not already in
				// single quotes)
				if p[s+1].tok == SINGLE {
					s++
					single = !single
					data += "''"
				}
			case OPEN:
				opens++
			case CLOSE:
				opens--
				if opens < 0 {
					return data, s
				}
			case COMMA:
				if opens == 0 {
					// if there is no data built up advance to the next on return
					if len(data) == 0 {
						return data, s + 1
					}
					return data, s
				}
				data += p[s].val
			default:
				data += p[s].val
			}
		}

		s++
	}
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
	action := d.p.getAction(s)

	if action == "DCT" {
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
		// get the hearder information
		if d.p.isInsert(s, "HEADER") {
			// header row found so skip to next CRLF + 1
			d.headerInsert = true
			return d.p.nextLine(s)
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
	table, action := parseTable(prsd[s+4].val)
	switch action {
	case "DCT", "CHG", "RSP", "LOAD":
		return table
	default:
		return "ERROR"
	}
}

func parseTable(text string) (name string, action string) {
	strgs := strings.Split(strings.ToUpper(text), "_")
	return strings.Join(strgs[0:len(strgs)-1], "_"), strgs[len(strgs)-1]

}

func (prsd parsed) getAction(s int) string {
	_, action := parseTable(prsd[s+4].val)
	return action
}
