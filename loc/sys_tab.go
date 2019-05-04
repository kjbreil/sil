package loc

// SysTab is the SYS_TAB definition
type SysTab struct {
	TransactionNumber    *int    `sil:"F1032"`
	DeployNumber         *string `sil:"F1059"`
	ImportExportNumber   *string `sil:"F1066"`
	LicenseControlCode   *string `sil:"F1243"`
	LicenseControlTotals *string `sil:"F1244"`
	SaleOrderNumber      *string `sil:"F1763"`
	InvoiceNumber        *string `sil:"F1764"`
	SpareCounter1        *int    `sil:"F1807"`
	SpareCounter2        *int    `sil:"F1808"`
	LastChangeDate       string  `sil:"F253" default:"NOW"`
	DateEnd              *string `sil:"F254"`
	DateChangeMinimum    *string `sil:"F2589"`
	DateChangeMaximum    *string `sil:"F2590"`
	LicenseControlString *string `sil:"F2874"`
	CouponOfferCode      *string `sil:"F302"`
	BatchIDentifier      *string `sil:"F902"`
	PurchaseOrderNumber  *string `sil:"F91"`
}
