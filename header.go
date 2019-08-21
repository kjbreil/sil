package sil

import (
	"fmt"
)

// Batch Type Name Constants
const (
	ADD    = "ADD"
	ADDRPL = "ADDRPL"
	CHANGE = "CHANGE"
	REMOVE = "REMOVE"
)

// Header tells the system what the SIL file is doing.
// Since the header insert is not needed the only sil tag that is used is
// INTEGER - the rest are dummy holders but should get correct and validate data
// against the sql data type. Pointers are not handled normally and should be
// used for optional elements. 0000
// note: a default of NOW inserts to JulianNow
// F912 can be ADD, ADDRPL, CHANGE and REMOVE
type Header struct {
	Type              string `sil:"F901" default:"HM"`             // Batch type
	Identifier        string `sil:"F902" default:"00000001"`       // Batch identifier
	Creator           string `sil:"F903" default:"MANUAL"`         // Batch creator
	Destination       string `sil:"F904" default:"PAL"`            // Batch destination
	AuditFile         string `sil:"F905"`                          // Batch audit file
	ResponseFile      string `sil:"F906"`                          // Batch response file
	EndingDate        int    `sil:"F907" default:"NOW"`            // Batch ending date
	EndingTime        int    `sil:"F908" default:"0000"`           // Batch ending time
	ActiveDate        int    `sil:"F909" default:"NOW"`            // Batch active date
	ActiveTime        int    `sil:"F910" default:"0000"`           // Batch active time
	PurgeDate         string `sil:"F911"`                          // Batch purge date
	ActionType        string `sil:"F912" default:"ADDRPL"`         // Batch action type
	Description       string `sil:"F913" default:"ADDRPL FROM GO"` // Batch description
	UserOneState      string `sil:"F914"`                          // Batch user 1 (state)
	MaximumErrorCount string `sil:"F918"`                          // Batch maximum error count
	FileVersion       string `sil:"F919"`                          // Batch file version
	CreatorVersion    string `sil:"F920"`                          // Batch creator version
	PrimaryKey        string `sil:"F921"`                          // Batch primary key
	SpecificCommand   string `sil:"F922"`                          // Batch specific command
	TagType           string `sil:"F930"`                          // Shelf tag type
	ExecutionPriority string `sil:"F931"`                          // Batch execution priority
	LongDescription   string `sil:"F932"`                          // Batch long description
}

// insert creates the insert line with a crlf for newline
func (h *Header) insert() []byte {
	// #nosec
	return []byte(fmt.Sprintf("INSERT INTO HEADER_DCT VALUES%s", crlf))
}

func (h *Header) row() (b []byte) {
	b = append(b, rowBytes(*h)...)
	// header row is a single row so insert endline
	b = append(b, endLine()...)
	return
}
