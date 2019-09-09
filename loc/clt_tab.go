package loc

// CltTab is the CLT_TAB definition
type CltTab struct {
	TargetIDentifier          string  `sil:"F1000" default:"PAL"`
	RecordStatus              int     `sil:"F1001" default:"1"`
	CustomerID                string  `sil:"F1148"`
	CustFirstName             *string `sil:"F1149"`
	CustLastName              *string `sil:"F1150"`
	CustPosComment            *string `sil:"F1151"`
	CustShopperLevel          *int    `sil:"F1152"`
	CustRiskLevel             *int    `sil:"F1153"`
	CustGroup                 *string `sil:"F1154"`
	CustCompanyName           *string `sil:"F1155"`
	CustAddress1              *string `sil:"F1156"`
	CustAddress2              *string `sil:"F1157"`
	CustCity                  *string `sil:"F1158"`
	CustZippostalCode         *string `sil:"F1159"`
	CustTax13Exempt           *string `sil:"F1160"`
	CustTax24Exempt           *string `sil:"F1161"`
	OperatorResponsible       *int    `sil:"F1168"`
	CustProvincestate         *string `sil:"F1170"`
	CustCountry               *string `sil:"F1171"`
	CustPhoneNumber           *string `sil:"F1172"`
	CustFaxNumber             *string `sil:"F1173"`
	DateCreation              string  `sil:"F1264" default:"NOW"`
	CustValidUntil            *string `sil:"F1265"`
	CustomerAlias             *string `sil:"F1501"`
	CustAliasReason           *string `sil:"F1502"`
	CustomerLanguage          *string `sil:"F1504"`
	BirthDate                 *string `sil:"F1520"`
	PasswordUpdateDate        *string `sil:"F1521"`
	MembershipRenewalDate     *string `sil:"F1522"`
	Gender                    *string `sil:"F1523"`
	CreatingStore             *string `sil:"F1524"`
	CallingStore              *string `sil:"F1525"`
	MembershipClosingDate     *string `sil:"F1526"`
	MailingContentCode        *string `sil:"F1527"`
	CallingReasonCode         *string `sil:"F1528"`
	VerifyAddress             *string `sil:"F1529"`
	CivicNumber               *string `sil:"F1550"`
	EmailAddress              *string `sil:"F1573"`
	CreditCardNumber          *string `sil:"F1574"`
	CreditCardExpiration      *string `sil:"F1575"`
	CreditCardSignature       *string `sil:"F1576"`
	SecretNumber              *string `sil:"F1581"`
	CreditCardType            *string `sil:"F1582"`
	LastCallingDate           *string `sil:"F1583"`
	LastMailingDate           *string `sil:"F1584"`
	ShipIDNumber              *string `sil:"F1642"`
	ShipName                  *string `sil:"F1643"`
	ShipContactName           *string `sil:"F1644"`
	ShipAddressLine1          *string `sil:"F1645"`
	ShipAddressLine2          *string `sil:"F1646"`
	ShipAddressCity           *string `sil:"F1647"`
	ShipAddressState          *string `sil:"F1648"`
	ShipAddressZip            *string `sil:"F1649"`
	ShipPhoneVoice            *string `sil:"F1650"`
	ShipPhoneFax              *string `sil:"F1651"`
	FreightRateWeight         *string `sil:"F1654"`
	FreightRateVolume         *string `sil:"F1655"`
	BackOrderDays             *int    `sil:"F1685"`
	ShippingNote              *string `sil:"F1692"`
	FreightRouteID            *string `sil:"F1697"`
	TemporaryShopperLevel     *int    `sil:"F1743"`
	MaintenanceOperatorLevel  *int    `sil:"F1759"`
	CustomerGroup2            *string `sil:"F1777"`
	CustomerInternalComment   *string `sil:"F1778"`
	SpareCLT                  *string `sil:"F1809"`
	SpecialAddingFunction     *int    `sil:"F1810"`
	CustomerStatus            *string `sil:"F1950"`
	Classification            *string `sil:"F1957"`
	StoreResponsible          *string `sil:"F1964"`
	LastChangeDate            string  `sil:"F253" default:"NOW"`
	Profil                    *string `sil:"F2597"`
	CurrencyCode              *string `sil:"F2602"`
	FreightRate3              *string `sil:"F2619"`
	URL                       *string `sil:"F2660"`
	PointProgramMember        *string `sil:"F2743"`
	LimitTransVolume          *string `sil:"F2807"`
	LimitTransAmount          *string `sil:"F2808"`
	LimitDailyVolume          *string `sil:"F2809"`
	LimitDailyAmount          *string `sil:"F2810"`
	LimitCumulVolume          *string `sil:"F2811"`
	LimitCumulAmount          *string `sil:"F2812"`
	CustCellPhone             *string `sil:"F2875"`
	ShipCountry               *string `sil:"F2889"`
	SubstitutionPreference    *string `sil:"F2899"`
	NotificationPreference    *string `sil:"F2900"`
	AcceptedTermsAndCondition *string `sil:"F2901"`
	RouteType                 *string `sil:"F2904"`
	PickupNote                *string `sil:"F2934"`
	BatchIDentifier           *string `sil:"F902"`
	CreatedByUser             *int    `sil:"F940"`
	ModifiedByUser            *int    `sil:"F941"`
}
