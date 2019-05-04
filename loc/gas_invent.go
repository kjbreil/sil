package loc

// GasInvent is the GAS_INVENT definition
type GasInvent struct {
	TerminalStore       string  `sil:"F1056"`
	ReferenceNumber     *string `sil:"F1079"`
	GasDollarsSales     *string `sil:"F1615"`
	GasVolumeSales      *string `sil:"F1620"`
	GasTankDescriptor   *string `sil:"F1626"`
	GasPrevInventoryVol *string `sil:"F1628"`
	GasPrevInventoryAmt *string `sil:"F1629"`
	GasDeliveryVol      *string `sil:"F1630"`
	GasDeliveryAmt      *string `sil:"F1631"`
	GasSalesCost        *string `sil:"F1632"`
	GasEndInventoryVol  *string `sil:"F1633"`
	GasEndInventoryAmt  *string `sil:"F1634"`
	GasTankDip          *string `sil:"F1635"`
	GasTankInvVol       *string `sil:"F1636"`
	GasTankInvAmt       *string `sil:"F1637"`
	GasTankVarVol       *string `sil:"F1638"`
	GasTankVarAmt       *string `sil:"F1639"`
	GasWaterDip         *string `sil:"F1640"`
	GasWaterVol         *string `sil:"F1641"`
	GasTankTemp         *string `sil:"F1842"`
	GasDeliveryGross    *string `sil:"F1847"`
	GasDeliveryTemp     *string `sil:"F1848"`
	GasTankDip2         *string `sil:"F1857"`
	GasWaterDip2        *string `sil:"F1858"`
	GasTankTemp2        *string `sil:"F1859"`
	LastChangeDate      string  `sil:"F253" default:"NOW"`
	DateEnd             *string `sil:"F254"`
}
