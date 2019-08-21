package loc

// RunTab is the RUN_TAB definition
type RunTab struct {
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	AutorunOrder             *int    `sil:"F1102"`
	AutorunCommand           *string `sil:"F1103"`
	AutorunDestinationStores *string `sil:"F1104"`
	AutorunLastExecDateTime  *string `sil:"F1105"`
	AutorunDateOffset        *int    `sil:"F1106"`
	AutorunNextExecDateTime  *string `sil:"F1107"`
	AutorunExecuteManually   *string `sil:"F1108"`
	AutorunDescriptor        *string `sil:"F1109"`
	AutorunExecTooLate       *string `sil:"F1110"`
	AutorunExecPastEvents    *string `sil:"F1111"`
	AutorunSpecificDate      *string `sil:"F1112"`
	AutorunIfFileExist       *string `sil:"F1113"`
	AutorunEveryNDays        *int    `sil:"F1114"`
	AutorunEveryNMonths      *int    `sil:"F1115"`
	AutorunEveryNPasses      *int    `sil:"F1116"`
	AutorunEveryNMinutes     *int    `sil:"F1117"`
}
