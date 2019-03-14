package sil

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
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
	F901 string `sil:"CHAR(30)" default:"00000001"` // Batch type
	F902 string `sil:"CHAR(30)"`                    // Batch identifier
	F903 string `sil:"CHAR(30)"`                    // Batch creator
	F904 string `sil:"CHAR(30)"`                    // Batch destination
	F905 string `sil:"CHAR(30)"`                    // Batch audit file
	F906 string `sil:"CHAR(30)"`                    // Batch response file
	F907 string `sil:"INTEGER"`                     // Batch ending date
	F908 string `sil:"INTEGER"`                     // Batch ending time
	F909 string `sil:"INTEGER"`                     // Batch active date
	F910 string `sil:"INTEGER"`                     // Batch active time
	F911 string `sil:"CHAR(30)"`                    // Batch purge date
	F912 string `sil:"CHAR(30)"`                    // Batch action type
	F913 string `sil:"CHAR(30)"`                    // Batch description
	F914 string `sil:"CHAR(30)"`                    // Batch user 1 (state)
	F918 string `sil:"CHAR(30)"`                    // Batch maximum error count
	F919 string `sil:"CHAR(30)"`                    // Batch file version
	F920 string `sil:"CHAR(30)"`                    // Batch creator version
	F921 string `sil:"CHAR(30)"`                    // Batch primary key
	F922 string `sil:"CHAR(30)"`                    // Batch specific command
	F930 string `sil:"CHAR(30)"`                    // Shelf tag type
	F931 string `sil:"CHAR(30)"`                    // Batch execution priority
	F932 string `sil:"CHAR(30)"`                    // Batch long description
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

// Some Constants
const (
	crlf   = "\r\n"
	sqlInt = "INTEGER"
)

// Bad constants
const (
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
	// #nosec
	f = append(f, []byte(fmt.Sprintf("INSERT INTO %s_CHG VALUES", s.Table.Name)))
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
}

// bytes creates the bytes of the header row
func (h *Header) bytes() []byte {
	return []byte(MakeRow(h) + ";" + crlf)
}

func (s *SIL) viewHeader() []byte {
	var itms []string
	for _, col := range s.Table.Columns {
		txt := col.Name
		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("CREATE VIEW " + s.Table.Name + "_CHG AS SELECT " + o + " FROM " + s.Table.Name + "_DCT;\r\n")
}

func (s *SIL) view() []byte {
	var lns []string
	for _, clk := range s.View.Data {
		var itms []string
		values := reflect.ValueOf(clk)

		for i := range s.Table.Columns {
			value := values.Field(i)
			var txt string
			if s.Table.Columns[i].Type == sqlInt {
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
