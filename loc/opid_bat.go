package loc

// OpidBat is the OPID_BAT definition
type OpidBat struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIdentifier string  `sil:"F1000"`
	PromotionCode    *string `sil:"F383"`
	BatchIdentifier  *string `sil:"F902"`
}
