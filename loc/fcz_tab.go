package loc

// FczTab is the FCZ_TAB definition
type FczTab struct {
	TargetIdentifier      string  `sil:"F1000"`
	TotalizerNumber       int     `sil:"F1034"`
	FunctionCode          *int    `sil:"F1063"`
	FunctionToTlzOperator *string `sil:"F1065"`
}
