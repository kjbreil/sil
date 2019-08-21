package loc

// AccTab is the ACC_TAB definition
type AccTab struct {
	UPCCode             string  `sil:"F01,zeropad"`
	TargetIDentifier    string  `sil:"F1000" default:"PAL"`
	MovementFile        string  `sil:"F1033"`
	TotalizerNumber     int     `sil:"F1034"`
	TotalizerDescriptor *string `sil:"F1039"`
	TransactionMode     *string `sil:"F1068"`
	TTLzGroup1          *string `sil:"F1128"`
	TTLzGroup2          *string `sil:"F1129"`
	TTLzGLAccount       *string `sil:"F1130"`
	SequenceNumber      *int    `sil:"F1147"`
	DebitOrCreditEntry  *string `sil:"F1253"`
}
