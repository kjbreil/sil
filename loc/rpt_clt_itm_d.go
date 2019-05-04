package loc

// RptCltItmD is the RPT_CLT_ITM_D definition
type RptCltItmD struct {
	UPCCode         string  `sil:"F01,zeropad"`
	TotalizerNumber int     `sil:"F1034"`
	TerminalStore   string  `sil:"F1056"`
	CustomerId      *string `sil:"F1148"`
	SoftField1      *string `sil:"F1301"`
	SoftField2      *string `sil:"F1302"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	DateEnd         *string `sil:"F254"`
	TotalUnits      *string `sil:"F64"`
	TotalDollars    *string `sil:"F65"`
	TotalWeight     *string `sil:"F67"`
}
