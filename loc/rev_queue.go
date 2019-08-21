package loc

// RevQueue is the REV_QUEUE definition
type RevQueue struct {
	RevisionInterfaceID      *string `sil:"F3000"`
	RevisionSequence         *string `sil:"F3001"`
	RevisionTriggerTheChange *string `sil:"F3002"`
	RevisionType             *string `sil:"F3003"`
	RevisionUsername         *string `sil:"F3004"`
	RevisionChangeDate       *string `sil:"F3005"`
	RevisionSequenceSource   *string `sil:"F3006"`
	RevisionChangeDateLocal  *string `sil:"F3007"`
	RevisionTableName        *string `sil:"F3009"`
}
