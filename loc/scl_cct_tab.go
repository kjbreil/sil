package loc

// SclCctTab is the SCL_CCT_TAB definition
type SclCctTab struct {
	SubDepartmentCode        *int    `sil:"F04"`
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	COOLClassNumber          *int    `sil:"F2799"`
	COOLCountryNumber        *int    `sil:"F2802"`
	COOLCountryText          *string `sil:"F2803"`
	BatchIDentifier          *string `sil:"F902"`
}
