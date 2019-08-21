package loc

// RcpTab is the RCP_TAB definition
type RcpTab struct {
	UPCCode       string  `sil:"F01,zeropad"`
	RecipeCode    *string `sil:"F2784"`
	RecipeQty     *string `sil:"F2785"`
	CutProductQty *string `sil:"F2788"`
}
