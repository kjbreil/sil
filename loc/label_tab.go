package loc

// LabelTab is the LABEL_TAB definition
type LabelTab struct {
	UPCCode                string  `sil:"F01,zeropad"`
	POSDescription         *string `sil:"F02"`
	SubDepartmentCode      *int    `sil:"F04"`
	BottleDepositLink      *int    `sil:"F05"`
	StatusCode             *string `sil:"F08"`
	FacingPosition         *string `sil:"F1030"`
	ShelfLife              *int    `sil:"F105"`
	PriceDiscountAmount    *string `sil:"F111"`
	PriceDiscountPercent   *string `sil:"F112"`
	ActivePriceDescription *string `sil:"F113"`
	CustomerID             *string `sil:"F1148"`
	ShelfLocation          *string `sil:"F117"`
	PriceLevel             *int    `sil:"F126"`
	PriceEndDate           *string `sil:"F129"`
	PriceEndTime           *string `sil:"F130"`
	SoftField1             *string `sil:"F1301"`
	SoftField2             *string `sil:"F1302"`
	PricePackage           *string `sil:"F140"`
	SoftAlpha1             *string `sil:"F1401"`
	SoftAlpha2             *string `sil:"F1402"`
	PricePackageQty        *string `sil:"F142"`
	AlternateCode          *string `sil:"F154"`
	BrandDescription       *string `sil:"F155"`
	FamilyCode             *int    `sil:"F16"`
	ItemLinkCode           string  `sil:"F164"`
	AlternatePrice         *string `sil:"F166"`
	AlternatePriceQuantity *string `sil:"F167"`
	SuggestedPrice         *string `sil:"F168"`
	PriceShopperPoints     *int    `sil:"F169"`
	CategoryCode           *int    `sil:"F17"`
	MeasureWeightOrVolume  *string `sil:"F1744"`
	ReportCode             *int    `sil:"F18"`
	NormalizedPrice        *string `sil:"F1806"`
	CaseSize               *string `sil:"F19"`
	MeasureSellPack        *string `sil:"F21"`
	SizeDescription        *string `sil:"F22"`
	MeasureDescription     *string `sil:"F23"`
	WeightDivisor          *string `sil:"F24"`
	Section                *string `sil:"F25"`
	LastChangeDate         string  `sil:"F253" default:"NOW"`
	LongDescription        *string `sil:"F255"`
	LabelFromDevice        *string `sil:"F2588"`
	VendorCode             *string `sil:"F26"`
	VendorID               *string `sil:"F27"`
	ExpandedDescription    *string `sil:"F29"`
	Price                  *string `sil:"F30"`
	PriceQty               *string `sil:"F31"`
	PriceMixmatch          *int    `sil:"F32"`
	PriceMethod            *string `sil:"F33"`
	PriceStartDate         *string `sil:"F35"`
	PriceStartTime         *string `sil:"F36"`
	BaseCost               *string `sil:"F38"`
	PriceLimitedQty        *string `sil:"F62"`
	PriceLimited           *string `sil:"F63"`
	TaxFlag1               *string `sil:"F81"`
	ScalableItem           *string `sil:"F82"`
	BatchIDentifier        *string `sil:"F902"`
	BatchCreator           *string `sil:"F903"`
	ShelfTagQuantity       *int    `sil:"F94"`
	ShelfTagType           *string `sil:"F95"`
	TaxFlag2               *string `sil:"F96"`
	TaxFlag3               *string `sil:"F97"`
	TaxFlag4               *string `sil:"F98"`
}
