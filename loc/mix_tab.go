package loc

// MixTab is the MIX_TAB definition
type MixTab struct {
	RecordStatus    int     `sil:"F1001" default:"1"`
	MixDescriptor   *string `sil:"F1019"`
	PriceMixmatch   *int    `sil:"F32"`
	BatchIDentifier *string `sil:"F902"`
}
