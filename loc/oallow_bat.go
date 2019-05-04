package loc

// OallowBat is the OALLOW_BAT definition
type OallowBat struct {
	UPCCode             string  `sil:"F01,zeropad"`
	UPCCodeFormat       *int    `sil:"F07"`
	TargetIDentifier    string  `sil:"F1000" default:"PAL"`
	RecordStatus        int     `sil:"F1001" default:"1"`
	AllowanceQualifier  *string `sil:"F156"`
	AllowancePercent    *string `sil:"F1657"`
	AllowanceMinQty     *string `sil:"F1977"`
	AllowanceAmount     *string `sil:"F201"`
	AllowanceStartDate  *string `sil:"F202"`
	AllowanceEndDate    *string `sil:"F203"`
	AllowanceCode       *string `sil:"F204"`
	AllowanceNumber     *string `sil:"F222"`
	AllowanceOffInvoice *string `sil:"F223"`
	AllowRetailRule     *string `sil:"F2591"`
	VendorID            *string `sil:"F27"`
	AllowanceMaxQty     *string `sil:"F329"`
	BatchIDentifier     *string `sil:"F902"`
}
