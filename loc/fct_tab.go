package loc

// FctTab is the FCT_TAB definition
type FctTab struct {
	TaxFlag6                 *string `sil:"F100"`
	TargetIdentifier         string  `sil:"F1000"`
	TaxFlag7                 *string `sil:"F101"`
	ProhibitQuantity         *string `sil:"F102"`
	NewUserAccessLevel       *int    `sil:"F1042"`
	OperUserAccessLevel      *int    `sil:"F1043"`
	AssistantUserAccessLevel *int    `sil:"F1044"`
	ManagerUserAccessLevel   *int    `sil:"F1045"`
	OwnerUserAccessLevel     *int    `sil:"F1046"`
	ProgrammerUserAccess     *int    `sil:"F1047"`
	MenuShortCut             *string `sil:"F1050"`
	MenuApplication          *string `sil:"F1051"`
	MenuTitle                *string `sil:"F1052"`
	MenuVisible              *string `sil:"F1053"`
	NoUserAccessLevel        *int    `sil:"F1054"`
	ClerkUserAccessLevel     *int    `sil:"F1055"`
	FunctionCode             *int    `sil:"F1063"`
	FunctionDescription      *string `sil:"F1064"`
	AlphaParameter           *string `sil:"F1081"`
	NoOverPayment            *string `sil:"F1082"`
	NoPartialPayment         *string `sil:"F1083"`
	NoDirectPayment          *string `sil:"F1084"`
	Rounding                 *string `sil:"F1085"`
	VoidProhibited           *string `sil:"F1086"`
	AccountNumberRequired    *string `sil:"F1088"`
	AlphaEntryRequired       *string `sil:"F1089"`
	ReasonCodeRequired       *string `sil:"F1090"`
	ManagerKeyRequired       *string `sil:"F1091"`
	ManagerPasswordRequired  *string `sil:"F1092"`
	SequenceNumber           *int    `sil:"F1147"`
	ProhibitRepeatKey        *string `sil:"F125"`
	ProhibitRefund           *string `sil:"F172"`
	OperatorAccessLevel      *int    `sil:"F1817"`
	GeneralManagerAccessLvl  *int    `sil:"F1818"`
	AltDescFct               *string `sil:"F1895"`
	ShowInEJ                 *string `sil:"F1897"`
	ShowPriority             *int    `sil:"F1965"`
	ShowFilter               *string `sil:"F1966"`
	MaximumAmount            *string `sil:"F239"`
	MinimumAmount            *string `sil:"F240"`
	MaximumVoid              *string `sil:"F241"`
	MaximumRefund            *string `sil:"F242"`
	TaxFlag1                 *string `sil:"F81"`
	RequireQuantity          *string `sil:"F85"`
	TaxFlag2                 *string `sil:"F96"`
	TaxFlag3                 *string `sil:"F97"`
	TaxFlag4                 *string `sil:"F98"`
	TaxFlag5                 *string `sil:"F99"`
}
