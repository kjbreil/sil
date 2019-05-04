package loc

// FtpriceBat is the FTPRICE_BAT definition
type FtpriceBat struct {
	UPCCode               string  `sil:"F01,zeropad"`
	UPCCodeFormat         *int    `sil:"F07"`
	TargetIdentifier      string  `sil:"F1000"`
	RecordStatus          int     `sil:"F1001" default:"1"`
	TPRPackagePrice       *string `sil:"F1186"`
	TPRPackagePriceQty    *string `sil:"F1187"`
	TPRLimitedQtyPrice    *string `sil:"F1188"`
	TPRLimitedQty         *string `sil:"F1189"`
	TPRMixMatch           *int    `sil:"F1190"`
	TPRShopperPoints      *int    `sil:"F1191"`
	TPRDiscountAmount     *string `sil:"F1218"`
	TPRDiscountPercent    *string `sil:"F1219"`
	PriceLevel            *int    `sil:"F126"`
	TPRPrice              *string `sil:"F181"`
	TPRPriceQty           *string `sil:"F182"`
	TPRStartDate          *string `sil:"F183"`
	TPREndDate            *string `sil:"F184"`
	TPRMarginRetail       *string `sil:"F1868"`
	TPRMarginPkg          *string `sil:"F1869"`
	TPRMarginLmt          *string `sil:"F1870"`
	TPRHostCompensation   *string `sil:"F1953"`
	TPRVendorCompensation *string `sil:"F1954"`
	TPRPriceMethod        *string `sil:"F221"`
	TPRNumber             *string `sil:"F2667"`
	ChangePriceReason     *string `sil:"F34"`
	ChangePriceOriginator *string `sil:"F37"`
	BatchIdentifier       *string `sil:"F902"`
}
