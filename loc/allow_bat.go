package loc

// AllowBat is the ALLOW_BAT definition
type AllowBat struct {
	UPCCode             string  `sil:"F01,zeropad"`
	UPCCodeFormat       *int    `sil:"F07"`
	TargetIdentifier    string  `sil:"F1000"`
	RecordStatus        int     `sil:"F1001" default:"1"`
	BuyingFormat        *string `sil:"F1184"`
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
	VendorId            *string `sil:"F27"`
	AllowanceMaxQty     *string `sil:"F329"`
	BatchIdentifier     *string `sil:"F902"`
}
