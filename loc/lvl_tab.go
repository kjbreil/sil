package loc

// LvlTab is the LVL_TAB definition
type LvlTab struct {
	TargetIDentifier     string  `sil:"F1000" default:"PAL"`
	PriceLevelDescriptor *string `sil:"F1017"`
	PriceLevel           *int    `sil:"F126"`
}
