package loc

// OposBat is the OPOS_BAT definition
type OposBat struct {
	UPCCode                    string  `sil:"F01,zeropad"`
	POSDescription             *string `sil:"F02"`
	DepartmentCode             *int    `sil:"F03"`
	SubDepartmentCode          *int    `sil:"F04"`
	BottleDepositLink          *int    `sil:"F05"`
	TareWeightLink             *int    `sil:"F06"`
	UPCCodeFormat              *int    `sil:"F07"`
	StatusCode                 *string `sil:"F08"`
	StatusIndicatorDate        *string `sil:"F09"`
	TaxFlag6                   *string `sil:"F100"`
	TargetIDentifier           string  `sil:"F1000" default:"PAL"`
	RecordStatus               int     `sil:"F1001" default:"1"`
	TaxFlag7                   *string `sil:"F101"`
	ProhibitQuantity           *string `sil:"F102"`
	VendorCoupon               *string `sil:"F104"`
	FollowSubdeptStatus        *string `sil:"F106"`
	RecordItemSale             *string `sil:"F107"`
	CouponRestricted           *string `sil:"F108"`
	RentAbleQuantity           *int    `sil:"F1099"`
	ElectronicCoupon           *string `sil:"F110"`
	POSSpecificFlags           *string `sil:"F1120"`
	CommissionRate             *string `sil:"F1124"`
	TicketTemplate             *string `sil:"F1136"`
	TareWeightProportional     *string `sil:"F1138"`
	MinimumAgeOperator         *int    `sil:"F1139"`
	AllowPriceOverride         *string `sil:"F114"`
	NotInNetSale               *string `sil:"F115"`
	AllowManualWeight          *string `sil:"F121"`
	PLUCode                    *string `sil:"F123"`
	AllowRaincheckTicket       *string `sil:"F1236"`
	ScaleDivisorSeeF24         *string `sil:"F1237"`
	ProhibitReturns            *string `sil:"F124"`
	ProhibitRepeatKey          *string `sil:"F125"`
	LinkQuantityLimit          *string `sil:"F141"`
	CouponMultiplication       *string `sil:"F149"`
	ProhibitDiscount           *string `sil:"F150"`
	LinkReasonCode             *int    `sil:"F153"`
	DepositItem                *string `sil:"F158"`
	RefundItem                 *string `sil:"F159"`
	BottleReturn               *string `sil:"F160"`
	MiscReceipt                *string `sil:"F161"`
	MiscPayout                 *string `sil:"F162"`
	CouponFamilyCodeprevious   *int    `sil:"F163"`
	RestrictionCode            *int    `sil:"F170"`
	MinimumAgeCustomer         *int    `sil:"F171"`
	ProhibitRefund             *string `sil:"F172"`
	ProhibitMultipleCoupon     *string `sil:"F173"`
	ControlledProductIndicator *string `sil:"F1735"`
	TaxIncluded                *string `sil:"F174"`
	KeyedDepartmentOverride    *string `sil:"F176"`
	WICEligible                *string `sil:"F178"`
	Behavior                   *string `sil:"F1785"`
	NormalizedUOMTag           *string `sil:"F1786"`
	ReplaceAddingFunction      *int    `sil:"F1787"`
	ReceiptCopyCount           *int    `sil:"F1788"`
	StoreCouponCount           *string `sil:"F1789"`
	POSValidItem               *string `sil:"F188"`
	SendToScale                *string `sil:"F189"`
	AltDescPOS                 *string `sil:"F1892"`
	StoreResponsible           *string `sil:"F1964"`
	CompetitivePriceQty        *string `sil:"F209"`
	CompetitivePrice           *string `sil:"F210"`
	CompetitivePriceStartDate  *string `sil:"F211"`
	ComparableSize             *string `sil:"F217"`
	WeightDivisor              *string `sil:"F24"`
	URL                        *string `sil:"F2660"`
	ItemAliasForWIC            *string `sil:"F2930"`
	CouponOfferCode            *string `sil:"F302"`
	CouponExpirationDate       *string `sil:"F303"`
	CouponHouseholdID          *string `sil:"F304"`
	CouponRedemptionMultiple   *string `sil:"F306"`
	PromotionCode              *string `sil:"F383"`
	PriceMultipleAsQuantity    *string `sil:"F388"`
	PromptForSKU               *string `sil:"F397"`
	ItemPricingRequired        *string `sil:"F40"`
	BottleDepositValue         *string `sil:"F50"`
	ExciseTaxAmount            *string `sil:"F60"`
	TaxExemptAmount            *string `sil:"F61"`
	SalesActivityLevel         *int    `sil:"F66"`
	CouponFamilyCode           *int    `sil:"F77"`
	RequireValidation          *string `sil:"F78"`
	FoodStamp                  *string `sil:"F79"`
	FSA                        *string `sil:"F80"`
	TaxFlag1                   *string `sil:"F81"`
	ScalableItem               *string `sil:"F82"`
	RequirePriceEntry          *string `sil:"F83"`
	RequireVisualVerify        *string `sil:"F84"`
	RequireQuantity            *string `sil:"F85"`
	NotForSaleInStore          *string `sil:"F86"`
	RestrictedSale             *string `sil:"F87"`
	StoreCoupon                *string `sil:"F88"`
	BatchIDentifier            *string `sil:"F902"`
	DepositContainerCode       *string `sil:"F92"`
	TaxFlag2                   *string `sil:"F96"`
	TaxFlag3                   *string `sil:"F97"`
	TaxFlag4                   *string `sil:"F98"`
	TaxFlag5                   *string `sil:"F99"`
}
