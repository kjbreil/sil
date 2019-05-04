package loc

// RptItmN is the RPT_ITM_N definition
type RptItmN struct {
	UPCCode           string  `sil:"F01,zeropad"`
	TotalizerNumber   int     `sil:"F1034"`
	TerminalStore     string  `sil:"F1056"`
	SoftField1        *string `sil:"F1301"`
	SoftField2        *string `sil:"F1302"`
	LastChangeDate    string  `sil:"F253" default:"NOW"`
	MovementResetDate *string `sil:"F381"`
	TotalUnits        *string `sil:"F64"`
	TotalDollars      *string `sil:"F65"`
	TotalWeight       *string `sil:"F67"`
}
