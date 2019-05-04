package loc

// RouteTab is the ROUTE_TAB definition
type RouteTab struct {
	TerminalStore     string  `sil:"F1056"`
	FreightRateWeight *string `sil:"F1654"`
	FreightRateVolume *string `sil:"F1655"`
	FreightRouteId    *string `sil:"F1697"`
	FreightRate3      *string `sil:"F2619"`
	PickupDate        *string `sil:"F2893"`
	PickupTime        *string `sil:"F2894"`
	OrderCountLimit   *int    `sil:"F2895"`
	OrderSizeLimit    *int    `sil:"F2896"`
	OrderCountTotal   *int    `sil:"F2897"`
	OrderSizeTotal    *int    `sil:"F2898"`
	OrderBefore       *string `sil:"F2902"`
	PickupOverdue     *string `sil:"F2903"`
	RouteType         *string `sil:"F2904"`
	RouteDescription  *string `sil:"F2905"`
}
