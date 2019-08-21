package loc

// RevHdr is the REV_HDR definition
type RevHdr struct {
	RevisionSequence         *string `sil:"F3001"`
	RevisionTriggerTheChange *string `sil:"F3002"`
	RevisionType             *string `sil:"F3003"`
	RevisionUsername         *string `sil:"F3004"`
	RevisionChangeDate       *string `sil:"F3005"`
	RevisionChangeDateLocal  *string `sil:"F3007"`
	RevisionTableName        *string `sil:"F3009"`
}
