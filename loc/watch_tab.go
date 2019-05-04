package loc

// WatchTab is the WATCH_TAB definition
type WatchTab struct {
	TargetIDentifier    string  `sil:"F1000" default:"PAL"`
	MovementPeriodType  string  `sil:"F1031"`
	TotalizerNumber     int     `sil:"F1034"`
	LineNumber          *int    `sil:"F1101"`
	UserOrderNumber     *int    `sil:"F1185"`
	TaskType            *string `sil:"F1823"`
	CubeFilename        *string `sil:"F2734"`
	CubeMnemonic        *string `sil:"F2735"`
	CubeCompareField    *string `sil:"F2736"`
	CubeCompareOperator *string `sil:"F2737"`
	CubeCompareValue    *string `sil:"F2738"`
	CubeCompareType     *string `sil:"F2739"`
	CubeKeyType         *string `sil:"F2740"`
	CubeKeyValue        *string `sil:"F2741"`
	CubePeriodDetail    *string `sil:"F2742"`
}
