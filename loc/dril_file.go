package loc

// DrilFile is the Dril_File definition
type DrilFile struct {
	TargetIDentifier string  `sil:"F1000" default:"PAL"`
	DrilKey          *string `sil:"F1597"`
	DrilFileTitle    *string `sil:"F1606"`
	DrilFileStyle    *string `sil:"F1607"`
	DrilFileName     *string `sil:"F1608"`
}
