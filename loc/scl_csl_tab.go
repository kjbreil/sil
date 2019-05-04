package loc

// SclCslTab is the SCL_CSL_TAB definition
type SclCslTab struct {
	SubDepartmentCode        *int    `sil:"F04"`
	TargetIdentifier         string  `sil:"F1000"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	COOLTextNumber           *int    `sil:"F2793"`
	COOLTextType             *string `sil:"F2795"`
	COOLShortListNumber      *int    `sil:"F2797"`
	COOLShortListSequence    *int    `sil:"F2798"`
	BatchIdentifier          *string `sil:"F902"`
}
