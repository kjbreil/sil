package loc

// ShfTab is the SHF_TAB definition
type ShfTab struct {
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	ShelfLife        *int    `sil:"F105"`
	AisleLocation    *string `sil:"F116"`
	ShelfLocation    *string `sil:"F117"`
	SideLocation     *string `sil:"F118"`
	Section          *string `sil:"F25"`
	ShelfPosX        *string `sil:"F2836"`
	ShelfPosY        *string `sil:"F2837"`
	ShelfPosZ        *string `sil:"F2838"`
	ShelfSegment     *string `sil:"F2839"`
	ShelfTagQuantity *int    `sil:"F94"`
	ShelfTagType     *string `sil:"F95"`
}
