package loc

// VendorTab is the VENDOR_TAB definition
type VendorTab struct {
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	OperatorResponsible      *int    `sil:"F1168"`
	ShipIDNumber             *string `sil:"F1642"`
	FreightRateWeight        *string `sil:"F1654"`
	FreightRateVolume        *string `sil:"F1655"`
	VendorTermDescriptor     *string `sil:"F1656"`
	BackOrderDays            *int    `sil:"F1685"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	MinimumCases             *string `sil:"F1760"`
	MinimumCubic             *string `sil:"F1761"`
	VendorInternalComment    *string `sil:"F1779"`
	DeliveryDays             *int    `sil:"F1793"`
	DeliveryCycleDays        *int    `sil:"F1794"`
	VendorComment            *string `sil:"F1882"`
	MinimumWeight            *string `sil:"F1883"`
	MinimumDollars           *string `sil:"F1884"`
	VendorContact2           *string `sil:"F1889"`
	VendorContact2Phone      *string `sil:"F1890"`
	VendorContact2Fax        *string `sil:"F1891"`
	VendorCountry            *string `sil:"F1948"`
	VendorStatus             *string `sil:"F1949"`
	StoreResponsible         *string `sil:"F1964"`
	ShowPriority             *int    `sil:"F1965"`
	ShowFilter               *string `sil:"F1966"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	Profil                   *string `sil:"F2597"`
	CurrencyCode             *string `sil:"F2602"`
	VendorEmail              *string `sil:"F2603"`
	InvisibleOnPDA           *string `sil:"F2658"`
	URL                      *string `sil:"F2660"`
	DefaultQuantity          *string `sil:"F2666"`
	VendorID                 *string `sil:"F27"`
	MaximumCases             *string `sil:"F2700"`
	MaximumCubic             *string `sil:"F2701"`
	MaximumWeight            *string `sil:"F2702"`
	AssociatedTaxRates       *string `sil:"F2710"`
	VendorName               *string `sil:"F334"`
	VendorContactName        *string `sil:"F335"`
	VendorAddressLine1       *string `sil:"F336"`
	VendorAddressLine2       *string `sil:"F337"`
	VendorAddressCity        *string `sil:"F338"`
	VendorAddressState       *string `sil:"F339"`
	VendorAddressZip         *string `sil:"F340"`
	VendorPhoneVoice         *string `sil:"F341"`
	VendorPhoneFax           *string `sil:"F342"`
	VendorProductCategories  *string `sil:"F343"`
	FreightbrokerCompany     *string `sil:"F344"`
	FOBPoint                 *string `sil:"F345"`
	VendorLimit              *string `sil:"F346"`
	SpecialDiscountPercent   *string `sil:"F347"`
	TermsDays                *int    `sil:"F348"`
	TermsDiscountPercent     *string `sil:"F349"`
	DUNSNumber               *string `sil:"F350"`
	SupplierLocationNumber   *string `sil:"F351"`
	VendorAccount            *string `sil:"F352"`
	DeliverMondayFlag        *string `sil:"F353"`
	DeliverTuesdayFlag       *string `sil:"F354"`
	DeliverWednesday         *string `sil:"F355"`
	DeliverThursday          *string `sil:"F356"`
	DeliverFriday            *string `sil:"F357"`
	DeliverSaturday          *string `sil:"F358"`
	DeliverSunday            *string `sil:"F359"`
	BatchIDentifier          *string `sil:"F902"`
}
