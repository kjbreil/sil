package loc

// SclStoTab is the SCL_STO_TAB definition
type SclStoTab struct {
	TargetIdentifier           string  `sil:"F1000"`
	RecordStatus               int     `sil:"F1001" default:"1"`
	StorageInstructionNumber   *int    `sil:"F2952"`
	StorageInstructionText     *string `sil:"F2953"`
	StorageInstructionFontSize *int    `sil:"F2954"`
}
