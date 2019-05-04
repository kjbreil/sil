package loc

// LocTab is the LOC_TAB definition
type LocTab struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	UPCCodeFormat            *int    `sil:"F07"`
	TargetIdentifier         string  `sil:"F1000"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	FacingY                  *int    `sil:"F1029"`
	FacingPosition           *string `sil:"F1030"`
	ShelfLife                *int    `sil:"F105"`
	AisleLocation            *string `sil:"F116"`
	ShelfLocation            *string `sil:"F117"`
	SideLocation             *string `sil:"F118"`
	MainLocation             *string `sil:"F157"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	ShelfTagBarcodeType      *string `sil:"F187"`
	StoreTreatement          *string `sil:"F190"`
	FacingX                  *int    `sil:"F191"`
	FacingZ                  *int    `sil:"F192"`
	LabelIndicator           *string `sil:"F1963"`
	StoreResponsible         *string `sil:"F1964"`
	Section                  *string `sil:"F25"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	PlanogramID              *string `sil:"F2851"`
	BatchIdentifier          *string `sil:"F902"`
	ShelfTagQuantity         *int    `sil:"F94"`
	ShelfTagType             *string `sil:"F95"`
}
