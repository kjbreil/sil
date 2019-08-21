package loc

// FcostBat is the FCOST_BAT definition
type FcostBat struct {
	UPCCode           string  `sil:"F01,zeropad"`
	UPCCodeFormat     *int    `sil:"F07"`
	TargetIDentifier  string  `sil:"F1000" default:"PAL"`
	RecordStatus      int     `sil:"F1001" default:"1"`
	CostFeeAmount     *string `sil:"F1121"`
	CostFeePercent    *string `sil:"F1122"`
	SplitCaseCost     *string `sil:"F120"`
	HalfCaseCost      *string `sil:"F185"`
	CaseSize          *string `sil:"F19"`
	StoreResponsible  *string `sil:"F1964"`
	ReceivingPackSize *string `sil:"F20"`
	BaseCostTime      *string `sil:"F212"`
	VendorID          *string `sil:"F27"`
	BaseCost          *string `sil:"F38"`
	BaseCostDate      *string `sil:"F39"`
	BatchIDentifier   *string `sil:"F902"`
}
