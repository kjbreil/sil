package loc

// RptClkN is the RPT_CLK_N definition
type RptClkN struct {
	TotalizerNumber int     `sil:"F1034"`
	UserOrderNumber *int    `sil:"F1185"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	TotalUnits      *string `sil:"F64"`
	TotalDollars    *string `sil:"F65"`
	TotalWeight     *string `sil:"F67"`
}
