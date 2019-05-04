package loc

// RevExec is the REV_EXEC definition
type RevExec struct {
	RevisionInterfaceID    *string `sil:"F3000"`
	RevisionSequence       *string `sil:"F3001"`
	RevisionChangeDate     *string `sil:"F3005"`
	RevisionSequenceSource *string `sil:"F3006"`
	RevisionTableName      *string `sil:"F3009"`
	RevisionPostedDate     *string `sil:"F3010"`
}
