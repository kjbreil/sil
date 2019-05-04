package loc

// BmpTab is the BMP_TAB definition
type BmpTab struct {
	UPCCode        string  `sil:"F01,zeropad"`
	RecordStatus   int     `sil:"F1001" default:"1"`
	LastChangeDate string  `sil:"F253" default:"NOW"`
	ImageType      *string `sil:"F2926"`
	ImageAngle     *string `sil:"F2927"`
	ImageSize      *string `sil:"F2928"`
	ImagePath      *string `sil:"F2929"`
}
