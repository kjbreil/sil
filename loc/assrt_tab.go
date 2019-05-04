package loc

// AssrtTab is the ASSRT_TAB definition
type AssrtTab struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIdentifier string  `sil:"F1000"`
	AssortmentType   *string `sil:"F2932"`
}
