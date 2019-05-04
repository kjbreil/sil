package loc

// RptVnd is the RPT_VND definition
type RptVnd struct {
	MovementPeriodType string  `sil:"F1031"`
	TotalizerNumber    int     `sil:"F1034"`
	TerminalStore      string  `sil:"F1056"`
	LastChangeDate     string  `sil:"F253" default:"NOW"`
	DateEnd            *string `sil:"F254"`
	VendorId           *string `sil:"F27"`
	TotalUnits         *string `sil:"F64"`
	TotalDollars       *string `sil:"F65"`
	TotalWeight        *string `sil:"F67"`
}
