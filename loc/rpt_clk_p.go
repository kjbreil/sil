package loc

// RptClkP is the RPT_CLK_P definition
type RptClkP struct {
	TotalizerNumber int     `sil:"F1034"`
	TerminalStore   string  `sil:"F1056"`
	TerminalNumber  *string `sil:"F1057"`
	UserOrderNumber *int    `sil:"F1185"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	DateEnd         *string `sil:"F254"`
	TotalUnits      *string `sil:"F64"`
	TotalDollars    *string `sil:"F65"`
	TotalWeight     *string `sil:"F67"`
}
