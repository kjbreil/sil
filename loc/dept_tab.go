package loc

// DeptTab is the DEPT_TAB definition
type DeptTab struct {
	DepartmentCode          *int    `sil:"F03"`
	TaxFlag6                *string `sil:"F100"`
	TaxFlag7                *string `sil:"F101"`
	VendorCoupon            *string `sil:"F104"`
	CouponRestricted        *string `sil:"F108"`
	POSSpecificFlags        *string `sil:"F1120"`
	KnownShrinkFactor       *string `sil:"F1123"`
	CommissionRate          *string `sil:"F1124"`
	DepartmentGroup         *string `sil:"F1132"`
	MinimumAgeOperator      *int    `sil:"F1139"`
	AllowPriceOverride      *string `sil:"F114"`
	SequenceNumber          *int    `sil:"F1147"`
	NotInNetSale            *string `sil:"F115"`
	OperatorResponsible     *int    `sil:"F1168"`
	AllowManualWeight       *string `sil:"F121"`
	ProhibitReturns         *string `sil:"F124"`
	DepartmentDiscount      *string `sil:"F1256"`
	ProhibitDiscount        *string `sil:"F150"`
	RestrictionCode         *int    `sil:"F170"`
	MinimumAgeCustomer      *int    `sil:"F171"`
	ProhibitRefund          *string `sil:"F172"`
	NotInAdmissibleSpending *string `sil:"F177"`
	WICEligible             *string `sil:"F178"`
	Behavior                *string `sil:"F1785"`
	AltDescDept             *string `sil:"F1894"`
	ShowPriority            *int    `sil:"F1965"`
	ShowFilter              *string `sil:"F1966"`
	SPARE                   *string `sil:"F1967"`
	DepartmentDescription   *string `sil:"F238"`
	MaximumAmount           *string `sil:"F239"`
	MinimumAmount           *string `sil:"F240"`
	MaximumVoid             *string `sil:"F241"`
	MaximumRefund           *string `sil:"F242"`
	URL                     *string `sil:"F2660"`
	PriceMargin             *string `sil:"F49"`
	RequireValidation       *string `sil:"F78"`
	FoodStamp               *string `sil:"F79"`
	FSA                     *string `sil:"F80"`
	TaxFlag1                *string `sil:"F81"`
	ScalableItem            *string `sil:"F82"`
	StoreCoupon             *string `sil:"F88"`
	TaxFlag2                *string `sil:"F96"`
	TaxFlag3                *string `sil:"F97"`
	TaxFlag4                *string `sil:"F98"`
	TaxFlag5                *string `sil:"F99"`
}
