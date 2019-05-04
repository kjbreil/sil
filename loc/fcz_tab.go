package loc

// FczTab is the FCZ_TAB definition
type FczTab struct {
	TargetIDentifier      string  `sil:"F1000" default:"PAL"`
	TotalizerNumber       int     `sil:"F1034"`
	FunctionCode          *int    `sil:"F1063"`
	FunctionToTlzOperator *string `sil:"F1065"`
}
