package loc

// FeclBat is the FECL_BAT definition
type FeclBat struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	POSDescription           *string `sil:"F02"`
	TargetIdentifier         string  `sil:"F1000"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	MovementPeriodType       string  `sil:"F1031"`
	MovementFile             string  `sil:"F1033"`
	TotalizerNumber          int     `sil:"F1034"`
	ItemLinkCode             string  `sil:"F164"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	StoreResponsible         *string `sil:"F1964"`
	PriceMethod              *string `sil:"F33"`
	TotalUnits               *string `sil:"F64"`
	TotalDollars             *string `sil:"F65"`
	TotalWeight              *string `sil:"F67"`
	BatchIdentifier          *string `sil:"F902"`
}
