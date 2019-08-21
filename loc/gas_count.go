package loc

// GasCount is the GAS_COUNT definition
type GasCount struct {
	UPCCode                string  `sil:"F01,zeropad"`
	TransactionNumber      *int    `sil:"F1032"`
	MovementEndTime        *string `sil:"F1036"`
	TerminalStore          string  `sil:"F1056"`
	PriceLevel             *int    `sil:"F126"`
	GasHoseID              *string `sil:"F1610"`
	GasDollarsEndCounter   *string `sil:"F1611"`
	GasDollarsStartCounter *string `sil:"F1612"`
	GasDollarsPumpTest     *string `sil:"F1613"`
	GasDollarsAdjustment   *string `sil:"F1614"`
	GasDollarsSales        *string `sil:"F1615"`
	GasVolumeEndCounter    *string `sil:"F1616"`
	GasVolumeStartCounter  *string `sil:"F1617"`
	GasVolumePumpTest      *string `sil:"F1618"`
	GasVolumeAdjustment    *string `sil:"F1619"`
	GasVolumeSales         *string `sil:"F1620"`
	GasDollarsCalculated   *string `sil:"F1621"`
	LastChangeDate         string  `sil:"F253" default:"NOW"`
	DateEnd                *string `sil:"F254"`
	Price                  *string `sil:"F30"`
}
