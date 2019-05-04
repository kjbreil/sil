package loc

// OspriceBat is the OSPRICE_BAT definition
type OspriceBat struct {
	UPCCode                string  `sil:"F01,zeropad"`
	UPCCodeFormat          *int    `sil:"F07"`
	TargetIDentifier       string  `sil:"F1000" default:"PAL"`
	RecordStatus           int     `sil:"F1001" default:"1"`
	SaleShopperPoints      *int    `sil:"F1192"`
	SaleMixmatch           *int    `sil:"F1193"`
	SaleDiscountAmount     *string `sil:"F1220"`
	SaleDiscountPercent    *string `sil:"F1221"`
	PriceLevel             *int    `sil:"F126"`
	SalePriceQty           *string `sil:"F135"`
	SalePrice              *string `sil:"F136"`
	SalePriceStartDate     *string `sil:"F137"`
	SalePriceEndDate       *string `sil:"F138"`
	SalePackagePrice       *string `sil:"F139"`
	SalePackageQty         *string `sil:"F143"`
	SalePriceStartTime     *string `sil:"F144"`
	SalePriceEndTime       *string `sil:"F145"`
	SalePriceMethod        *string `sil:"F146"`
	SaleLimitedPriceQty    *string `sil:"F147"`
	SaleLimitedPrice       *string `sil:"F148"`
	FlyerImpact            *string `sil:"F179"`
	FlyerPage              *int    `sil:"F1801"`
	SaleMarginRetail       *string `sil:"F1865"`
	SaleMarginPkg          *string `sil:"F1866"`
	SaleMarginLmt          *string `sil:"F1867"`
	SaleHostCompensation   *string `sil:"F1955"`
	SaleVendorCompensation *string `sil:"F1956"`
	OverrideNumber         *string `sil:"F232"`
	SaleNumber             *string `sil:"F2668"`
	SaleQuantityLimit      *string `sil:"F2955"`
	ChangePriceReason      *string `sil:"F34"`
	ChangePriceOriginator  *string `sil:"F37"`
	AdFlag                 *string `sil:"F42"`
	BatchIDentifier        *string `sil:"F902"`
}
