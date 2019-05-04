package loc

// RebateBat is the REBATE_BAT definition
type RebateBat struct {
	UPCCode          string  `sil:"F01,zeropad"`
	UPCCodeFormat    *int    `sil:"F07"`
	TargetIdentifier string  `sil:"F1000"`
	RecordStatus     int     `sil:"F1001" default:"1"`
	BuyingFormat     *string `sil:"F1184"`
	RebateOffInvoice *string `sil:"F1660"`
	RebateQualifier  *string `sil:"F1974"`
	RebateMinQty     *string `sil:"F1979"`
	RebateAmount     *string `sil:"F233"`
	RebateStartDate  *string `sil:"F234"`
	RebateEndDate    *string `sil:"F235"`
	RebateNumber     *string `sil:"F236"`
	RebatePercent    *string `sil:"F237"`
	RebateMaxQty     *string `sil:"F2567"`
	RebateRetailRule *string `sil:"F2592"`
	VendorId         *string `sil:"F27"`
	BatchIdentifier  *string `sil:"F902"`
}
