package loc

// RpcTab is the RPC_TAB definition
type RpcTab struct {
	ReportCodeDescriptor *string `sil:"F1024"`
	KnownShrinkFactor    *string `sil:"F1123"`
	CommissionRate       *string `sil:"F1124"`
	OperatorResponsible  *int    `sil:"F1168"`
	ReportCode           *int    `sil:"F18"`
	ShowPriority         *int    `sil:"F1965"`
	ShowFilter           *string `sil:"F1966"`
	SPARE                *string `sil:"F1967"`
	PriceMargin          *string `sil:"F49"`
}
