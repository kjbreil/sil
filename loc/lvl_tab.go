package loc

// LvlTab is the LVL_TAB definition
type LvlTab struct {
	TargetIdentifier     string  `sil:"F1000"`
	PriceLevelDescriptor *string `sil:"F1017"`
	PriceLevel           *int    `sil:"F126"`
}
