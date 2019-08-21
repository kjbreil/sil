package loc

// OrebateBat is the OREBATE_BAT definition
type OrebateBat struct {
	UPCCode          string  `sil:"F01,zeropad"`
	UPCCodeFormat    *int    `sil:"F07"`
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	RecordStatus     int     `sil:"F1001" default:"1"`
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
	VendorID         *string `sil:"F27"`
	BatchIDentifier  *string `sil:"F902"`
}
