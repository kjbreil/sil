package loc

// HeaderDct is the HEADER_DCT definition
type HeaderDct struct {
	BatchType              *string `sil:"F901"`
	BatchIdentifier        *string `sil:"F902"`
	BatchCreator           *string `sil:"F903"`
	BatchDestination       *string `sil:"F904"`
	BatchAuditFile         *string `sil:"F905"`
	BatchResponseFile      *string `sil:"F906"`
	BatchEndingDate        *string `sil:"F907"`
	BatchEndingTime        *string `sil:"F908"`
	BatchActiveDate        *string `sil:"F909"`
	BatchActiveTime        *string `sil:"F910"`
	BatchPurgeDate         *string `sil:"F911"`
	BatchActionType        *string `sil:"F912"`
	BatchDescription       *string `sil:"F913"`
	BatchUser1state        *string `sil:"F914"`
	BatchUser2Count        *string `sil:"F915"`
	BatchUser3Macro        *string `sil:"F916"`
	BatchWarningLevel      *int    `sil:"F917"`
	BatchMaximumErrorCount *int    `sil:"F918"`
	BatchFileVersion       *string `sil:"F919"`
	BatchCreatorVersion    *string `sil:"F920"`
	BatchPrimaryKey        *string `sil:"F921"`
	BatchSpecificCommand   *string `sil:"F922"`
	ShelfTagType           *string `sil:"F930"`
	BatchExecutionPriority *int    `sil:"F931"`
	BatchLongDescription   *string `sil:"F932"`
}
