package loc

// SalDetSav is the SAL_DET_SAV definition
type SalDetSav struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	TransactionNumber       *int    `sil:"F1032"`
	DescriptionRegistration *string `sil:"F1041"`
	TerminalStore           string  `sil:"F1056"`
	TerminalNumber          *string `sil:"F1057"`
	ReferenceNumber         *string `sil:"F1079"`
	AlphaParameter          *string `sil:"F1081"`
	LineNumber              *int    `sil:"F1101"`
	TransactionArchive      *string `sil:"F1505"`
	RegistrationNote        *string `sil:"F1691"`
	AddingTotalizerfunction *int    `sil:"F1802"`
	DetailSequence          *int    `sil:"F2770"`
	DetailType              *string `sil:"F2771"`
	TimestampTransArchive   *string `sil:"F2890"`
	TotalUnits              *string `sil:"F64"`
	TotalDollars            *string `sil:"F65"`
}
