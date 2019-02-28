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
	TableHeader Header
	Table       Table
	ViewHeader  Header
	View        View
	TableType   interface{}
}

// Header tells the system what the SIL file is doing - an audit of sorts
type Header struct {
	F902 string // Batch identifier
	F903 string // Batch creator
	F901 string // Batch type
	F904 string // Batch destination
	F905 string // Batch audit file
	F906 string // Batch response file
	F907 string // Batch ending date
	F908 string // Batch ending time
	F909 string // Batch active date
	F910 string // Batch active time
	F911 string // Batch purge date
	F912 string // Batch action type
	F913 string // Batch description
	F914 string // Batch user 1 (state)
	F918 string // Batch maximum error count
	F919 string // Batch file version
	F920 string // Batch creator version
	F921 string // Batch primary key
	F922 string // Batch specific command
	F930 string // Shelf tag type
	F931 string // Batch execution priority
	F932 string // Batch long description
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
	Data []CLK
}

// New returns a new SIL
func New() SIL {
	var s SIL
	return s
}

// MakeCLK makes a CLK formatted SIL file
func MakeCLK() (s SIL) {
	s.CreateDCT()
	s.AddRplDCT()
	// s.TableCLK()
	s.MakeTable("CLK", CLK{})
	return s
}

// Make makes a sil file of the definiton (as struct) passed
func Make(name string, definition interface{}) (s SIL) {
	s.CreateDCT()
	s.AddRplDCT()
	// s.TableCLK()
	s.MakeTable(name, definition)
	return s
}

// CreateDCT Creates and returns the DCT information
func (s *SIL) CreateDCT() {
	s.TableHeader.F902 = "00001345" // Batch identifier
	s.TableHeader.F903 = "001901"   // Batch creator
	s.TableHeader.F901 = "HC"       // Batch type
	s.TableHeader.F904 = "PAL"      // Batch destination
	s.TableHeader.F909 = "000000"
	s.TableHeader.F910 = "0000"
	s.TableHeader.F912 = "LOAD"
	s.TableHeader.F913 = "CREATE DCT"
}

// AddRplDCT Creates and returns the DCT information
func (s *SIL) AddRplDCT() {
	s.ViewHeader.F902 = "00001345" // Batch identifier
	s.ViewHeader.F903 = "001901"   // Batch creator
	s.ViewHeader.F901 = "HM"       // Batch type
	s.ViewHeader.F910 = "0000"
	s.ViewHeader.F904 = "PAL" // Batch destination
	s.ViewHeader.F909 = "000000"
	s.ViewHeader.F912 = "ADDRPL"
	s.ViewHeader.F913 = "ADDRPL CHANGED OPERATORS"
}

func (s *SIL) Write(filename string) {
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
	f = append(f, []byte("INSERT INTO HEADER_DCT VALUES"))
	f = append(f, s.TableHeader.bytes())
	f = append(f, s.tableHeader())
	f = append(f, []byte("INSERT INTO HEADER_DCT VALUES"))
	f = append(f, s.ViewHeader.bytes())
	f = append(f, s.viewHeader())
	f = append(f, []byte("INSERT INTO CLK_CHG VALUES"))
	f = append(f, s.view())
	f = append(f, []byte("\n\n@EXEC(SQL=PCC_ACTIVATE_CLK);\n\n@EXEC(SQL=DEPLOY_CHG);"))

	var fwn []byte
	for _, eba := range f {
		fwn = append(fwn, eba...)
		fwn = append(fwn, []byte("\n")...)
	}

	return fwn
}

func (h *Header) bytes() []byte {
	var itms []string
	len := 30
	for i := 0; i < len; i++ {
		var txt string
		switch i {
		case 0:
			txt = text(&h.F901)
		case 1:
			txt = text(&h.F902)
		case 2:
			txt = text(&h.F903)
		case 3:
			txt = text(&h.F904)
		case 6:
			txt = julianNow()
		case 7:
			txt = "0000"
		case 8:
			txt = julianNow()
		case 9:
			txt = "0000"
		case 11:
			txt = text(&h.F912)
		case 12:
			txt = text(&h.F913)
		default:
			txt = ""
		}
		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("(" + o + ");\n")
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

	return []byte("(" + o + ");\n")
}

func (s *SIL) tableHeader() []byte {
	var itms []string
	for _, col := range s.Table.Columns {
		var txt string

		txt = col.Name + " " + col.Type

		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("CREATE TABLE " + s.Table.Name + "_DCT(" + o + ");\n")
}

func (s *SIL) viewHeader() []byte {
	var itms []string
	for _, col := range s.Table.Columns {
		var txt string

		txt = col.Name

		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("CREATE VIEW " + s.Table.Name + "_CHG AS SELECT " + o + " FROM " + s.Table.Name + "_DCT;\n")
}

func itoa(i *int) string {
	if i != nil {
		return strconv.Itoa(*i)
	}
	return ""
}

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
	cmb := strings.Join(lns, ",\n")
	cmb = cmb + ";\n"

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
	l.F253 = julianTime()
	l.F1001 = 1
	l.F902 = "MANUAL"
	l.F1000 = "PAL"
	l.F1964 = "999"

	v.Data = append(v.Data, l)
}

// JulianDate takes a time.Time and turns it into a julien date - just formatting
func JulianDate(t time.Time) string {
	return fmt.Sprintf("%04d%03d", t.Year(), t.YearDay())
}

// julianNow returns the julian date for right now
func julianNow() string {
	return JulianDate(time.Now())
}

func julianTime() string {
	n := time.Now()
	return fmt.Sprintf("%v %02d:%02d:%02d", julianNow(), n.Hour(), n.Minute(), n.Second())
}
