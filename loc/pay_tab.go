package loc

// PayTab is the PAY_TAB definition
type PayTab struct {
	DepartmentCode       *int    `sil:"F03"`
	SubDepartmentCode    *int    `sil:"F04"`
	UserPayrollTableLink *int    `sil:"F1567"`
	JobDescription       *string `sil:"F2648"`
	Salary               *string `sil:"F2649"`
	Overtime1            *string `sil:"F2650"`
	Overtime2            *string `sil:"F2651"`
	RealCost             *string `sil:"F2652"`
	Extra1               *string `sil:"F2779"`
	Extra2               *string `sil:"F2780"`
	HourLimit            *int    `sil:"F2781"`
}
