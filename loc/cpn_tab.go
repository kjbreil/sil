package loc

// CpnTab is the CPN_TAB definition
type CpnTab struct {
	TargetIDentifier   string  `sil:"F1000" default:"PAL"`
	RecordStatus       int     `sil:"F1001" default:"1"`
	CouponCode         *int    `sil:"F2034"`
	CouponData         *string `sil:"F2035"`
	CouponSignificance *string `sil:"F2036"`
	CouponMinQtyBuy    *string `sil:"F2037"`
	CouponPrice        *string `sil:"F2038"`
	CouponType         *string `sil:"F2039"`
}
