package loc

// ClrTab is the CLR_TAB definition
type ClrTab struct {
	CustRiskLevel             *int    `sil:"F1153"`
	CustPrintStatement        *string `sil:"F1162"`
	CustRiskMaxCredit         *string `sil:"F1165"`
	CustRiskInterestRate      *string `sil:"F1166"`
	CustRiskTerm              *string `sil:"F1167"`
	CustRiskDescriptor        *string `sil:"F1175"`
	CheckLimit                *string `sil:"F1257"`
	CheckChangeLimit          *string `sil:"F1258"`
	CheckWeekLimitAmt         *string `sil:"F1259"`
	AcceptChecks              *string `sil:"F1260"`
	InvoiceTemplateName       *string `sil:"F1762"`
	ReceiptCopyCount          *int    `sil:"F1788"`
	CheckWeekLimitQty         *int    `sil:"F1876"`
	PrivateCardAccept         *string `sil:"F1877"`
	PrivateCardLimit          *string `sil:"F1878"`
	PrivateCardWeeklyLimit    *string `sil:"F1879"`
	PrivateCardWeeklyLimitQty *int    `sil:"F1880"`
	InvoiceCopyCount          *int    `sil:"F1881"`
	DaysForPayment            *int    `sil:"F1921"`
	EarlyPaymentDays          *int    `sil:"F1922"`
	EarlyPaymentDiscount      *string `sil:"F1923"`
	LastChangeDate            string  `sil:"F253" default:"NOW"`
}
