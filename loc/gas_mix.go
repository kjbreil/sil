package loc

// GasMix is the GAS_MIX definition
type GasMix struct {
	TargetIdentifier  string  `sil:"F1000"`
	GasHoseId         *string `sil:"F1610"`
	GasTankModel      *string `sil:"F1622"`
	GasTankDescriptor *string `sil:"F1626"`
	GasTankSplit      *string `sil:"F1627"`
}
