package loc

// ClfTab is the CLF_TAB definition
type ClfTab struct {
	LimitTransaction         *string `sil:"F1073"`
	LimitPeriod              *string `sil:"F1074"`
	LimitPeriodValue         *string `sil:"F1075"`
	LimitTotalizer           *int    `sil:"F1076"`
	CustShopperLevel         *int    `sil:"F1152"`
	ShopperPriceLevel        *int    `sil:"F1163"`
	ShopperDiscount          *string `sil:"F1164"`
	ShopperDescriptor        *string `sil:"F1174"`
	ShopperCouponMultiple    *string `sil:"F1269"`
	ShopperMaxCouponTotal    *string `sil:"F1270"`
	ShopperGetDscDollar      *string `sil:"F1271"`
	ShopperGetDscPct         *string `sil:"F1272"`
	ShopperGetLinkedItem     *string `sil:"F1273"`
	ShopperGetCouponLink     *string `sil:"F1274"`
	ShopperMaxCouponValue    *string `sil:"F1275"`
	ShopperOperatorLevel     *int    `sil:"F1276"`
	ShopperGetDscDept        *string `sil:"F1277"`
	ShopperPointRatio        *string `sil:"F1278"`
	ShopperPointMethod       *string `sil:"F1279"`
	ShopperPointMultiplier   *string `sil:"F1280"`
	RedeemPointRatio         *string `sil:"F1281"`
	RedeemPointMethod        *string `sil:"F1282"`
	RedeemPointMultiplier    *string `sil:"F1283"`
	BestPriceMethod          *string `sil:"F1284"`
	RedeemPointMinBalance    *int    `sil:"F1285"`
	ShopperPointMinSale      *string `sil:"F1286"`
	ShopperGetPackagePrice   *string `sil:"F1287"`
	ShopperGetLimitedPrice   *string `sil:"F1288"`
	UpgradeTotalTransaction  *string `sil:"F1289"`
	UpgradePointBalance      *int    `sil:"F1290"`
	UpgradeSpending          *string `sil:"F1291"`
	UpgradeSpendingPeriod    *string `sil:"F1292"`
	UpgradeSpendingTotalizer *int    `sil:"F1293"`
	UpgradeToLevel           *int    `sil:"F1294"`
	MultiplePromotionMethod  *string `sil:"F1295"`
	ShopperPointExpirDays    *int    `sil:"F1296"`
	ShopperCouponAddValue    *string `sil:"F1297"`
	ShopperCpnDblMaxTrs      *string `sil:"F1298"`
	PriceLevelFallBack       *int    `sil:"F1711"`
	MondayStartTime          *string `sil:"F1745"`
	MondayStopTime           *string `sil:"F1746"`
	TuesdayStartTime         *string `sil:"F1747"`
	TuesdayStopTime          *string `sil:"F1748"`
	WednesdayStartTime       *string `sil:"F1749"`
	WednesdayStopTime        *string `sil:"F1750"`
	ThursdayStartTime        *string `sil:"F1751"`
	ThursdayStopTime         *string `sil:"F1752"`
	FridayStartTime          *string `sil:"F1753"`
	FridayStopTime           *string `sil:"F1754"`
	SaturdayStartTime        *string `sil:"F1755"`
	SaturdayStopTime         *string `sil:"F1756"`
	SundayStartTime          *string `sil:"F1757"`
	SundayStopTime           *string `sil:"F1758"`
	BasePriceLevel           *int    `sil:"F1811"`
	SpareCLF                 *string `sil:"F1812"`
	AutomaticRedeem          *string `sil:"F1886"`
	CostPlusPercent          *string `sil:"F1938"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	PointProgramMember       *string `sil:"F2743"`
}
