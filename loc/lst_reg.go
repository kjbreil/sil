package loc

// LstReg is the LST_REG definition
type LstReg struct {
	UPCCode                 string  `sil:"F01,zeropad"`
	DepartmentCode          *int    `sil:"F03"`
	SubDepartmentCode       *int    `sil:"F04"`
	RecordStatus            int     `sil:"F1001" default:"1"`
	DescriptionRegistration *string `sil:"F1041"`
	RegistrationMode        *string `sil:"F1067"`
	TentativeItemFlag       *string `sil:"F1071"`
	SelectPkgPrice          *int    `sil:"F1874"`
	LastChangeDate          string  `sil:"F253" default:"NOW"`
	ListID                  *string `sil:"F2891"`
	TotalUnits              *string `sil:"F64"`
	TotalWeight             *string `sil:"F67"`
	ModifiedByUser          *int    `sil:"F941"`
}
