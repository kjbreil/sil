package loc

// ObbackBat is the OBBACK_BAT definition
type ObbackBat struct {
	UPCCode             string  `sil:"F01,zeropad"`
	UPCCodeFormat       *int    `sil:"F07"`
	TargetIdentifier    string  `sil:"F1000"`
	RecordStatus        int     `sil:"F1001" default:"1"`
	BuyingFormat        *string `sil:"F1184"`
	BillBack1Amount     *string `sil:"F1662"`
	BillBack1Percent    *string `sil:"F1663"`
	BillBack2Amount     *string `sil:"F1664"`
	BillBack2Percent    *string `sil:"F1665"`
	BillBack3Amount     *string `sil:"F1666"`
	BillBack3Percent    *string `sil:"F1667"`
	BillBack4Descriptor *string `sil:"F1668"`
	BillBack4Amount     *string `sil:"F1669"`
	BillBack4Percent    *string `sil:"F1670"`
	VendorId            *string `sil:"F27"`
	BatchIdentifier     *string `sil:"F902"`
}
