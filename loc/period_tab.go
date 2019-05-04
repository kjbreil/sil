package loc

// PeriodTab is the PERIOD_TAB definition
type PeriodTab struct {
	TargetIdentifier   string  `sil:"F1000"`
	MovementPeriodType string  `sil:"F1031"`
	TransactionNumber  *int    `sil:"F1032"`
	PeriodDescriptor   *string `sil:"F1951"`
	LastChangeDate     string  `sil:"F253" default:"NOW"`
	DateEnd            *string `sil:"F254"`
}
