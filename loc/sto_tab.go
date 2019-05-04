package loc

// StoTab is the STO_TAB definition
type StoTab struct {
	TargetIDentifier      string  `sil:"F1000" default:"PAL"`
	TargetDescriptor      *string `sil:"F1018"`
	TargetForReportFlag   *string `sil:"F1180"`
	TargetForExchangeFlag *string `sil:"F1181"`
	TargetForProgramFlag  *string `sil:"F1182"`
	TargetPriority        *int    `sil:"F1937"`
	ShowPriority          *int    `sil:"F1965"`
	ShowFilter            *string `sil:"F1966"`
	ShowFiltersOperator   *string `sil:"F2691"`
}
