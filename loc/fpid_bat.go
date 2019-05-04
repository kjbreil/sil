package loc

// FpidBat is the FPID_BAT definition
type FpidBat struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIdentifier string  `sil:"F1000"`
	PromotionCode    *string `sil:"F383"`
	BatchIdentifier  *string `sil:"F902"`
}
