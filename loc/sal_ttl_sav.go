package loc

// SalTtlSav is the SAL_TTL_SAV definition
type SalTtlSav struct {
	TransactionNumber     *int    `sil:"F1032"`
	TotalizerNumber       int     `sil:"F1034"`
	TotalizerDescriptor   *string `sil:"F1039"`
	TerminalStore         string  `sil:"F1056"`
	TerminalNumber        *string `sil:"F1057"`
	RegistrationMode      *string `sil:"F1067"`
	TotalTaxable1         *string `sil:"F1093"`
	TotalTaxable2         *string `sil:"F1094"`
	TotalTaxable3         *string `sil:"F1095"`
	TotalTaxable4         *string `sil:"F1096"`
	TotalFoodstamplable   *string `sil:"F1097"`
	TotalWICable          *string `sil:"F1098"`
	TransactionArchive    *string `sil:"F1505"`
	TimestampTransArchive *string `sil:"F2890"`
	TotalUnits            *string `sil:"F64"`
	TotalDollars          *string `sil:"F65"`
	TotalWeight           *string `sil:"F67"`
}
