package loc

// ResTab is the RES_TAB definition
type ResTab struct {
	RecordStatus            int     `sil:"F1001" default:"1"`
	RestrictionCode         *int    `sil:"F170"`
	MondayStartTime         *string `sil:"F1745"`
	MondayStopTime          *string `sil:"F1746"`
	TuesdayStartTime        *string `sil:"F1747"`
	TuesdayStopTime         *string `sil:"F1748"`
	WednesdayStartTime      *string `sil:"F1749"`
	WednesdayStopTime       *string `sil:"F1750"`
	ThursdayStartTime       *string `sil:"F1751"`
	ThursdayStopTime        *string `sil:"F1752"`
	FridayStartTime         *string `sil:"F1753"`
	FridayStopTime          *string `sil:"F1754"`
	SaturdayStartTime       *string `sil:"F1755"`
	SaturdayStopTime        *string `sil:"F1756"`
	SundayStartTime         *string `sil:"F1757"`
	SundayStopTime          *string `sil:"F1758"`
	RestrictionDescriptor   *string `sil:"F1830"`
	RestrictionOrPermission *string `sil:"F1841"`
}
