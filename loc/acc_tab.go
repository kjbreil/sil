package loc

// AccTab is the ACC_TAB definition
type AccTab struct {
	UPCCode             string  `sil:"F01,zeropad"`
	TargetIdentifier    string  `sil:"F1000"`
	MovementFile        string  `sil:"F1033"`
	TotalizerNumber     int     `sil:"F1034"`
	TotalizerDescriptor *string `sil:"F1039"`
	TransactionMode     *string `sil:"F1068"`
	TtlzGroup1          *string `sil:"F1128"`
	TtlzGroup2          *string `sil:"F1129"`
	TtlzGLAccount       *string `sil:"F1130"`
	SequenceNumber      *int    `sil:"F1147"`
	DebitOrCreditEntry  *string `sil:"F1253"`
}
