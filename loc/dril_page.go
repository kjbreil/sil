package loc

// DrilPage is the Dril_Page definition
type DrilPage struct {
	NewUserAccessLevel       *int    `sil:"F1042"`
	OperUserAccessLevel      *int    `sil:"F1043"`
	AssistantUserAccessLevel *int    `sil:"F1044"`
	ManagerUserAccessLevel   *int    `sil:"F1045"`
	OwnerUserAccessLevel     *int    `sil:"F1046"`
	ProgrammerUserAccess     *int    `sil:"F1047"`
	NoUserAccessLevel        *int    `sil:"F1054"`
	ClerkUserAccessLevel     *int    `sil:"F1055"`
	DrilKey                  *string `sil:"F1597"`
	DrilSequence             *int    `sil:"F1598"`
	DrilButtonStyle          *string `sil:"F1599"`
	DrilBgColor              *string `sil:"F1600"`
	DrilText                 *string `sil:"F1601"`
	DrilTextColor            *string `sil:"F1602"`
	DrilTextSize             *int    `sil:"F1603"`
	DrilTextAlign            *string `sil:"F1604"`
	DrilCommand              *string `sil:"F1605"`
	DrilDisplayFilter        *string `sil:"F1609"`
	DrillTargetFrame         *string `sil:"F1698"`
	OperatorAccessLevel      *int    `sil:"F1817"`
	GeneralManagerAccessLvl  *int    `sil:"F1818"`
}
