package loc

// RecReg is the REC_REG definition
type RecReg struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	DepartmentCode          *int    `sil:"F03"`
	SubDepartmentCode       *int    `sil:"F04"`
	BottleDepositLink       *int    `sil:"F05"`
	StatusCode              *string `sil:"F08"`
	TaxFlag6                *string `sil:"F100"`
	CaseQuantity            *string `sil:"F1003"`
	TaxFlag7                *string `sil:"F101"`
	TransactionNumber       *int    `sil:"F1032"`
	TotalizerNumber         int     `sil:"F1034"`
	DescriptionRegistration *string `sil:"F1041"`
	FunctionCode            *int    `sil:"F1063"`
	RegistrationMode        *string `sil:"F1067"`
	VoidPreviousFlag        *string `sil:"F1069"`
	RefundKeyFlag           *string `sil:"F1070"`
	TentativeItemFlag       *string `sil:"F1071"`
	ReferenceNumber         *string `sil:"F1079"`
	AlphaParameter          *string `sil:"F1081"`
	VoidProhibited          *string `sil:"F1086"`
	LineNumber              *int    `sil:"F1101"`
	CostFeeAmount           *string `sil:"F1121"`
	CostFeePercent          *string `sil:"F1122"`
	UnitNetCost             *string `sil:"F1140"`
	SplitCaseCost           *string `sil:"F120"`
	DateDelivery            *string `sil:"F1246"`
	ItemFreeFlag            *string `sil:"F1247"`
	ItemReturnFlag          *string `sil:"F1248"`
	AllowanceTotal          *string `sil:"F1249"`
	RebateTotal             *string `sil:"F1250"`
	DiscountTotal           *string `sil:"F1251"`
	CostBreakTotal          *string `sil:"F1252"`
	UnitOnOrder             *string `sil:"F1266"`
	WeightOnOrder           *string `sil:"F1267"`
	SequenceNumber          *int    `sil:"F131"`
	NextCostNet             *string `sil:"F151"`
	AllowanceQualifier      *string `sil:"F156"`
	EnvirfeesValue          *string `sil:"F1595"`
	EnvirfeesTlz            *int    `sil:"F1596"`
	AllowancePercent        *string `sil:"F1657"`
	DiscountAmount          *string `sil:"F1658"`
	DiscountOffInvoice      *string `sil:"F1659"`
	RebateOffInvoice        *string `sil:"F1660"`
	CostBreakOffInvoice     *string `sil:"F1661"`
	BillBack1Amount         *string `sil:"F1662"`
	BillBack1Percent        *string `sil:"F1663"`
	BillBack2Amount         *string `sil:"F1664"`
	BillBack2Percent        *string `sil:"F1665"`
	BillBack3Amount         *string `sil:"F1666"`
	BillBack3Percent        *string `sil:"F1667"`
	BillBack4Descriptor     *string `sil:"F1668"`
	BillBack4Amount         *string `sil:"F1669"`
	BillBack4Percent        *string `sil:"F1670"`
	BillBack1Total          *string `sil:"F1671"`
	BillBack2Total          *string `sil:"F1672"`
	BillBack3Total          *string `sil:"F1673"`
	BillBack4Total          *string `sil:"F1674"`
	EventRegularType        *string `sil:"F1675"`
	EventRegularDiscount    *string `sil:"F1676"`
	EventPromoType          *string `sil:"F1677"`
	EventPromoDiscount      *string `sil:"F1678"`
	EventRegularDays        *int    `sil:"F1679"`
	EventPromoDate          *string `sil:"F1680"`
	BottleDepositTotal      *string `sil:"F1681"`
	EnvirfeesTotal          *string `sil:"F1682"`
	FreightWeightTotal      *string `sil:"F1683"`
	FreightCubicTotal       *string `sil:"F1684"`
	BackOrderDays           *int    `sil:"F1685"`
	OperatorReceiptPrinted  *int    `sil:"F1687"`
	RegistrationNote        *string `sil:"F1691"`
	BackOrderFilled         *string `sil:"F1732"`
	RetailTotal             *string `sil:"F1771"`
	CostFeeAmountTotal      *string `sil:"F1772"`
	CostFeePercentTotal     *string `sil:"F1773"`
	TotalMatchUpdown        *string `sil:"F1774"`
	RetailPrice             *string `sil:"F1775"`
	SpareRECREG             *int    `sil:"F1776"`
	CostMethod              *string `sil:"F1791"`
	DeliveryDays            *int    `sil:"F1793"`
	SplitItemQty            *string `sil:"F1795"`
	SpareRECREGGCOST        *string `sil:"F1797"`
	OpenAmount              *string `sil:"F1815"`
	OpenTotalizer           *int    `sil:"F1816"`
	CaseDepositValue        *string `sil:"F186"`
	SOMatching              *string `sil:"F1899"`
	CaseSize                *string `sil:"F19"`
	BackOrderCase           *string `sil:"F1918"`
	BackOrderUnit           *string `sil:"F1919"`
	BackOrderWeight         *string `sil:"F1920"`
	DiscountQualifier       *string `sil:"F1973"`
	RebateQualifier         *string `sil:"F1974"`
	CostBreakQualifier      *string `sil:"F1975"`
	CostComment             *string `sil:"F1976"`
	AllowanceMinQty         *string `sil:"F1977"`
	DiscountMinQty          *string `sil:"F1978"`
	RebateMinQty            *string `sil:"F1979"`
	ReceivingPackSize       *string `sil:"F20"`
	AllowanceAmount         *string `sil:"F201"`
	AllowanceStartDate      *string `sil:"F202"`
	AllowanceEndDate        *string `sil:"F203"`
	AllowanceCode           *string `sil:"F204"`
	CostBreakEndDate        *string `sil:"F219"`
	SplitItemCode           *string `sil:"F220"`
	AllowanceNumber         *string `sil:"F222"`
	AllowanceOffInvoice     *string `sil:"F223"`
	CostBreakMinQty         *string `sil:"F224"`
	CostBreakAmount         *string `sil:"F225"`
	CostBreakPercent        *string `sil:"F226"`
	CostBreakStartDate      *string `sil:"F227"`
	DiscountPercent         *string `sil:"F228"`
	DiscountStartDate       *string `sil:"F229"`
	DiscountEndDate         *string `sil:"F230"`
	DiscountNumber          *string `sil:"F231"`
	RebateAmount            *string `sil:"F233"`
	RebateStartDate         *string `sil:"F234"`
	RebateEndDate           *string `sil:"F235"`
	RebateNumber            *string `sil:"F236"`
	RebatePercent           *string `sil:"F237"`
	LastChangeDate          string  `sil:"F253" default:"NOW"`
	DateEnd                 *string `sil:"F254"`
	Fee1                    *string `sil:"F2551"`
	Fee2                    *string `sil:"F2552"`
	Fee3                    *string `sil:"F2553"`
	Fee4                    *string `sil:"F2554"`
	Fee5                    *string `sil:"F2555"`
	Fee6                    *string `sil:"F2556"`
	Fee7                    *string `sil:"F2557"`
	Fee8                    *string `sil:"F2558"`
	Fee9                    *string `sil:"F2559"`
	Fee10                   *string `sil:"F2560"`
	Fee11                   *string `sil:"F2561"`
	Fee12                   *string `sil:"F2562"`
	Fee13                   *string `sil:"F2563"`
	Fee14                   *string `sil:"F2564"`
	Fee15                   *string `sil:"F2565"`
	DiscountMaxQty          *string `sil:"F2566"`
	RebateMaxQty            *string `sil:"F2567"`
	CostBreakMaxQty         *string `sil:"F2568"`
	VendorCode              *string `sil:"F26"`
	SPARERECREGGCOST        *int    `sil:"F2624"`
	SPARERECREG             *int    `sil:"F2625"`
	GrossPurchase           *string `sil:"F2627"`
	BackstorePickup         *string `sil:"F2629"`
	URL                     *string `sil:"F2660"`
	DefaultQuantity         *string `sil:"F2666"`
	WeightNet               *string `sil:"F270"`
	TotalTaxIncluded        *string `sil:"F2709"`
	Tax1Paid                *string `sil:"F2711"`
	Tax2Paid                *string `sil:"F2712"`
	Tax3Paid                *string `sil:"F2713"`
	Tax4Paid                *string `sil:"F2714"`
	Tax5Paid                *string `sil:"F2715"`
	Tax6Paid                *string `sil:"F2716"`
	Tax7Paid                *string `sil:"F2717"`
	Tax8Paid                *string `sil:"F2718"`
	FreightCaseWeight       *string `sil:"F327"`
	FreightCaseCube         *string `sil:"F328"`
	AllowanceMaxQty         *string `sil:"F329"`
	BaseCost                *string `sil:"F38"`
	BottleDepositValue      *string `sil:"F50"`
	TotalUnits              *string `sil:"F64"`
	TotalDollars            *string `sil:"F65"`
	TotalWeight             *string `sil:"F67"`
	UnitQuantity            *string `sil:"F70"`
	CaseOnOrder             *string `sil:"F75"`
	DateOrder               *string `sil:"F76"`
	TaxFlag1                *string `sil:"F81"`
	ScalableItem            *string `sil:"F82"`
	VendorAuthorizedItem    *string `sil:"F90"`
	DepositContainerCode    *string `sil:"F92"`
	TaxFlag2                *string `sil:"F96"`
	TaxFlag3                *string `sil:"F97"`
	TaxFlag4                *string `sil:"F98"`
	TaxFlag5                *string `sil:"F99"`
}