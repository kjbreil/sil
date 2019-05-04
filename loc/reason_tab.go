package loc

// ReasonTab is the REASON_TAB definition
type ReasonTab struct {
	TargetIdentifier string  `sil:"F1000"`
	ExceptCode       *int    `sil:"F1511"`
	ReasonCode       *int    `sil:"F2772"`
	ReasonDescriptor *string `sil:"F2773"`
}
