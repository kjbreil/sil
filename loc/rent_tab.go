package loc

// RentTab is the RENT_TAB definition
type RentTab struct {
	UPCCode           string  `sil:"F01,zeropad"`
	TransactionNumber *int    `sil:"F1032"`
	TerminalStore     string  `sil:"F1056"`
	TerminalNumber    *string `sil:"F1057"`
	ReferenceNumber   *string `sil:"F1079"`
	LineNumber        *int    `sil:"F1101"`
	CustomerId        *string `sil:"F1148"`
	PriceLevel        *int    `sil:"F126"`
	Behavior          *string `sil:"F1785"`
	RentStartDate     *string `sil:"F1845"`
	RentEndDate       *string `sil:"F1846"`
	TotalDollars      *string `sil:"F65"`
}
