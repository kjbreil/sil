package loc

// UntTab is the UNT_TAB definition
type UntTab struct {
	UnitDescriptor          *string `sil:"F1503"`
	NormalizedUOMTag        *string `sil:"F1786"`
	MeasureDescription      *string `sil:"F23"`
	NormalizedUOMMultiplier *string `sil:"F2876"`
	NormalizedUOMMain       *string `sil:"F2877"`
	NormalizedUOMDesc       *string `sil:"F2878"`
}
