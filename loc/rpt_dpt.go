package loc

// RptDpt is the RPT_DPT definition
type RptDpt struct {
	DepartmentCode     *int    `sil:"F03"`
	MovementPeriodType string  `sil:"F1031"`
	TotalizerNumber    int     `sil:"F1034"`
	TerminalStore      string  `sil:"F1056"`
	TerminalNumber     *string `sil:"F1057"`
	LastChangeDate     string  `sil:"F253" default:"NOW"`
	DateEnd            *string `sil:"F254"`
	TotalUnits         *string `sil:"F64"`
	TotalDollars       *string `sil:"F65"`
	TotalWeight        *string `sil:"F67"`
}