package loc

// DssTab is the DSS_TAB definition
type DssTab struct {
	LastChangeDate    string  `sil:"F253" default:"NOW"`
	CustomPriority    *int    `sil:"F2727"`
	CustomAuthor      *string `sil:"F2728"`
	CustomOption      *string `sil:"F2729"`
	CustomDestination *string `sil:"F2730"`
	CustomScript      *string `sil:"F2731"`
	CustomFileDate    *string `sil:"F2732"`
	CustomSignature   *string `sil:"F2733"`
}
