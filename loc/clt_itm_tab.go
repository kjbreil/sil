package loc

// CltItmTab is the CLT_ITM_TAB definition
type CltItmTab struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	PriceRoundingMethod      *string `sil:"F1005"`
	CustomerId               *string `sil:"F1148"`
	PriceEndDate             *string `sil:"F129"`
	SuggestedPrice           *string `sil:"F168"`
	SuggestedPriceMargin     *string `sil:"F1713"`
	SuggestedPriceRounding   *string `sil:"F1714"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	ContractRefNumber        *string `sil:"F1780"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	ContractNextPrice        *string `sil:"F2604"`
	ContractNextPriceQty     *string `sil:"F2605"`
	ContractNextStartDate    *string `sil:"F2606"`
	ContractNextStopDate     *string `sil:"F2607"`
	PriceMethod              *string `sil:"F33"`
	PriceStartDate           *string `sil:"F35"`
	PriceMargin              *string `sil:"F49"`
	PriceLimitedQty          *string `sil:"F62"`
	PriceLimited             *string `sil:"F63"`
	BatchIdentifier          *string `sil:"F902"`
}
