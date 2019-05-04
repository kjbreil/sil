package loc

// FredeemBat is the FREDEEM_BAT definition
type FredeemBat struct {
	UPCCode              string  `sil:"F01,zeropad"`
	UPCCodeFormat        *int    `sil:"F07"`
	TargetIdentifier     string  `sil:"F1000"`
	RecordStatus         int     `sil:"F1001" default:"1"`
	RedeemDiscountPoints *int    `sil:"F1228"`
	RedeemDiscountAmount *string `sil:"F1229"`
	RedeemFreePoints     *int    `sil:"F1232"`
	RedeemStartDate      *string `sil:"F1234"`
	RedeemEndDate        *string `sil:"F1235"`
	PriceLevel           *int    `sil:"F126"`
	PointRedeemProgram   *string `sil:"F2744"`
	BatchIdentifier      *string `sil:"F902"`
}
