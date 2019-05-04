package loc

// CkzTab is the CKZ_TAB definition
type CkzTab struct {
	RecordStatus    int     `sil:"F1001" default:"1"`
	TotalizerNumber int     `sil:"F1034"`
	LastChangeDate  string  `sil:"F253" default:"NOW"`
	UserLimit       *string `sil:"F2825"`
	UserLimitMethod *string `sil:"F2826"`
	UserGroup       *string `sil:"F2827"`
	BatchIDentifier *string `sil:"F902"`
	CreatedByUser   *int    `sil:"F940"`
	ModifiedByUser  *int    `sil:"F941"`
}
