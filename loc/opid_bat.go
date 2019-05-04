package loc

// OpidBat is the OPID_BAT definition
type OpidBat struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	PromotionCode    *string `sil:"F383"`
	BatchIDentifier  *string `sil:"F902"`
}
