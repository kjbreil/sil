package loc

// ClsTab is the CLS_TAB definition
type ClsTab struct {
	UPCCode      string  `sil:"F01,zeropad"`
	RecordStatus int     `sil:"F1001" default:"1"`
	ClassID      *string `sil:"F2935"`
}
