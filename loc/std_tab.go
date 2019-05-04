package loc

// StdTab is the STD_TAB definition
type StdTab struct {
	TerminalStore       string  `sil:"F1056"`
	StoreLongNumber     *string `sil:"F1530"`
	StoreName           *string `sil:"F1531"`
	StoreAddress1       *string `sil:"F1532"`
	StoreAddress2       *string `sil:"F1533"`
	StoreCity           *string `sil:"F1534"`
	StoreZipCode        *string `sil:"F1535"`
	StorePhone          *string `sil:"F1536"`
	StoreFax            *string `sil:"F1537"`
	StoreEmail          *string `sil:"F1538"`
	StoreClass          *string `sil:"F1539"`
	StoreGroup          *string `sil:"F1540"`
	StoreLocationType   *string `sil:"F1541"`
	StoreRegion         *string `sil:"F1542"`
	StoreManager        *int    `sil:"F1543"`
	StoreInterim        *int    `sil:"F1544"`
	StoreManhourWeek1   *string `sil:"F1545"`
	StoreOpeningDate    *string `sil:"F1546"`
	StoreClosingDate    *string `sil:"F1547"`
	StoreType           *string `sil:"F1548"`
	StoreProvince       *string `sil:"F1549"`
	StoreElible         *string `sil:"F1551"`
	StoreManhourWeek2   *string `sil:"F1579"`
	StoreManhourWeek3   *string `sil:"F1580"`
	StoreLongNumber2    *string `sil:"F2688"`
	AccountingPrefix    *string `sil:"F2689"`
	AccountingPrefix2   *string `sil:"F2690"`
	BackstoreCapacity   *string `sil:"F2698"`
	BankAccount         *string `sil:"F2774"`
	BankAccount2        *string `sil:"F2775"`
	CameraServerAddress *string `sil:"F2776"`
	CameraTerminalLink  *string `sil:"F2777"`
	MStoreServer        *string `sil:"F2840"`
	EStoreServer        *string `sil:"F2841"`
	StoreLongitude      *string `sil:"F2842"`
	StoreLatitude       *string `sil:"F2843"`
	BankName1           *string `sil:"F2849"`
	BankName2           *string `sil:"F2850"`
	StoreServices       *string `sil:"F2937"`
}
