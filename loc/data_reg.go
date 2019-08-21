package loc

// DataReg is the DATA_REG definition
type DataReg struct {
	MovementPeriodType string  `sil:"F1031"`
	TotalizerNumber    int     `sil:"F1034"`
	TerminalStore      string  `sil:"F1056"`
	TransactionMode    *string `sil:"F1068"`
	ReferenceNumber    *string `sil:"F1079"`
	AlternateCode      *string `sil:"F154"`
	LastChangeDate     string  `sil:"F253" default:"NOW"`
	DateEnd            *string `sil:"F254"`
	TotalUnits         *string `sil:"F64"`
	TotalDollars       *string `sil:"F65"`
	TotalWeight        *string `sil:"F67"`
}
