package loc

// LnkTab is the LNK_TAB definition
type LnkTab struct {
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	TerminalStore    string  `sil:"F1056"`
	TerminalNumber   *string `sil:"F1057"`
}
