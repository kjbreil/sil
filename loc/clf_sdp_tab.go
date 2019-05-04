package loc

// ClfSdpTab is the CLF_SDP_TAB definition
type ClfSdpTab struct {
	SubDepartmentCode *int    `sil:"F04"`
	CustShopperLevel  *int    `sil:"F1152"`
	ShopperDiscount   *string `sil:"F1164"`
	CostPlusPercent   *string `sil:"F1938"`
}
