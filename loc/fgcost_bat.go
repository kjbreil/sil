package loc

// FgcostBat is the FGCOST_BAT definition
type FgcostBat struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	UPCCodeFormat            *int    `sil:"F07"`
	StatusCode               *string `sil:"F08"`
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	ActiveVendor             *string `sil:"F1037"`
	OperatorResponsible      *int    `sil:"F1168"`
	AlternateItemCode        *string `sil:"F127"`
	SequenceNumber           *int    `sil:"F131"`
	VendorBrokerNumber       *string `sil:"F165"`
	BackOrderDays            *int    `sil:"F1685"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	MinimumCases             *string `sil:"F1760"`
	MinimumCubic             *string `sil:"F1761"`
	ReturnFees               *string `sil:"F1766"`
	CostMethod               *string `sil:"F1791"`
	SpareFieldCOS2           *string `sil:"F1792"`
	DeliveryDays             *int    `sil:"F1793"`
	DeliveryCycleDays        *int    `sil:"F1794"`
	SplitItemQty             *string `sil:"F1795"`
	SpareRECREGGCOST         *string `sil:"F1797"`
	CostForSupplier          *string `sil:"F1798"`
	InventoryMode            *string `sil:"F1875"`
	ReturnPolicy             *string `sil:"F1961"`
	CostComment              *string `sil:"F1976"`
	AvailabilityDate         *string `sil:"F216"`
	SplitItemCode            *string `sil:"F220"`
	VendorCode               *string `sil:"F26"`
	VendorSubstitutionPolicy *string `sil:"F2601"`
	SPARERECREGGCOST         *int    `sil:"F2624"`
	DefaultQuantity          *string `sil:"F2666"`
	CostMethodDate           *string `sil:"F2699"`
	VendorID                 *string `sil:"F27"`
	CaseUPCCode              *string `sil:"F28"`
	PalletCases              *string `sil:"F325"`
	PalletLayers             *string `sil:"F326"`
	FreightCaseWeight        *string `sil:"F327"`
	FreightCaseCube          *string `sil:"F328"`
	VendorAuthorizedItem     *string `sil:"F90"`
	BatchIDentifier          *string `sil:"F902"`
}
