package loc

// ECL is Electronic Coupon Link Table
type ECL struct {
	F902  string `sil:"CHAR(8)"`
	F164  string `sil:"CHAR(13)"`
	F1033 string `sil:"CHAR(3)"`
	F01   string `sil:"CHAR(13)"`
	F1031 string `sil:"CHAR(1)"`
	F1034 string `sil:"INTEGER"`
	F1000 string `sil:"CHAR(5)"`
	F02   string `sil:"CHAR(40)"`
	F33   string `sil:"CHAR(8)"`
	F64   string `sil:"NUMBER(10,3)"`
	F65   string `sil:"MONEY"`
	F67   string `sil:"NUMBER(10,3)"`
	F1001 string `sil:"INTEGER"`
	F1759 string `sil:"INTEGER"`
	F1964 string `sil:"CHAR(4)"`
}