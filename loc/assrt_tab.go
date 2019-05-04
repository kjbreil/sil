package loc

// AssrtTab is the ASSRT_TAB definition
type AssrtTab struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	AssortmentType   *string `sil:"F2932"`
}
