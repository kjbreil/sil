package loc

// GasHose is the GAS_HOSE definition
type GasHose struct {
	UPCCode              string  `sil:"F01,zeropad"`
	TargetIdentifier     string  `sil:"F1000"`
	PriceLevel           *int    `sil:"F126"`
	GasHoseId            *string `sil:"F1610"`
	GasDollarsEndCounter *string `sil:"F1611"`
	GasVolumeEndCounter  *string `sil:"F1616"`
	GasHoseDescriptor    *string `sil:"F1625"`
}
