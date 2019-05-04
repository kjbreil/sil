package loc

// ItzTab is the ITZ_TAB definition
type ItzTab struct {
	UPCCode          string  `sil:"F01,zeropad"`
	TargetIdentifier string  `sil:"F1000"`
	RecordStatus     int     `sil:"F1001" default:"1"`
	TotalizerNumber  int     `sil:"F1034"`
	LastChangeDate   string  `sil:"F253" default:"NOW"`
	ItemizerMethod   *string `sil:"F2823"`
	ItemizerFactor   *string `sil:"F2824"`
	BatchIdentifier  *string `sil:"F902"`
	CreatedByUser    *int    `sil:"F940"`
	ModifiedByUser   *int    `sil:"F941"`
}
