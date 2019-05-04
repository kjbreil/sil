package loc

// GasHose is the GAS_HOSE definition
type GasHose struct {
	UPCCode              string  `sil:"F01,zeropad"`
	TargetIDentifier     string  `sil:"F1000" default:"PAL"`
	PriceLevel           *int    `sil:"F126"`
	GasHoseID            *string `sil:"F1610"`
	GasDollarsEndCounter *string `sil:"F1611"`
	GasVolumeEndCounter  *string `sil:"F1616"`
	GasHoseDescriptor    *string `sil:"F1625"`
}
