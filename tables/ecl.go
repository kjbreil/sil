package tables

// ECL is the ECL_TAB definition
type ECL struct {
	UPCCode        string  `sil:"F01"`
	POSDescription *string `sil:"F02"`
	ItemLinkCode   string  `sil:"F164,zeropad"`
}
