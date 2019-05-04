package loc

// OfrTab is the OFR_TAB definition
type OfrTab struct {
	VendorCoupon             *string `sil:"F104"`
	TerminalStore            string  `sil:"F1056"`
	TerminalNumber           *string `sil:"F1057"`
	LineNumber               *int    `sil:"F1101"`
	CustomerId               *string `sil:"F1148"`
	DateCreation             string  `sil:"F1264" default:"NOW"`
	CallingStore             *string `sil:"F1525"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	CouponShortDescription   *string `sil:"F2813"`
	CouponLongDescription    *string `sil:"F2814"`
	CouponManufacture        *string `sil:"F2815"`
	CouponCodeType           *string `sil:"F2817"`
	CouponStartDate          *string `sil:"F2818"`
	DateClipped              *string `sil:"F2819"`
	CouponState              *string `sil:"F2820"`
	CouponQtyAvailable       *string `sil:"F2821"`
	CouponQtyUsed            *string `sil:"F2822"`
	CouponExtendedCode       *string `sil:"F301"`
	CouponOfferCode          *string `sil:"F302"`
	CouponExpirationDate     *string `sil:"F303"`
	CouponHouseholdID        *string `sil:"F304"`
	CouponFaceValue          *string `sil:"F305"`
	CouponRedemptionMultiple *string `sil:"F306"`
	CouponRedemptionAmount   *string `sil:"F307"`
	CouponKeyEntered         *string `sil:"F308"`
	CouponOverrideReason     *string `sil:"F309"`
	CouponRejectReason       *string `sil:"F310"`
	CouponTransactionNumber  *int    `sil:"F311"`
	CouponTransactionDate    *string `sil:"F312"`
	CpnQualifyingItemCode1   *string `sil:"F313"`
	CpnQualifyingItemAmt1    *string `sil:"F314"`
	CpnQualifyingItemCode2   *string `sil:"F315"`
	CpnQualifyingItemAmt2    *string `sil:"F316"`
	CpnQualifyingItemCode3   *string `sil:"F317"`
	CpnQualifyingItemAmt3    *string `sil:"F318"`
	CpnQualifyingItemCode4   *string `sil:"F319"`
	CpnQualifyingItemAmt4    *string `sil:"F320"`
	CpnQualifyingItemCode5   *string `sil:"F321"`
	CpnQualifyingItemAmt5    *string `sil:"F322"`
	WeightCouponLinkFlag     *string `sil:"F385"`
	StoreCoupon              *string `sil:"F88"`
	CreatedByUser            *int    `sil:"F940"`
	ModifiedByUser           *int    `sil:"F941"`
}
