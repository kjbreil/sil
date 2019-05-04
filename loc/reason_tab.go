package loc

// ReasonTab is the REASON_TAB definition
type ReasonTab struct {
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	ExceptCode       *int    `sil:"F1511"`
	ReasonCode       *int    `sil:"F2772"`
	ReasonDescriptor *string `sil:"F2773"`
}
