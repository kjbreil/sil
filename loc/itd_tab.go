package loc

// ItdTab is the ITD_TAB definition
type ItdTab struct {
	UPCCode               string  `sil:"F01,zeropad"`
	TargetIDentifier      string  `sil:"F1000" default:"PAL"`
	RecordStatus          int     `sil:"F1001" default:"1"`
	LastChangeDate        string  `sil:"F253" default:"NOW"`
	ItemAttribMethod      *string `sil:"F2834"`
	ItemAttribDescription *string `sil:"F2835"`
	BatchIDentifier       *string `sil:"F902"`
	CreatedByUser         *int    `sil:"F940"`
	ModifiedByUser        *int    `sil:"F941"`
}
