package loc

// TimTab is the TIM_TAB definition
type TimTab struct {
	MovementPeriodType   string  `sil:"F1031"`
	TransactionNumber    *int    `sil:"F1032"`
	TerminalStore        string  `sil:"F1056"`
	TerminalNumber       *string `sil:"F1057"`
	TransactionMode      *string `sil:"F1068"`
	LineNumber           *int    `sil:"F1101"`
	CustomerId           *string `sil:"F1148"`
	UserOrderNumber      *int    `sil:"F1185"`
	SoftField1           *string `sil:"F1301"`
	UserDepartmentNumber *int    `sil:"F1569"`
	ShiftNumber          *int    `sil:"F1819"`
	PunchInDatetime      *string `sil:"F1820"`
	PunchOutDatetime     *string `sil:"F1821"`
	SMSReference         *int    `sil:"F2014"`
	LastChangeDate       string  `sil:"F253" default:"NOW"`
	DateEnd              *string `sil:"F254"`
	LongDescription      *string `sil:"F255"`
	WageType             *int    `sil:"F2778"`
}
