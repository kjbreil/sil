package loc

// RptCltItmN is the RPT_CLT_ITM_N definition
type RptCltItmN struct {
	UPCCode         string  `sil:"F01,zeropad"`
	TotalizerNumber int     `sil:"F1034"`
	CustomerID      *string `sil:"F1148"`
	SoftField1      *string `sil:"F1301"`
	SoftField2      *string `sil:"F1302"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	TotalUnits      *string `sil:"F64"`
	TotalDollars    *string `sil:"F65"`
	TotalWeight     *string `sil:"F67"`
}
