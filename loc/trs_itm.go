package loc

// TrsItm is the TRS_ITM definition
type TrsItm struct {
	UPCCode           string  `sil:"F01,zeropad"`
	TransactionNumber *int    `sil:"F1032"`
	TotalizerNumber   int     `sil:"F1034"`
	MovementEndTime   *string `sil:"F1036"`
	TerminalStore     string  `sil:"F1056"`
	TerminalNumber    *string `sil:"F1057"`
	ReferenceNumber   *string `sil:"F1079"`
	LineNumber        *int    `sil:"F1101"`
	OperatorValidated *int    `sil:"F1765"`
	LastChangeDate    string  `sil:"F253" default:"NOW"`
	DateEnd           *string `sil:"F254"`
	Price             *string `sil:"F30"`
	PriceQty          *string `sil:"F31"`
	TotalUnits        *string `sil:"F64"`
	TotalDollars      *string `sil:"F65"`
	TotalWeight       *string `sil:"F67"`
}
