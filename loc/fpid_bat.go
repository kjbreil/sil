package loc

// FpidBat is the FPID_BAT definition
type FpidBat struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	PromotionCode    *string `sil:"F383"`
	BatchIDentifier  *string `sil:"F902"`
}
