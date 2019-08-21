package loc

// BtlTab is the BTL_TAB definition
type BtlTab struct {
	BottleDepositLink     *int    `sil:"F05"`
	TaxFlag6              *string `sil:"F100"`
	TaxFlag7              *string `sil:"F101"`
	BottleLinkDescriptor  *string `sil:"F1020"`
	TotalizerNumber       int     `sil:"F1034"`
	FollowSubdeptStatus   *string `sil:"F106"`
	BottleEnvirDescriptor *string `sil:"F1594"`
	EnvirfeesValue        *string `sil:"F1595"`
	EnvirfeesTlz          *int    `sil:"F1596"`
	WICEligible           *string `sil:"F178"`
	CaseDepositValue      *string `sil:"F186"`
	BottleDepositValue    *string `sil:"F50"`
	FoodStamp             *string `sil:"F79"`
	TaxFlag1              *string `sil:"F81"`
	TaxFlag2              *string `sil:"F96"`
	TaxFlag3              *string `sil:"F97"`
	TaxFlag4              *string `sil:"F98"`
	TaxFlag5              *string `sil:"F99"`
}
