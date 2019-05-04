package loc

// LnkTab is the LNK_TAB definition
type LnkTab struct {
	TargetIdentifier string  `sil:"F1000"`
	TerminalStore    string  `sil:"F1056"`
	TerminalNumber   *string `sil:"F1057"`
}
