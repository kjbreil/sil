package loc

// CatTab is the CAT_TAB definition
type CatTab struct {
	CategoryDescriptor  *string `sil:"F1023"`
	KnownShrinkFactor   *string `sil:"F1123"`
	CommissionRate      *string `sil:"F1124"`
	OperatorResponsible *int    `sil:"F1168"`
	CategoryCode        *int    `sil:"F17"`
	SubCategory1        *string `sil:"F1943"`
	SubCategory2        *string `sil:"F1944"`
	SubCategory3        *string `sil:"F1945"`
	SubDeptLink         *int    `sil:"F1946"`
	SubDeptLink2        *int    `sil:"F1947"`
	ShowPriority        *int    `sil:"F1965"`
	ShowFilter          *string `sil:"F1966"`
	SubDeptLink3        *int    `sil:"F2653"`
	SubDeptLink4        *int    `sil:"F2654"`
	PriceMargin         *string `sil:"F49"`
}
