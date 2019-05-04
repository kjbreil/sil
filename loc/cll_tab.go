package loc

// CllTab is the CLL_TAB definition
type CllTab struct {
	TargetIdentifier         string  `sil:"F1000"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	CustomerId               *string `sil:"F1148"`
	AlternateCustNumber      *string `sil:"F1577"`
	AlternateCustType        *string `sil:"F1578"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	MainAltCode              *string `sil:"F1898"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	BatchIdentifier          *string `sil:"F902"`
}
