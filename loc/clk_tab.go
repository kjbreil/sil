package loc

// ClkTab is the CLK_TAB definition
type ClkTab struct {
	TargetIDentifier      string  `sil:"F1000" default:"PAL"`
	RecordStatus          int     `sil:"F1001" default:"1"`
	TerminalStore         string  `sil:"F1056"`
	UserNumber            *int    `sil:"F1126"`
	UserShortName         *string `sil:"F1127"`
	UserSecretNumber      *string `sil:"F1141"`
	UserRestrictionLevel  *int    `sil:"F1142"`
	UserFirstName         *string `sil:"F1143"`
	UserLastName          *string `sil:"F1144"`
	UserBirthDate         *string `sil:"F1145"`
	UserLanguage          *string `sil:"F1146"`
	CustomerID            *string `sil:"F1148"`
	UserActiveOnNode      *string `sil:"F1176"`
	UserOrderNumber       *int    `sil:"F1185"`
	DateCreation          string  `sil:"F1264" default:"NOW"`
	UserStatus            *string `sil:"F1552"`
	UserHiringDate        *string `sil:"F1553"`
	UserBreakDate         *string `sil:"F1554"`
	UserLastPromotionDate *string `sil:"F1555"`
	UserSuperior          *int    `sil:"F1556"`
	UserAddress1          *string `sil:"F1557"`
	UserAddress2          *string `sil:"F1558"`
	UserCity              *string `sil:"F1559"`
	UserProvinceState     *string `sil:"F1560"`
	UserCountry           *string `sil:"F1561"`
	UserZipPostalCode     *string `sil:"F1562"`
	UserTelephone1        *string `sil:"F1563"`
	UserTelephone2        *string `sil:"F1564"`
	UserInsuranceNumber   *string `sil:"F1565"`
	UserHolidays          *string `sil:"F1566"`
	UserPayrollTableLink  *int    `sil:"F1567"`
	UserAssignDrawerNum   *int    `sil:"F1568"`
	UserDepartmentNumber  *int    `sil:"F1569"`
	UserAddedExperience   *string `sil:"F1570"`
	UserEmailAddress      *string `sil:"F1571"`
	UserJobDescription    *string `sil:"F1585"`
	UserGender            *string `sil:"F1586"`
	UserBreakReason       *string `sil:"F1587"`
	UserCivilState        *string `sil:"F1588"`
	UserEmployeeNumber    *string `sil:"F1589"`
	UserNote              *string `sil:"F1590"`
	RestrictionCode       *int    `sil:"F170"`
	StoreResponsible      *string `sil:"F1964"`
	LastChangeDate        string  `sil:"F253" default:"NOW"`
	ClockInout            *string `sil:"F2587"`
	Profil                *string `sil:"F2597"`
	ShowFilterUser        *string `sil:"F2692"`
	VendorID              *string `sil:"F27"`
	CellularPhone         *string `sil:"F2806"`
	UserGroup             *string `sil:"F2827"`
	WeeklyHomeBreakfast   *string `sil:"F2828"`
	WeeklyHomeLunch       *string `sil:"F2829"`
	WeeklyHomeDinners     *string `sil:"F2830"`
	WeeklyHomeSnacks      *string `sil:"F2831"`
	PetsAtHome            *string `sil:"F2832"`
	ActivityLevel         *int    `sil:"F2833"`
	NotificationOptOut    *string `sil:"F2844"`
	BatchIDentifier       *string `sil:"F902"`
	CreatedByUser         *int    `sil:"F940"`
	ModifiedByUser        *int    `sil:"F941"`
}
