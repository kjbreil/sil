package loc

// OpriceBat is the OPRICE_BAT definition
type OpriceBat struct {
	UPCCode                    string  `sil:"F01,zeropad"`
	UPCCodeFormat              *int    `sil:"F07"`
	TargetIdentifier           string  `sil:"F1000"`
	RecordStatus               int     `sil:"F1001" default:"1"`
	PriceRoundingMethod        *string `sil:"F1005"`
	PriceDiscountAmount        *string `sil:"F111"`
	PriceDiscountPercent       *string `sil:"F112"`
	PriceLevel                 *int    `sil:"F126"`
	PriceEndDate               *string `sil:"F129"`
	PriceEndTime               *string `sil:"F130"`
	TargetGrossMarginStartDate *string `sil:"F133"`
	PricePackage               *string `sil:"F140"`
	PricePackageQty            *string `sil:"F142"`
	SuggestedPrice             *string `sil:"F168"`
	PriceShopperPoints         *int    `sil:"F169"`
	SuggestedPriceMargin       *string `sil:"F1713"`
	SuggestedPriceRounding     *string `sil:"F1714"`
	PricePackageMargin         *string `sil:"F1767"`
	PricePackageRounding       *string `sil:"F1768"`
	PriceLimitedMargin         *string `sil:"F1769"`
	PriceLimitedRounding       *string `sil:"F1770"`
	SPAREREGSALFPRICE          *string `sil:"F1927"`
	SPAREREGSALFPRICE          *string `sil:"F1934"`
	StoreResponsible           *string `sil:"F1964"`
	TargetDollarMarkup         *string `sil:"F205"`
	Price                      *string `sil:"F30"`
	PriceQty                   *string `sil:"F31"`
	PriceMixmatch              *int    `sil:"F32"`
	PriceMethod                *string `sil:"F33"`
	ChangePriceReason          *string `sil:"F34"`
	PriceStartDate             *string `sil:"F35"`
	PriceStartTime             *string `sil:"F36"`
	ChangePriceOriginator      *string `sil:"F37"`
	PriceMargin                *string `sil:"F49"`
	PriceLimitedQty            *string `sil:"F62"`
	PriceLimited               *string `sil:"F63"`
	BatchIdentifier            *string `sil:"F902"`
}
