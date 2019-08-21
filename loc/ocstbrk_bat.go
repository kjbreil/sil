package loc

// OcstbrkBat is the OCSTBRK_BAT definition
type OcstbrkBat struct {
	UPCCode             string  `sil:"F01,zeropad"`
	UPCCodeFormat       *int    `sil:"F07"`
	TargetIDentifier    string  `sil:"F1000" default:"PAL"`
	RecordStatus        int     `sil:"F1001" default:"1"`
	CostBreakOffInvoice *string `sil:"F1661"`
	CostBreakQualifier  *string `sil:"F1975"`
	CostBreakEndDate    *string `sil:"F219"`
	CostBreakMinQty     *string `sil:"F224"`
	CostBreakAmount     *string `sil:"F225"`
	CostBreakPercent    *string `sil:"F226"`
	CostBreakStartDate  *string `sil:"F227"`
	CostBreakMaxQty     *string `sil:"F2568"`
	CstbrkRetailRule    *string `sil:"F2594"`
	VendorID            *string `sil:"F27"`
	BatchIDentifier     *string `sil:"F902"`
}
