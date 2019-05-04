package loc

// CfgTab is the CFG_TAB definition
type CfgTab struct {
	TargetIdentifier string `sil:"F1000"`
	RecordStatus     int    `sil:"F1001" default:"1"`
	DateCreation     string `sil:"F1264" default:"NOW"`
	LastChangeDate   string `sil:"F253" default:"NOW"`
	ConfigName       string `sil:"F2845"`
	ConfigGroup      string `sil:"F2846"`
	ConfigValue      string `sil:"F2847"`
	CreatedByUser    *int   `sil:"F940"`
	ModifiedByUser   *int   `sil:"F941"`
}
