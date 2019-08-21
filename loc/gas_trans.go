package loc

// GasTrans is the GAS_TRANS definition
type GasTrans struct {
	UPCCode             string  `sil:"F01,zeropad"`
	ActivePrice         *string `sil:"F1007"`
	TransactionNumber   *int    `sil:"F1032"`
	MovementEndTime     *string `sil:"F1036"`
	TerminalStore       string  `sil:"F1056"`
	TerminalNumber      *string `sil:"F1057"`
	FunctionCode        *int    `sil:"F1063"`
	LineNumber          *int    `sil:"F1101"`
	UserOrderNumber     *int    `sil:"F1185"`
	PriceLevel          *int    `sil:"F126"`
	GasHoseID           *string `sil:"F1610"`
	ShiftNumber         *int    `sil:"F1819"`
	PassiveDiscountInfo *string `sil:"F1936"`
	LastChangeDate      string  `sil:"F253" default:"NOW"`
	DateEnd             *string `sil:"F254"`
	ConsoleTransNumber  *int    `sil:"F2719"`
	ConsoleTransService *int    `sil:"F2720"`
	ConsoleTransStatus  *int    `sil:"F2721"`
	GasPrepaySMSTerm    *string `sil:"F2722"`
	GasPrepaySMSTrans   *int    `sil:"F2723"`
	GasPrepaySMSLine    *int    `sil:"F2724"`
	GasPrepaySMSOper    *int    `sil:"F2725"`
	GasPrepayAmount     *string `sil:"F2726"`
	ConsoleTransType    *int    `sil:"F2782"`
	TotalDollars        *string `sil:"F65"`
	TotalWeight         *string `sil:"F67"`
}
