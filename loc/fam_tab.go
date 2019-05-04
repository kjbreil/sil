package loc

// FamTab is the FAM_TAB definition
type FamTab struct {
	FamilyGroupDescriptor *string `sil:"F1040"`
	KnownShrinkFactor     *string `sil:"F1123"`
	CommissionRate        *string `sil:"F1124"`
	OperatorResponsible   *int    `sil:"F1168"`
	FamilyCode            *int    `sil:"F16"`
	ShowPriority          *int    `sil:"F1965"`
	ShowFilter            *string `sil:"F1966"`
	PriceMargin           *string `sil:"F49"`
}
