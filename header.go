package sil

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
// against the sql data type. Pointers are not handled normally and shoulk be
// used for optional elements.
// note: a default of NOW inserts to JulianNow
// F912 can be ADD, ADDRPL, CHANGE and REMOVE
type Header struct {
	F901 string `sil:"CHAR(30)" default:"HM"`             // Batch type
	F902 string `sil:"CHAR(30)" default:"00000001"`       // Batch identifier
	F903 string `sil:"CHAR(30)" default:"MANUAL"`         // Batch creator
	F904 string `sil:"CHAR(30)" default:"PAL"`            // Batch destination
	F905 string `sil:"CHAR(30)"`                          // Batch audit file
	F906 string `sil:"CHAR(30)"`                          // Batch response file
	F907 string `sil:"INTEGER" default:"NOW"`             // Batch ending date
	F908 string `sil:"INTEGER" default:"0000"`            // Batch ending time
	F909 string `sil:"INTEGER" default:"NOW"`             // Batch active date
	F910 string `sil:"INTEGER" default:"0000"`            // Batch active time
	F911 string `sil:"CHAR(30)"`                          // Batch purge date
	F912 string `sil:"CHAR(30)" default:"ADDRPL"`         // Batch action type
	F913 string `sil:"CHAR(30)" default:"ADDRPL FROM GO"` // Batch description
	F914 string `sil:"CHAR(30)"`                          // Batch user 1 (state)
	F918 string `sil:"CHAR(30)"`                          // Batch maximum error count
	F919 string `sil:"CHAR(30)"`                          // Batch file version
	F920 string `sil:"CHAR(30)"`                          // Batch creator version
	F921 string `sil:"CHAR(30)"`                          // Batch primary key
	F922 string `sil:"CHAR(30)"`                          // Batch specific command
	F930 string `sil:"CHAR(30)"`                          // Shelf tag type
	F931 string `sil:"CHAR(30)"`                          // Batch execution priority
	F932 string `sil:"CHAR(30)"`                          // Batch long description
}

// check and fix a header with default tag
func (h *Header) check() error {
	fields, values := fieldValue(h)

	_, err := forFields(fields, values, true)
	if err != nil {
		return err
	}
	return nil
}
