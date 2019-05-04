package loc

// RptCltY is the RPT_CLT_Y definition
type RptCltY struct {
	TotalizerNumber int     `sil:"F1034"`
	TerminalStore   string  `sil:"F1056"`
	CustomerId      *string `sil:"F1148"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	DateEnd         *string `sil:"F254"`
	TotalUnits      *string `sil:"F64"`
	TotalDollars    *string `sil:"F65"`
	TotalWeight     *string `sil:"F67"`
}
