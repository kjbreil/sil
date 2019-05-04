package loc

// BioTab is the BIO_TAB definition
type BioTab struct {
	RecordStatus   int     `sil:"F1001" default:"1"`
	FunctionCode   *int    `sil:"F1063"`
	CustomerID     *string `sil:"F1148"`
	SecretNumber   *string `sil:"F1581"`
	BioData1       *string `sil:"F1849"`
	LastChangeDate string  `sil:"F253" default:"NOW"`
}
