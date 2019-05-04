package loc

// OobjBat is the OOBJ_BAT definition
type OobjBat struct {
	UPCCode                                string  `sil:"F01,zeropad"`
	UPCCodeFormat                          *int    `sil:"F07"`
	TargetIdentifier                       string  `sil:"F1000"`
	RecordStatus                           int     `sil:"F1001" default:"1"`
	SizeCubic                              *string `sil:"F1002"`
	ContainerType                          *int    `sil:"F1004"`
	MeasurementSystem                      *int    `sil:"F11"`
	ManufacturerId                         *string `sil:"F1118"`
	GraphicFile                            *string `sil:"F1119"`
	OperatorResponsible                    *int    `sil:"F1168"`
	SizeHeight                             *string `sil:"F12"`
	SizeWidth                              *string `sil:"F13"`
	SizeDepth                              *string `sil:"F14"`
	BrandDescription                       *string `sil:"F155"`
	FamilyCode                             *int    `sil:"F16"`
	ShippingPieceCount                     *string `sil:"F1699"`
	CategoryCode                           *int    `sil:"F17"`
	BestBeforeDays                         *int    `sil:"F1736"`
	MedicalProductCode                     *string `sil:"F1737"`
	NDCDINNumber                           *string `sil:"F1738"`
	MeasureWeightOrVolume                  *string `sil:"F1744"`
	SpareFieldOBJ1                         *string `sil:"F1781"`
	SpareFieldOBJ2                         *string `sil:"F1782"`
	NutritionIndex                         *int    `sil:"F1783"`
	SpareFieldOBJ4                         *int    `sil:"F1784"`
	ReportCode                             *int    `sil:"F18"`
	ManufactureCode                        *string `sil:"F180"`
	ReportingDepartment                    *string `sil:"F193"`
	AltBrandDesc                           *string `sil:"F1939"`
	AltExpandedDesc                        *string `sil:"F1940"`
	AltSizeDesc                            *string `sil:"F1941"`
	AltLongDesc                            *string `sil:"F1942"`
	Classification                         *string `sil:"F1957"`
	TargetCustomerType                     *string `sil:"F1958"`
	TargetStoreType                        *string `sil:"F1959"`
	HandlingType                           *string `sil:"F1960"`
	MarketingJustification                 *string `sil:"F1962"`
	StoreResponsible                       *string `sil:"F1964"`
	MeasureSellPack                        *string `sil:"F21"`
	DUNSNumberPlusSuffix                   *string `sil:"F213"`
	AliasCode                              *string `sil:"F214"`
	AliasCodeFormat                        *int    `sil:"F215"`
	ComparableSizeUnitOfMeasureDescription *string `sil:"F218"`
	SizeDescription                        *string `sil:"F22"`
	MeasureDescription                     *string `sil:"F23"`
	LongDescription                        *string `sil:"F255"`
	ItemSubstitutionPolicy                 *string `sil:"F2600"`
	CompetitiveCode                        *string `sil:"F2693"`
	WeightNet                              *string `sil:"F270"`
	InterDeptCode                          *string `sil:"F2789"`
	ExpandedDescription                    *string `sil:"F29"`
	NACSCode                               *int    `sil:"F2931"`
	BatchIdentifier                        *string `sil:"F902"`
	AccountCode                            *string `sil:"F93"`
	CreatedByUser                          *int    `sil:"F940"`
	ModifiedByUser                         *int    `sil:"F941"`
}
