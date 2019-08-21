package loc

// FinstoreBat is the FINSTORE_BAT definition
type FinstoreBat struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	UPCCodeFormat           *int    `sil:"F07"`
	TargetIDentifier        string  `sil:"F1000" default:"PAL"`
	RecordStatus            int     `sil:"F1001" default:"1"`
	InstorePrice            *string `sil:"F1133"`
	InstorePriceQty         *string `sil:"F1134"`
	InstoreOverTPR          *string `sil:"F1194"`
	InstoreOverSale         *string `sil:"F1195"`
	InstorePackagePrice     *string `sil:"F1196"`
	InstorePackagePriceQty  *string `sil:"F1197"`
	InstoreLimitedQtyPrice  *string `sil:"F1198"`
	InstoreLimitedQty       *string `sil:"F1199"`
	InstoreMixmatch         *int    `sil:"F1200"`
	InstorePriceMethod      *string `sil:"F1201"`
	InstoreShopperPoints    *int    `sil:"F1202"`
	InstoreStartDate        *string `sil:"F1216"`
	InstoreStopDate         *string `sil:"F1217"`
	InstoreDiscountAmount   *string `sil:"F1222"`
	InstoreDiscountPercent  *string `sil:"F1223"`
	PriceLevel              *int    `sil:"F126"`
	InstoreMarginRetail     *string `sil:"F1871"`
	InstoreMarginPkg        *string `sil:"F1872"`
	InstoreMarginLmt        *string `sil:"F1873"`
	ChangePriceReason       *string `sil:"F34"`
	ChangePriceOriginator   *string `sil:"F37"`
	InstoreOverRegularPrice *string `sil:"F59"`
	BatchIDentifier         *string `sil:"F902"`
}
