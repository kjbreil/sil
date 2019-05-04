package loc

// PubTab is the PUB_TAB definition
type PubTab struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	TargetIdentifier         string  `sil:"F1000"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	MovementFile             string  `sil:"F1033"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	AdvertisingName          *string `sil:"F1912"`
	AdvertisingDuration      *int    `sil:"F1913"`
	AdvertisingURL           *string `sil:"F1914"`
	AdvertisingForm          *string `sil:"F1915"`
	AdvertisingStartDate     *string `sil:"F1916"`
	AdvertisingStopDate      *string `sil:"F1917"`
	StoreResponsible         *string `sil:"F1964"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
}
