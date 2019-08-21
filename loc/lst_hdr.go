package loc

// LstHdr is the LST_HDR definition
type LstHdr struct {
	RecordStatus     int     `sil:"F1001" default:"1"`
	RegistrationMode *string `sil:"F1067"`
	TransactionMode  *string `sil:"F1068"`
	CustomerID       *string `sil:"F1148"`
	DateCreation     string  `sil:"F1264" default:"NOW"`
	LastChangeDate   string  `sil:"F253" default:"NOW"`
	ListID           *string `sil:"F2891"`
	ListDescription  *string `sil:"F2892"`
	ModifiedByUser   *int    `sil:"F941"`
}
