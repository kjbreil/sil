package loc

// KitTab is the KIT_TAB definition
type KitTab struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	UPCCodeFormat            *int    `sil:"F07"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	PriceLevel               *int    `sil:"F126"`
	KitUPCLink               *string `sil:"F1507"`
	KitUPCFactor             *string `sil:"F1508"`
	KitUPCReason             *string `sil:"F1509"`
	KitUPCSplit              *string `sil:"F1510"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	BatchIDentifier          *string `sil:"F902"`
}
