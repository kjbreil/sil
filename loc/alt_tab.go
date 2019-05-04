package loc

// AltTab is the ALT_TAB definition
type AltTab struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	PriceLevel               *int    `sil:"F126"`
	AlternateCode            *string `sil:"F154"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	SelectPkgPrice           *int    `sil:"F1874"`
	MainAltCode              *string `sil:"F1898"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	MainCodeByFormat         *string `sil:"F2637"`
	URL                      *string `sil:"F2660"`
	BatchIDentifier          *string `sil:"F902"`
}
