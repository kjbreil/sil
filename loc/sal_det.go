package loc

// SalDet is the SAL_DET definition
type SalDet struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	TransactionNumber       *int    `sil:"F1032"`
	DescriptionRegistration *string `sil:"F1041"`
	ReferenceNumber         *string `sil:"F1079"`
	AlphaParameter          *string `sil:"F1081"`
	LineNumber              *int    `sil:"F1101"`
	RegistrationNote        *string `sil:"F1691"`
	AddingTotalizerfunction *int    `sil:"F1802"`
	DetailSequence          *int    `sil:"F2770"`
	DetailType              *string `sil:"F2771"`
	TotalUnits              *string `sil:"F64"`
	TotalDollars            *string `sil:"F65"`
}
