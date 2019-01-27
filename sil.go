package sil

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// SIL is the structure of a SIL file
type SIL struct {
	TableHeader Header
	Table       Table
	ViewHeader  Header
	View        View
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

type Column struct {
	Name string
	Type string
}

// View holds the data
type View struct {
	Name string
	Data [][]string
}

// New returns a new SIL
func New() SIL {
	var s SIL
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

func (s *SIL) Create() []byte {
	var f [][]byte
	f = append(f, []byte("INSERT INTO HEADER_DCT VALUES"))
	f = append(f, s.TableHeader.bytes())
	f = append(f, []byte("CREATE TABLE CLK_DCT(F1185 INTEGER,F1001 INTEGER,F1126 INTEGER,F1571 CHAR(60),F27 CHAR(14),F170 INTEGER,F253 DATE(7),F902 CHAR(8),F940 INTEGER,F941 INTEGER,F1000 CHAR(5),F1056 CHAR(4),F1127 CHAR(30),F1141 CHAR(200),F1142 INTEGER,F1143 CHAR(30),F1144 CHAR(30),F1145 DATE(7),F1146 CHAR(10),F1148 CHAR(14),F1176 CHAR(3),F1264 DATE(7),F1552 CHAR(4),F1553 DATE(7),F1554 DATE(7),F1555 DATE(7),F1556 INTEGER,F1557 CHAR(40),F1558 CHAR(40),F1559 CHAR(40),F1560 CHAR(20),F1561 CHAR(20),F1562 CHAR(15),F1563 CHAR(20),F1564 CHAR(20),F1565 CHAR(20),F1566 NUMBER(10,0),F1567 INTEGER,F1568 INTEGER,F1569 INTEGER,F1570 NUMBER(10,0),F1585 CHAR(20),F1586 CHAR(4),F1587 CHAR(20),F1588 CHAR(4),F1589 CHAR(15),F1590 CHAR(200),F1964 CHAR(4),F2587 DATE(7),F2597 CHAR(20),F2692 CHAR(1),F2806 CHAR(40),F2827 CHAR(10),F2828 NUMBER(8,4),F2829 NUMBER(8,4),F2830 NUMBER(8,4),F2831 NUMBER(8,4),F2832 CHAR(20),F2833 INTEGER,F2844 CHAR(200));\n"))
	f = append(f, s.table())
	f = append(f, []byte("INSERT INTO HEADER_DCT VALUES"))
	f = append(f, s.ViewHeader.bytes())
	f = append(f, []byte("CREATE VIEW CLK_CHG AS SELECT F1185,F1001,F1126,F1571,F27,F170,F253,F902,F940,F941,F1000,F1056,F1127,F1141,F1142,F1143,F1144,F1145,F1146,F1148,F1176,F1264,F1552,F1553,F1554,F1555,F1556,F1557,F1558,F1559,F1560,F1561,F1562,F1563,F1564,F1565,F1566,F1567,F1568,F1569,F1570,F1585,F1586,F1587,F1588,F1589,F1590,F1964,F2587,F2597,F2692,F2806,F2827,F2828,F2829,F2830,F2831,F2832,F2833,F2844 FROM CLK_DCT;\n"))
	f = append(f, []byte("INSERT INTO CLK_CHG VALUES"))

	var fwn []byte
	for _, eba := range f {
		fwn = append(fwn, eba...)
		fwn = append(fwn, []byte("\n")...)
	}

	return fwn
}

func text(t string) string {
	return "'" + t + "'"
}

func (h *Header) bytes() []byte {
	var itms []string
	len := 30
	for i := 0; i < len; i++ {
		var txt string
		switch i {
		case 0:
			txt = text(h.F901)
		case 1:
			txt = text(h.F902)
		case 2:
			txt = text(h.F903)
		case 3:
			txt = text(h.F904)
		case 8:
			txt = h.F909
		case 11:
			txt = text(h.F912)
		case 12:
			txt = text(h.F913)
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

func (s *SIL) table() []byte {
	var itms []string
	for _, col := range s.Table.Columns {
		var txt string

		txt = col.Name + " " + col.Type

		itms = append(itms, txt)
	}
	o := strings.Join(itms, ",")

	return []byte("CREATE TABLE " + s.Table.Name + "_DCT(" + o + ");\n")
}
