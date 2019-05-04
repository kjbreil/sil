package loc

// SalBat is the SAL_BAT definition
type SalBat struct {
	TransactionNumber       *int    `sil:"F1032"`
	MovementStartTime       *string `sil:"F1035"`
	MovementEndTime         *string `sil:"F1036"`
	TerminalStore           string  `sil:"F1056"`
	TerminalNumber          *string `sil:"F1057"`
	RegistrationMode        *string `sil:"F1067"`
	TransactionMode         *string `sil:"F1068"`
	LineNumber              *int    `sil:"F1101"`
	UserNumber              *int    `sil:"F1126"`
	UserShortName           *string `sil:"F1127"`
	TransactionVoidCount    *int    `sil:"F1137"`
	CustomerId              *string `sil:"F1148"`
	CustFirstName           *string `sil:"F1149"`
	CustLastName            *string `sil:"F1150"`
	CustPosComment          *string `sil:"F1151"`
	CustShopperLevel        *int    `sil:"F1152"`
	CustRiskLevel           *int    `sil:"F1153"`
	CustGroup               *string `sil:"F1154"`
	CustCompanyName         *string `sil:"F1155"`
	CustAddress1            *string `sil:"F1156"`
	CustAddress2            *string `sil:"F1157"`
	CustCity                *string `sil:"F1158"`
	CustZippostalCode       *string `sil:"F1159"`
	CustTax13Exempt         *string `sil:"F1160"`
	CustTax24Exempt         *string `sil:"F1161"`
	ShopperPriceLevel       *int    `sil:"F1163"`
	ShopperDiscount         *string `sil:"F1164"`
	CustRiskMaxCredit       *string `sil:"F1165"`
	CustRiskTerm            *string `sil:"F1167"`
	OperatorResponsible     *int    `sil:"F1168"`
	CustProvincestate       *string `sil:"F1170"`
	CustCountry             *string `sil:"F1171"`
	CustPhoneNumber         *string `sil:"F1172"`
	CustFaxNumber           *string `sil:"F1173"`
	UserOrderNumber         *int    `sil:"F1185"`
	ClientPointsBeginning   *int    `sil:"F1238"`
	ClientBalanceBeginning  *string `sil:"F1242"`
	ReferenceNumber         *string `sil:"F1245"`
	DateDelivery            *string `sil:"F1246"`
	TransactionNote         *string `sil:"F1254"`
	TransBatchFileName      *string `sil:"F1255"`
	ShopperGetDsc           *string `sil:"F1271"`
	ShopperGetDsc           *string `sil:"F1272"`
	ShopperGetLinkedItem    *string `sil:"F1273"`
	ShopperGetCouponLink    *string `sil:"F1274"`
	ShopperGetDscDept       *string `sil:"F1277"`
	ShopperGetPackagePrice  *string `sil:"F1287"`
	ShopperGetLimitedPrice  *string `sil:"F1288"`
	MultiplePromotionMethod *string `sil:"F1295"`
	CustomerLanguage        *string `sil:"F1504"`
	BirthDate               *string `sil:"F1520"`
	EmailAddress            *string `sil:"F1573"`
	ShipIdNumber            *string `sil:"F1642"`
	ShipName                *string `sil:"F1643"`
	ShipContactName         *string `sil:"F1644"`
	ShipAddressLine1        *string `sil:"F1645"`
	ShipAddressLine2        *string `sil:"F1646"`
	ShipAddressCity         *string `sil:"F1647"`
	ShipAddressState        *string `sil:"F1648"`
	ShipAddressZip          *string `sil:"F1649"`
	ShipPhoneVoice          *string `sil:"F1650"`
	ShipPhoneFax            *string `sil:"F1651"`
	ShipProductCategories   *string `sil:"F1652"`
	DateShipping            *string `sil:"F1653"`
	FreightRateWeight       *string `sil:"F1654"`
	FreightRateVolume       *string `sil:"F1655"`
	BackOrderDays           *int    `sil:"F1685"`
	OperatorOrderPrinted    *int    `sil:"F1686"`
	OperatorReceiptPrinted  *int    `sil:"F1687"`
	OperatorInvoicePrinted  *int    `sil:"F1688"`
	OperatorExportMade      *int    `sil:"F1689"`
	ShippingNote            *string `sil:"F1692"`
	OperatorOrderEntered    *int    `sil:"F1693"`
	OperatorReceiptConfirm  *int    `sil:"F1694"`
	OperatorInvoiceEntered  *int    `sil:"F1695"`
	OperatorTrucker         *int    `sil:"F1696"`
	FreightRouteId          *string `sil:"F1697"`
	ShippingPieceCount      *string `sil:"F1699"`
	PriceLevelFallBack      *int    `sil:"F1711"`
	SaleOrderNumber         *string `sil:"F1763"`
	InvoiceNumber           *string `sil:"F1764"`
	CostPlusPercent         *string `sil:"F1938"`
	LastChangeDate          string  `sil:"F253" default:"NOW"`
	DateEnd                 *string `sil:"F254"`
	DatePayable             *string `sil:"F2596"`
	CurrencyExchangeNow     *string `sil:"F2598"`
	CurrencyExchangeOrder   *string `sil:"F2599"`
	SPAREHDRREGSAL          *int    `sil:"F2613"`
	SPAREHDRREGSAL          *string `sil:"F2614"`
	SPARE                   *string `sil:"F2615"`
	SPARE                   *int    `sil:"F2616"`
	SPARE                   *string `sil:"F2617"`
	SPARE                   *string `sil:"F2618"`
	FreightRate3            *string `sil:"F2619"`
	SPARE                   *string `sil:"F2620"`
	SPARE                   *int    `sil:"F2621"`
	SPARE                   *int    `sil:"F2622"`
	SPARE                   *string `sil:"F2623"`
	ShipInfoSource          *string `sil:"F2816"`
	TransactionComment      *string `sil:"F2848"`
	ShipCountry             *string `sil:"F2889"`
	RouteType               *string `sil:"F2904"`
	PickupNote              *string `sil:"F2934"`
	DateOrder               *string `sil:"F76"`
	BatchIdentifier         *string `sil:"F902"`
	PurchaseOrderNumber     *string `sil:"F91"`
}
