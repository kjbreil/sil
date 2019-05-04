package loc

// SalRegSav is the SAL_REG_SAV definition
type SalRegSav struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	DepartmentCode           *int    `sil:"F03"`
	SubDepartmentCode        *int    `sil:"F04"`
	BottleDepositLink        *int    `sil:"F05"`
	TareWeightLink           *int    `sil:"F06"`
	StatusCode               *string `sil:"F08"`
	TaxFlag6                 *string `sil:"F100"`
	SizeCubic                *string `sil:"F1002"`
	ActivePriceQuantity      *string `sil:"F1006"`
	ActivePrice              *string `sil:"F1007"`
	TaxFlag7                 *string `sil:"F101"`
	ProhibitQuantity         *string `sil:"F102"`
	TransactionNumber        *int    `sil:"F1032"`
	TotalizerNumber          int     `sil:"F1034"`
	VendorCoupon             *string `sil:"F104"`
	DescriptionRegistration  *string `sil:"F1041"`
	TerminalStore            string  `sil:"F1056"`
	TerminalNumber           *string `sil:"F1057"`
	FollowSubdeptStatus      *string `sil:"F106"`
	FunctionCode             *int    `sil:"F1063"`
	RegistrationMode         *string `sil:"F1067"`
	VoidPreviousFlag         *string `sil:"F1069"`
	RefundKeyFlag            *string `sil:"F1070"`
	TentativeItemFlag        *string `sil:"F1071"`
	SalesPerson              *int    `sil:"F1072"`
	OrderedQuantity          *string `sil:"F1078"`
	ReferenceNumber          *string `sil:"F1079"`
	CouponRestricted         *string `sil:"F108"`
	ElapseTime               *string `sil:"F1080"`
	AlphaParameter           *string `sil:"F1081"`
	VoidProhibited           *string `sil:"F1086"`
	DiscountPrintdetailtrack *string `sil:"F109"`
	ElectronicCoupon         *string `sil:"F110"`
	LineNumber               *int    `sil:"F1101"`
	POSSpecificFlags         *string `sil:"F1120"`
	UserNumber               *int    `sil:"F1126"`
	ActivePriceDescription   *string `sil:"F113"`
	TicketTemplate           *string `sil:"F1136"`
	AllowPriceOverride       *string `sil:"F114"`
	NotInNetSale             *string `sil:"F115"`
	ShopperDiscount          *string `sil:"F1164"`
	ShelfLocation            *string `sil:"F117"`
	TareLinkValue            *string `sil:"F1178"`
	UserOrderNumber          *int    `sil:"F1185"`
	ActivePackagePrice       *string `sil:"F1203"`
	ActivePackageQty         *string `sil:"F1204"`
	ActiveLimitedPrice       *string `sil:"F1205"`
	ActiveLimitedQty         *string `sil:"F1206"`
	ActiveMixmatch           *int    `sil:"F1207"`
	ActivePriceMethod        *string `sil:"F1208"`
	ActiveShopperPoints      *int    `sil:"F1209"`
	ActiveDiscountAmount     *string `sil:"F1224"`
	ActiveDiscountPercent    *string `sil:"F1225"`
	ItemDiscountFct          *int    `sil:"F1239"`
	ProhibitReturns          *string `sil:"F124"`
	ItemDiscountAmount       *string `sil:"F1240"`
	ClientDiscountTotal      *string `sil:"F1241"`
	ProhibitRepeatKey        *string `sil:"F125"`
	DepartmentDiscount       *string `sil:"F1256"`
	PriceLevel               *int    `sil:"F126"`
	TotalRoundedTax          *string `sil:"F1263"`
	CouponMultiplication     *string `sil:"F149"`
	ProhibitDiscount         *string `sil:"F150"`
	TransactionArchive       *string `sil:"F1505"`
	EnvirfeesValue           *string `sil:"F1595"`
	EnvirfeesTlz             *int    `sil:"F1596"`
	BottleReturn             *string `sil:"F160"`
	CouponFamilyCodeprevious *int    `sil:"F163"`
	SuggestedPrice           *string `sil:"F168"`
	FreightWeightTotal       *string `sil:"F1683"`
	FreightCubicTotal        *string `sil:"F1684"`
	OperatorReceiptPrinted   *int    `sil:"F1687"`
	PriceShopperPoints       *int    `sil:"F169"`
	RegistrationNote         *string `sil:"F1691"`
	OperatorOrderEntered     *int    `sil:"F1693"`
	OperatorReceiptConfirm   *int    `sil:"F1694"`
	ShippingPieceCount       *string `sil:"F1699"`
	RestrictionCode          *int    `sil:"F170"`
	MinimumAgeCustomer       *int    `sil:"F171"`
	CommittedQty             *string `sil:"F1712"`
	BottleDepositTotal       *string `sil:"F1715"`
	EnvirfeesTotal           *string `sil:"F1716"`
	TareTotal                *string `sil:"F1717"`
	StoreCouponFct           *int    `sil:"F1718"`
	ElectstoreCouponAmt      *string `sil:"F1719"`
	ProhibitRefund           *string `sil:"F172"`
	StoreCouponAmt           *string `sil:"F1720"`
	VendorCouponFct          *int    `sil:"F1721"`
	ElectvendorCouponAmt     *string `sil:"F1722"`
	VendorCouponAmt          *string `sil:"F1723"`
	DoubleCouponAmt          *string `sil:"F1724"`
	ItemDiscountQty          *string `sil:"F1725"`
	ItemDiscount             *string `sil:"F1726"`
	ItemDiscount             *string `sil:"F1727"`
	ProportionalDiscountAmt  *string `sil:"F1728"`
	DepartmentDiscountTotal  *string `sil:"F1729"`
	ProhibitMultipleCoupon   *string `sil:"F173"`
	GrossSale                *string `sil:"F1730"`
	ProportionalDiscountFct  *int    `sil:"F1731"`
	BackOrderFilled          *string `sil:"F1732"`
	PackageDiscount          *string `sil:"F1733"`
	LimitedDiscount          *string `sil:"F1734"`
	EntryUnit                *string `sil:"F1739"`
	EntryWeight              *string `sil:"F1740"`
	EnvirfeesFlags           *string `sil:"F1741"`
	CostSection              *string `sil:"F1742"`
	ProhibitItemMarkdown     *string `sil:"F175"`
	NotInAdmissibleSpending  *string `sil:"F177"`
	WICEligible              *string `sil:"F178"`
	Behavior                 *string `sil:"F1785"`
	ReplaceAddingFunction    *int    `sil:"F1787"`
	StoreCouponCount         *string `sil:"F1789"`
	AddingTotalizerfunction  *int    `sil:"F1802"`
	TimeIncluded             *string `sil:"F1803"`
	LabelPrice               *string `sil:"F1805"`
	OpenAmount               *string `sil:"F1815"`
	OpenTotalizer            *int    `sil:"F1816"`
	PointsOnStoreECoupons    *int    `sil:"F1831"`
	RedeemedPoints           *int    `sil:"F1832"`
	DiscountFromPoints       *string `sil:"F1833"`
	POMatching               *string `sil:"F1834"`
	RegLastCost              *string `sil:"F1835"`
	ElectvendorCouponFct     *int    `sil:"F1860"`
	ElectvendorCouponCnt     *string `sil:"F1861"`
	ElectstoreCouponFct      *int    `sil:"F1862"`
	ElectstoreCouponCnt      *string `sil:"F1863"`
	VendorCouponCnt          *string `sil:"F1864"`
	SelectPkgPrice           *int    `sil:"F1874"`
	DoubleCpnCnt             *string `sil:"F1888"`
	ProportionalDiscPercent  *string `sil:"F1924"`
	ComparePriceQty          *string `sil:"F1925"`
	SPAREREGSALPOS           *string `sil:"F1926"`
	SPAREREGSALFPRICE        *string `sil:"F1927"`
	SPAREREGSALGPRICE        *string `sil:"F1928"`
	SPAREREGSALPOS           *int    `sil:"F1929"`
	SPAREREGSALGPRICE        *int    `sil:"F1930"`
	RedeemOnStoreECoupon     *int    `sil:"F1931"`
	ComparePrice             *string `sil:"F1932"`
	SPAREREGSALPOS           *string `sil:"F1933"`
	SPAREREGSALFPRICE        *string `sil:"F1934"`
	SPAREREGSALGPRICE        *string `sil:"F1935"`
	PassiveDiscountInfo      *string `sil:"F1936"`
	CostPlusPercent          *string `sil:"F1938"`
	WeightDivisor            *string `sil:"F24"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	DateEnd                  *string `sil:"F254"`
	Fee1                     *string `sil:"F2551"`
	Fee2                     *string `sil:"F2552"`
	Fee3                     *string `sil:"F2553"`
	Fee4                     *string `sil:"F2554"`
	Fee5                     *string `sil:"F2555"`
	SPAREREGSALPOS           *string `sil:"F2608"`
	RedeemOnStoreCoupon      *int    `sil:"F2609"`
	PODistributionMatch      *string `sil:"F2610"`
	PriceModification        *string `sil:"F2611"`
	SPAREREGSAL              *string `sil:"F2612"`
	SPAREHDRREGSAL           *int    `sil:"F2613"`
	SPAREHDRREGSAL           *string `sil:"F2614"`
	URL                      *string `sil:"F2660"`
	WeightNet                *string `sil:"F270"`
	PointRedeemProgram       *string `sil:"F2744"`
	StoreECouponMethod       *string `sil:"F2745"`
	VendorECouponMethod      *string `sil:"F2746"`
	PointsOnVendorECoupons   *int    `sil:"F2747"`
	RedeemOnVendorECoupon    *int    `sil:"F2748"`
	RedeemOnVendorCoupon     *int    `sil:"F2749"`
	PointsOnStoreCoupon      *int    `sil:"F2750"`
	PointsOnVendorCoupon     *int    `sil:"F2751"`
	StoreCouponMethod        *string `sil:"F2752"`
	VendorCouponMethod       *string `sil:"F2753"`
	OpenQuantity             *string `sil:"F2860"`
	BagCount                 *int    `sil:"F2861"`
	SPAREREGSAL              *int    `sil:"F2862"`
	SPAREREGSAL              *int    `sil:"F2863"`
	SPAREREGSAL              *string `sil:"F2865"`
	SPAREREGSAL              *string `sil:"F2866"`
	SPAREREGSAL              *string `sil:"F2867"`
	SPAREREGSAL              *string `sil:"F2869"`
	SaleAsDiscount           *string `sil:"F2870"`
	SPAREREGSAL              *string `sil:"F2871"`
	TimestampTransArchive    *string `sil:"F2890"`
	Price                    *string `sil:"F30"`
	PriceQty                 *string `sil:"F31"`
	PromotionCode            *string `sil:"F383"`
	ApplyItemDiscount        *string `sil:"F43"`
	BottleDepositValue       *string `sil:"F50"`
	ExciseTaxAmount          *string `sil:"F60"`
	TaxExemptAmount          *string `sil:"F61"`
	TotalUnits               *string `sil:"F64"`
	TotalDollars             *string `sil:"F65"`
	TotalWeight              *string `sil:"F67"`
	CouponFamilyCode         *int    `sil:"F77"`
	FoodStamp                *string `sil:"F79"`
	FSA                      *string `sil:"F80"`
	TaxFlag1                 *string `sil:"F81"`
	ScalableItem             *string `sil:"F82"`
	RequirePriceEntry        *string `sil:"F83"`
	StoreCoupon              *string `sil:"F88"`
	BatchCreator             *string `sil:"F903"`
	TaxFlag2                 *string `sil:"F96"`
	TaxFlag3                 *string `sil:"F97"`
	TaxFlag4                 *string `sil:"F98"`
	TaxFlag5                 *string `sil:"F99"`
}
