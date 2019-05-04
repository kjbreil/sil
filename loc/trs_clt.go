package loc

// TrsClt is the TRS_CLT definition
type TrsClt struct {
	TransactionNumber *int    `sil:"F1032"`
	TotalizerNumber   int     `sil:"F1034"`
	MovementEndTime   *string `sil:"F1036"`
	TerminalStore     string  `sil:"F1056"`
	TerminalNumber    *string `sil:"F1057"`
	ReferenceNumber   *string `sil:"F1079"`
	MatchingNumber    *int    `sil:"F1100"`
	LineNumber        *int    `sil:"F1101"`
	CustomerId        *string `sil:"F1148"`
	OperatorValidated *int    `sil:"F1765"`
	LastChangeDate    string  `sil:"F253" default:"NOW"`
	DateEnd           *string `sil:"F254"`
	TotalUnits        *string `sil:"F64"`
	TotalDollars      *string `sil:"F65"`
	TotalWeight       *string `sil:"F67"`
}
