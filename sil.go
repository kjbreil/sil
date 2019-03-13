package sil

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// SIL is the structure of a SIL file
type SIL struct {
	Table          Table
	ViewHeader     Header
	View           View
	Footer         []string
	TableType      interface{}
	SILType        string
	SILDescription string
}

// Header tells the system what the SIL file is doing
type Header struct {
	F901 string `sil:"CHAR(30)"` // Batch type
	F902 string `sil:"CHAR(30)"` // Batch identifier
	F903 string `sil:"CHAR(30)"` // Batch creator
	F904 string `sil:"CHAR(30)"` // Batch destination
	F905 string `sil:"CHAR(30)"` // Batch audit file
	F906 string `sil:"CHAR(30)"` // Batch response file
	F907 string `sil:"INTEGER"`  // Batch ending date
	F908 string `sil:"INTEGER"`  // Batch ending time
	F909 string `sil:"INTEGER"`  // Batch active date
	F910 string `sil:"INTEGER"`  // Batch active time
	F911 string `sil:"CHAR(30)"` // Batch purge date
	F912 string `sil:"CHAR(30)"` // Batch action type
	F913 string `sil:"CHAR(30)"` // Batch description
	F914 string `sil:"CHAR(30)"` // Batch user 1 (state)
	F918 string `sil:"CHAR(30)"` // Batch maximum error count
	F919 string `sil:"CHAR(30)"` // Batch file version
	F920 string `sil:"CHAR(30)"` // Batch creator version
	F921 string `sil:"CHAR(30)"` // Batch primary key
	F922 string `sil:"CHAR(30)"` // Batch specific command
	F930 string `sil:"CHAR(30)"` // Shelf tag type
	F931 string `sil:"CHAR(30)"` // Batch execution priority
	F932 string `sil:"CHAR(30)"` // Batch long description
}

// Table contains the definition of the columns to be inserted
type Table struct {
	Name    string
	Columns []Column
}

// Column is each column in a SIL file containing both the name and the type contained
type Column struct {
	Name string
	Type string
}

// View holds the data
type View struct {
	Name string
	Data []interface{}
}

// New returns a new SIL
func New() SIL {
	var s SIL
	return s
}

const (
	crlf = "\r\n"
	f902 = "00000001"
	f903 = "MANUAL"
	// f901 = "HC"
	f904 = "PAL"
	f909 = "000000"
	f910 = "0000"
	// f912 = "LOAD"
	// f913 = "CREATE DCT"
)

// Make makes a sil file of the definiton (as struct) passed
func Make(name string, definition interface{}) (s SIL) {
	// AddRpl header information - needs to be dynamic so deletes are possible
	s.AddRplDCT()
	s.Table.Name = name
	s.MakeTable(definition)
	return s
}

// Write writes a SIL to a file
func (s *SIL) Write(filename string) {
	// create the bytes of the SIL file
	mydata := s.Create()

	err := ioutil.WriteFile(filename, mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}

// Create creates the SIL structure from the information in the SIL type
func (s *SIL) Create() []byte {
	var f [][]byte
	// Header Insert
	f = append(f, []byte("INSERT INTO HEADER_DCT VALUES"))
	// Values to insert into header
	f = append(f, s.ViewHeader.bytes())
	// Create View for Data
	f = append(f, s.viewHeader())
	f = append(f, []byte("INSERT INTO "+s.Table.Name+"_CHG VALUES"))
	f = append(f, s.view())

	if len(s.Footer) > 0 {
		f = append(f, []byte("\r\n\r\n"))
		for _, es := range s.Footer {
			f = append(f, []byte(es))
		}
		f = append(f, []byte("\r\n"))
	}

	var fwn []byte
	for _, eba := range f {
		fwn = append(fwn, eba...)
		fwn = append(fwn, []byte("\r\n")...)
	}

	return fwn
}

// Append adds a line to the bottom of the SIL file
func (s *SIL) Append(str string) {
	s.Footer = append(s.Footer, str)
	return
}

// bytes creates the bytes of the header row
func (h *Header) bytes() []byte {
	return []byte(MakeRow(h) + ";" + crlf)
}

func (s *SIL) bytes() []byte {
	var itms []string
	for i := range s.Table.Columns {
		var txt string
		switch i {
		case 0:
		default:
			txt = ""
		}
		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("(" + o + ");\r\n")
}

func (s *SIL) viewHeader() []byte {
	var itms []string
	for _, col := range s.Table.Columns {
		var txt string

		txt = col.Name

		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("CREATE VIEW " + s.Table.Name + "_CHG AS SELECT " + o + " FROM " + s.Table.Name + "_DCT;\r\n")
}

// convert int to string for batches - ints should not have single quotes
func itoa(i *int) string {
	if i != nil {
		return strconv.Itoa(*i)
	}
	return ""
}

// convert anything besides int for LOC - this means put single quotes around it
func text(t *string) string {
	if t != nil {
		return "'" + *t + "'"
	}
	return ""
}

func (s *SIL) view() []byte {
	var lns []string
	for _, clk := range s.View.Data {
		var itms []string
		values := reflect.ValueOf(clk)

		for i := range s.Table.Columns {
			value := values.Field(i)
			var txt string
			if s.Table.Columns[i].Type == "INTEGER" {
				txt = fmt.Sprintf("%v", value)
			} else {
				txt = fmt.Sprintf("'%v'", value)

			}

			itms = append(itms, txt)
		}
		lns = append(lns, "("+strings.Join(itms, ",")+")")

	}
	cmb := strings.Join(lns, ",\r\n")
	cmb = cmb + ";\r\n"

	return []byte(cmb)

}

func (v *View) addCLK(c CLK) {
	v.Data = append(v.Data, c)
	return
}

// AddUser adds a user to the CLK
func (v *View) AddUser(u User) {
	if u.Level == 0 {
		return
	}
	var l CLK
	l.F1185 = u.Number
	l.F1126 = u.Number
	l.F1127 = u.Short
	l.F1142 = u.Level
	l.F1143 = u.First
	l.F1144 = u.Last
	l.F1145 = u.Birthdate

	// constants
	l.F253 = JulianTime()
	l.F1001 = 1
	l.F902 = "MANUAL"
	l.F1000 = "PAL"
	// l.F1964 = "999"

	v.Data = append(v.Data, l)
}

// JulianDate takes a time.Time and turns it into a julien date - just formatting
func JulianDate(t time.Time) string {
	return fmt.Sprintf("%04d%03d", t.Year(), t.YearDay())
}

// JulianNow returns the julian date for right now
func JulianNow() string {
	return JulianDate(time.Now())
}

// JulianTime is the julien date with a time formatted after
func JulianTime() string {
	n := time.Now()
	return fmt.Sprintf("%v %02d:%02d:%02d", JulianNow(), n.Hour(), n.Minute(), n.Second())
}
