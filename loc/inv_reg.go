package loc

// InvReg is the INV_REG definition
type InvReg struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	DepartmentCode          *int    `sil:"F03"`
	SubDepartmentCode       *int    `sil:"F04"`
	BottleDepositLink       *int    `sil:"F05"`
	CaseQuantity            *string `sil:"F1003"`
	ActivePriceQuantity     *string `sil:"F1006"`
	ActivePrice             *string `sil:"F1007"`
	TransactionNumber       *int    `sil:"F1032"`
	TotalizerNumber         int     `sil:"F1034"`
	MovementStartTime       *string `sil:"F1035"`
	MovementEndTime         *string `sil:"F1036"`
	DescriptionRegistration *string `sil:"F1041"`
	TerminalStore           string  `sil:"F1056"`
	FunctionCode            *int    `sil:"F1063"`
	RegistrationMode        *string `sil:"F1067"`
	TransactionMode         *string `sil:"F1068"`
	SalesPerson             *int    `sil:"F1072"`
	ReferenceNumber         *string `sil:"F1079"`
	AlphaParameter          *string `sil:"F1081"`
	LineNumber              *int    `sil:"F1101"`
	UnitNetCost             *string `sil:"F1140"`
	SequenceNumber          *int    `sil:"F131"`
	NextCostNet             *string `sil:"F151"`
	TargetStore             *string `sil:"F1690"`
	RegistrationNote        *string `sil:"F1691"`
	OperatorReceiptConfirm  *int    `sil:"F1694"`
	InventoryQtytare        *string `sil:"F1706"`
	InventoryCost           *string `sil:"F1707"`
	InventoryRetail         *string `sil:"F1708"`
	PackagePrice            *string `sil:"F1813"`
	SpareREGINV             *int    `sil:"F1814"`
	InventoryMode           *string `sil:"F1875"`
	CaseSize                *string `sil:"F19"`
	LastBaseCost            *string `sil:"F195"`
	CaseNetCost             *string `sil:"F196"`
	LastChangeDate          string  `sil:"F253" default:"NOW"`
	DateEnd                 *string `sil:"F254"`
	VendorCode              *string `sil:"F26"`
	URL                     *string `sil:"F2660"`
	VendorID                *string `sil:"F27"`
	WeightNet               *string `sil:"F270"`
	Price                   *string `sil:"F30"`
	PriceQty                *string `sil:"F31"`
	BaseCost                *string `sil:"F38"`
	TotalUnits              *string `sil:"F64"`
	TotalDollars            *string `sil:"F65"`
	TotalWeight             *string `sil:"F67"`
	UnitQuantity            *string `sil:"F70"`
	ScalableItem            *string `sil:"F82"`
}
