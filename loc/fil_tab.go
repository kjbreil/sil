package loc

// FilTab is the FIL_TAB definition
type FilTab struct {
	FilterLookup    *string `sil:"F1025"`
	FilterCondition *string `sil:"F1026"`
	FilterKey       *string `sil:"F1027"`
	FilterMore      *string `sil:"F1028"`
	LineNumber      *int    `sil:"F1101"`
}
