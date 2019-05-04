package loc

// RptCltN is the RPT_CLT_N definition
type RptCltN struct {
	TotalizerNumber int     `sil:"F1034"`
	CustomerID      *string `sil:"F1148"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	TotalUnits      *string `sil:"F64"`
	TotalDollars    *string `sil:"F65"`
	TotalWeight     *string `sil:"F67"`
}
