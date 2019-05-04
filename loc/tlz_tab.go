package loc

// TlzTab is the TLZ_TAB definition
type TlzTab struct {
	TotalizerNumber          int     `sil:"F1034"`
	TotalizerDescriptor      *string `sil:"F1039"`
	NewUserAccessLevel       *int    `sil:"F1042"`
	OperUserAccessLevel      *int    `sil:"F1043"`
	AssistantUserAccessLevel *int    `sil:"F1044"`
	ManagerUserAccessLevel   *int    `sil:"F1045"`
	OwnerUserAccessLevel     *int    `sil:"F1046"`
	ProgrammerUserAccess     *int    `sil:"F1047"`
	AffectedFiles            *string `sil:"F1048"`
	NoUserAccessLevel        *int    `sil:"F1054"`
	ClerkUserAccessLevel     *int    `sil:"F1055"`
	TtlzGroup1               *string `sil:"F1128"`
	TtlzGroup2               *string `sil:"F1129"`
	TtlzGLAccount            *string `sil:"F1130"`
	TtlzGroupingExpression   *string `sil:"F1131"`
	SequenceNumber           *int    `sil:"F1147"`
	TotalizerSectionNumber   *int    `sil:"F1179"`
	DebitOrCreditEntry       *string `sil:"F1253"`
	AffectFilesDetails       *string `sil:"F1709"`
	AddingTriggers           *string `sil:"F1710"`
	OperatorAccessLevel      *int    `sil:"F1817"`
	GeneralManagerAccessLvl  *int    `sil:"F1818"`
	AltDescTlz               *string `sil:"F1896"`
	ShowInEJ                 *string `sil:"F1897"`
	ShowPriority             *int    `sil:"F1965"`
	ShowFilter               *string `sil:"F1966"`
}
