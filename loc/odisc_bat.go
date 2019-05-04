package loc

// OdiscBat is the ODISC_BAT definition
type OdiscBat struct {
	UPCCode            string  `sil:"F01,zeropad"`
	UPCCodeFormat      *int    `sil:"F07"`
	TargetIDentifier   string  `sil:"F1000" default:"PAL"`
	RecordStatus       int     `sil:"F1001" default:"1"`
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
	VendorID           *string `sil:"F27"`
	BatchIDentifier    *string `sil:"F902"`
}
