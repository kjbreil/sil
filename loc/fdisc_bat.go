package loc

// FdiscBat is the FDISC_BAT definition
type FdiscBat struct {
	UPCCode            string  `sil:"F01,zeropad"`
	UPCCodeFormat      *int    `sil:"F07"`
	TargetIdentifier   string  `sil:"F1000"`
	RecordStatus       int     `sil:"F1001" default:"1"`
	BuyingFormat       *string `sil:"F1184"`
	DiscountAmount     *string `sil:"F1658"`
	DiscountOffInvoice *string `sil:"F1659"`
	DiscountQualifier  *string `sil:"F1973"`
	DiscountMinQty     *string `sil:"F1978"`
	DiscountPercent    *string `sil:"F228"`
	DiscountStartDate  *string `sil:"F229"`
	DiscountEndDate    *string `sil:"F230"`
	DiscountNumber     *string `sil:"F231"`
	DiscountMaxQty     *string `sil:"F2566"`
	DiscountRetailRule *string `sil:"F2593"`
	VendorId           *string `sil:"F27"`
	BatchIdentifier    *string `sil:"F902"`
}
