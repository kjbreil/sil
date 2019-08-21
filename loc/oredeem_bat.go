package loc

// OredeemBat is the OREDEEM_BAT definition
type OredeemBat struct {
	UPCCode              string  `sil:"F01,zeropad"`
	UPCCodeFormat        *int    `sil:"F07"`
	TargetIDentifier     string  `sil:"F1000" default:"PAL"`
	RecordStatus         int     `sil:"F1001" default:"1"`
	RedeemDiscountPoints *int    `sil:"F1228"`
	RedeemDiscountAmount *string `sil:"F1229"`
	RedeemFreePoints     *int    `sil:"F1232"`
	RedeemStartDate      *string `sil:"F1234"`
	RedeemEndDate        *string `sil:"F1235"`
	PriceLevel           *int    `sil:"F126"`
	PointRedeemProgram   *string `sil:"F2744"`
	BatchIDentifier      *string `sil:"F902"`
}
