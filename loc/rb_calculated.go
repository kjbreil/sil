package loc

// RbCalculated is the RB_CALCULATED definition
type RbCalculated struct {
	UPCCode            string  `sil:"F01,zeropad"`
	POSDescription     *string `sil:"F02"`
	ProductDescription *string `sil:"F256"`
}
