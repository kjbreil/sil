package loc

// RptSub is the RPT_SUB definition
type RptSub struct {
	SubDepartmentCode  *int    `sil:"F04"`
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
