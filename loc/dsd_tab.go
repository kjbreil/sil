package loc

// DsdTab is the DSD_TAB definition
type DsdTab struct {
	UPCCode                   string  `sil:"F01,zeropad"`
	UPCCodeFormat             *int    `sil:"F07"`
	TargetIDentifier          string  `sil:"F1000" default:"PAL"`
	RecordStatus              int     `sil:"F1001" default:"1"`
	KnownShrinkFactor         *string `sil:"F1123"`
	DepartmentNumberReceiving *int    `sil:"F15"`
	MaintenanceOperatorLevel  *int    `sil:"F1759"`
	StoreResponsible          *string `sil:"F1964"`
	LastChangeDate            string  `sil:"F253" default:"NOW"`
	TargetTurnOver            *string `sil:"F2697"`
	BackstoreCapacity         *string `sil:"F2698"`
	InventoryTracking         *string `sil:"F68"`
	ReorderFlag               *string `sil:"F69"`
	OrderTrigger              *string `sil:"F72"`
	SuggestedOrder            *string `sil:"F73"`
	MinimumInventory          *string `sil:"F74"`
	BatchIDentifier           *string `sil:"F902"`
}
