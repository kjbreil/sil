package loc

// SclCptTab is the SCL_CPT_TAB definition
type SclCptTab struct {
	SubDepartmentCode        *int    `sil:"F04"`
	TargetIdentifier         string  `sil:"F1000"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	COOLClassNumber          *int    `sil:"F2799"`
	COOLProdNumber           *int    `sil:"F2804"`
	COOLProdText             *string `sil:"F2805"`
	BatchIdentifier          *string `sil:"F902"`
}
