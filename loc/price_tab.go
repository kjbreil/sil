package loc

// PriceTab is the PRICE_TAB definition
type PriceTab struct {
	UPCCode                    string  `sil:"F01,zeropad"`
	UPCCodeFormat              *int    `sil:"F07"`
	TargetIDentifier           string  `sil:"F1000" default:"PAL"`
	RecordStatus               int     `sil:"F1001" default:"1"`
	PriceRoundingMethod        *string `sil:"F1005"`
	ActivePriceQuantity        *string `sil:"F1006"`
	ActivePrice                *string `sil:"F1007"`
	ActivePriceChangedDate     *string `sil:"F1008"`
	ActivePriceChangedTime     *string `sil:"F1009"`
	NextPriceMethod            *string `sil:"F1010"`
	NextPriceSource            *string `sil:"F1011"`
	NextPriceQuantity          *string `sil:"F1012"`
	NextPrice                  *string `sil:"F1013"`
	NextPriceChangedDate       *string `sil:"F1014"`
	NextPriceChangedTime       *string `sil:"F1015"`
	DiscountPrintdetailtrack   *string `sil:"F109"`
	PriceDiscountAmount        *string `sil:"F111"`
	PriceDiscountPercent       *string `sil:"F112"`
	ActivePriceDescription     *string `sil:"F113"`
	InstorePrice               *string `sil:"F1133"`
	InstorePriceQty            *string `sil:"F1134"`
	AutoCalculateStorePrice    *string `sil:"F1135"`
	TPRPackagePrice            *string `sil:"F1186"`
	TPRPackagePriceQty         *string `sil:"F1187"`
	TPRLimitedQtyPrice         *string `sil:"F1188"`
	TPRLimitedQty              *string `sil:"F1189"`
	HoldRetailPrice            *string `sil:"F119"`
	TPRMixMatch                *int    `sil:"F1190"`
	TPRShopperPoints           *int    `sil:"F1191"`
	SaleShopperPoints          *int    `sil:"F1192"`
	SaleMixmatch               *int    `sil:"F1193"`
	InstoreOverTPR             *string `sil:"F1194"`
	InstoreOverSale            *string `sil:"F1195"`
	InstorePackagePrice        *string `sil:"F1196"`
	InstorePackagePriceQty     *string `sil:"F1197"`
	InstoreLimitedQtyPrice     *string `sil:"F1198"`
	InstoreLimitedQty          *string `sil:"F1199"`
	InstoreMixmatch            *int    `sil:"F1200"`
	InstorePriceMethod         *string `sil:"F1201"`
	InstoreShopperPoints       *int    `sil:"F1202"`
	ActivePackagePrice         *string `sil:"F1203"`
	ActivePackageQty           *string `sil:"F1204"`
	ActiveLimitedPrice         *string `sil:"F1205"`
	ActiveLimitedQty           *string `sil:"F1206"`
	ActiveMixmatch             *int    `sil:"F1207"`
	ActivePriceMethod          *string `sil:"F1208"`
	ActiveShopperPoints        *int    `sil:"F1209"`
	NextPackagePrice           *string `sil:"F1210"`
	NextPachagePriceQty        *string `sil:"F1211"`
	NextLimitedQtyPrice        *string `sil:"F1212"`
	NextLimitedQty             *string `sil:"F1213"`
	NextMixmatch               *int    `sil:"F1214"`
	NextShopperPoints          *int    `sil:"F1215"`
	InstoreStartDate           *string `sil:"F1216"`
	InstoreStopDate            *string `sil:"F1217"`
	TPRDiscountAmount          *string `sil:"F1218"`
	TPRDiscountPercent         *string `sil:"F1219"`
	LikeCode                   *string `sil:"F122"`
	SaleDiscountAmount         *string `sil:"F1220"`
	SaleDiscountPercent        *string `sil:"F1221"`
	InstoreDiscountAmount      *string `sil:"F1222"`
	InstoreDiscountPercent     *string `sil:"F1223"`
	ActiveDiscountAmount       *string `sil:"F1224"`
	ActiveDiscountPercent      *string `sil:"F1225"`
	NextDiscountAmount         *string `sil:"F1226"`
	NextDiscountPercent        *string `sil:"F1227"`
	RedeemDiscountPoints       *int    `sil:"F1228"`
	RedeemDiscountAmount       *string `sil:"F1229"`
	LastLabelRegularPrice      *string `sil:"F1231"`
	RedeemFreePoints           *int    `sil:"F1232"`
	RedeemStartDate            *string `sil:"F1234"`
	RedeemEndDate              *string `sil:"F1235"`
	PriceLevel                 *int    `sil:"F126"`
	PriceEndDate               *string `sil:"F129"`
	PriceEndTime               *string `sil:"F130"`
	TargetGrossMarginStartDate *string `sil:"F133"`
	SalePriceQty               *string `sil:"F135"`
	SalePrice                  *string `sil:"F136"`
	SalePriceStartDate         *string `sil:"F137"`
	SalePriceEndDate           *string `sil:"F138"`
	SalePackagePrice           *string `sil:"F139"`
	PricePackage               *string `sil:"F140"`
	PricePackageQty            *string `sil:"F142"`
	SalePackageQty             *string `sil:"F143"`
	SalePriceStartTime         *string `sil:"F144"`
	SalePriceEndTime           *string `sil:"F145"`
	SalePriceMethod            *string `sil:"F146"`
	SaleLimitedPriceQty        *string `sil:"F147"`
	SaleLimitedPrice           *string `sil:"F148"`
	ItemLinkCode               string  `sil:"F164"`
	AlternatePrice             *string `sil:"F166"`
	AlternatePriceQuantity     *string `sil:"F167"`
	SuggestedPrice             *string `sil:"F168"`
	PriceShopperPoints         *int    `sil:"F169"`
	SuggestedPriceMargin       *string `sil:"F1713"`
	SuggestedPriceRounding     *string `sil:"F1714"`
	ProhibitItemMarkdown       *string `sil:"F175"`
	MaintenanceOperatorLevel   *int    `sil:"F1759"`
	PricePackageMargin         *string `sil:"F1767"`
	PricePackageRounding       *string `sil:"F1768"`
	PriceLimitedMargin         *string `sil:"F1769"`
	NotInAdmissibleSpending    *string `sil:"F177"`
	PriceLimitedRounding       *string `sil:"F1770"`
	FlyerImpact                *string `sil:"F179"`
	SpareFieldPRI1             *string `sil:"F1799"`
	SpareFieldPRI2             *string `sil:"F1800"`
	FlyerPage                  *int    `sil:"F1801"`
	AddingTotalizerfunction    *int    `sil:"F1802"`
	TimeIncluded               *string `sil:"F1803"`
	SpareFieldPRI6             *string `sil:"F1804"`
	LabelPrice                 *string `sil:"F1805"`
	NormalizedPrice            *string `sil:"F1806"`
	TPRPrice                   *string `sil:"F181"`
	TPRPriceQty                *string `sil:"F182"`
	TPRStartDate               *string `sil:"F183"`
	TPREndDate                 *string `sil:"F184"`
	CommissionDollars          *string `sil:"F1885"`
	StoreResponsible           *string `sil:"F1964"`
	TargetDollarMarkup         *string `sil:"F205"`
	MeasureSellPack            *string `sil:"F21"`
	PrintLabelOnNextBatch      *string `sil:"F2119"`
	TPRPriceMethod             *string `sil:"F221"`
	OverrideNumber             *string `sil:"F232"`
	LastChangeDate             string  `sil:"F253" default:"NOW"`
	LabelPriceMethod           *string `sil:"F2569"`
	LabelPriceQuantity         *string `sil:"F2570"`
	LabelPackagePrice          *string `sil:"F2571"`
	LabelPackagePriceQty       *string `sil:"F2572"`
	LabelLimitedQtyPrice       *string `sil:"F2573"`
	LabelLimitedQty            *string `sil:"F2574"`
	LabelMixmatch              *int    `sil:"F2575"`
	LabelShopperPoints         *int    `sil:"F2576"`
	LabelDiscountAmount        *string `sil:"F2577"`
	LabelDiscountPercent       *string `sil:"F2578"`
	LabelPriceSource           *string `sil:"F2579"`
	LabelQuantity              *int    `sil:"F2580"`
	TPRNumber                  *string `sil:"F2667"`
	SaleNumber                 *string `sil:"F2668"`
	Price1Description          *string `sil:"F2694"`
	Price2Description          *string `sil:"F2695"`
	Price3Description          *string `sil:"F2696"`
	PointRedeemProgram         *string `sil:"F2744"`
	SaleQuantityLimit          *string `sil:"F2955"`
	Price                      *string `sil:"F30"`
	PriceQty                   *string `sil:"F31"`
	PriceMixmatch              *int    `sil:"F32"`
	PriceMethod                *string `sil:"F33"`
	ChangePriceReason          *string `sil:"F34"`
	PriceStartDate             *string `sil:"F35"`
	PriceStartTime             *string `sil:"F36"`
	ChangePriceOriginator      *string `sil:"F37"`
	PriceAudit                 *string `sil:"F41"`
	AdFlag                     *string `sil:"F42"`
	ApplyItemDiscount          *string `sil:"F43"`
	PriceMargin                *string `sil:"F49"`
	InstoreOverRegularPrice    *string `sil:"F59"`
	PriceLimitedQty            *string `sil:"F62"`
	PriceLimited               *string `sil:"F63"`
	BatchIDentifier            *string `sil:"F902"`
	BatchCreator               *string `sil:"F903"`
}