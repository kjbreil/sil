package loc

// TerTab is the TER_TAB definition
type TerTab struct {
	TerminalStore           string  `sil:"F1056"`
	TerminalNumber          *string `sil:"F1057"`
	TerminalDescription     *string `sil:"F1058"`
	TerminalSendingDriver   *string `sil:"F1125"`
	TerminalReceivingDriver *string `sil:"F1169"`
}
