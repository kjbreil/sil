package loc

// RecTtl is the REC_TTL definition
type RecTtl struct {
	TransactionNumber   *int    `sil:"F1032"`
	TotalizerNumber     int     `sil:"F1034"`
	TotalizerDescriptor *string `sil:"F1039"`
	RegistrationMode    *string `sil:"F1067"`
	TotalTaxable1       *string `sil:"F1093"`
	TotalTaxable2       *string `sil:"F1094"`
	TotalTaxable3       *string `sil:"F1095"`
	TotalTaxable4       *string `sil:"F1096"`
	TotalFoodstamplable *string `sil:"F1097"`
	TotalWICable        *string `sil:"F1098"`
	TotalUnits          *string `sil:"F64"`
	TotalDollars        *string `sil:"F65"`
	TotalWeight         *string `sil:"F67"`
}
