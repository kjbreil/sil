package loc

// InvHdr is the INV_HDR definition
type InvHdr struct {
	SubDepartmentCode      *int    `sil:"F04"`
	TransactionNumber      *int    `sil:"F1032"`
	MovementStartTime      *string `sil:"F1035"`
	MovementEndTime        *string `sil:"F1036"`
	TerminalStore          string  `sil:"F1056"`
	TerminalNumber         *string `sil:"F1057"`
	RegistrationMode       *string `sil:"F1067"`
	TransactionMode        *string `sil:"F1068"`
	LineNumber             *int    `sil:"F1101"`
	UserNumber             *int    `sil:"F1126"`
	UserShortName          *string `sil:"F1127"`
	ReferenceNumber        *string `sil:"F1245"`
	TransactionNote        *string `sil:"F1254"`
	TransBatchFileName     *string `sil:"F1255"`
	DateCreation           string  `sil:"F1264" default:"NOW"`
	OperatorReceiptPrinted *int    `sil:"F1687"`
	OperatorInvoicePrinted *int    `sil:"F1688"`
	OperatorExportMade     *int    `sil:"F1689"`
	CategoryCode           *int    `sil:"F17"`
	LastChangeDate         string  `sil:"F253" default:"NOW"`
	DateEnd                *string `sil:"F254"`
	SPARE                  *string `sil:"F2641"`
	SPARE                  *int    `sil:"F2642"`
	SPARE                  *string `sil:"F2643"`
	SPARE                  *int    `sil:"F2644"`
	SPARE                  *string `sil:"F2645"`
	SPARE                  *string `sil:"F2646"`
	SPARE                  *string `sil:"F2647"`
	VendorId               *string `sil:"F27"`
	TransactionComment     *string `sil:"F2848"`
	BatchIdentifier        *string `sil:"F902"`
}
