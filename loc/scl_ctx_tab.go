package loc

// SclCtxTab is the SCL_CTX_TAB definition
type SclCtxTab struct {
	SubDepartmentCode        *int    `sil:"F04"`
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	COOLTextNumber           *int    `sil:"F2793"`
	COOLTextDescription      *string `sil:"F2794"`
	COOLTextType             *string `sil:"F2795"`
	COOLTextFontSize         *int    `sil:"F2938"`
	BatchIDentifier          *string `sil:"F902"`
}
