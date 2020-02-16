package tables

// PRICE is a sample PRICE_TAB definition
type PRICE struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	Price            *string `sil:"F30"`
	PriceQty         *string `sil:"F31"`
	PriceMixmatch    *int    `sil:"F32"`
	PriceMethod      *string `sil:"F33"`
	BatchIDentifier  *string `sil:"F902"`
	BatchCreator     *string `sil:"F903"`
}
