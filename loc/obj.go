package loc

// OBJ is the main object table
type OBJ struct {
	F01   string  `sil:"CHAR(13),zeropad"`
	F902  *string `sil:"CHAR(8)"`
	F07   *int    `sil:"INTEGER"`
	F11   *int    `sil:"INTEGER"`
	F12   *string `sil:"NUMBER(5,3)"`
	F13   *string `sil:"NUMBER(5,3)"`
	F14   *string `sil:"NUMBER(5,3)"`
	F16   *int    `sil:"INTEGER"`
	F17   *int    `sil:"INTEGER"`
	F18   *int    `sil:"INTEGER"`
	F21   *string `sil:"NUMBER(4,0)"`
	F22   *string `sil:"CHAR(30)"`
	F23   *string `sil:"CHAR(10)"`
	F29   *string `sil:"CHAR(60)"`
	F93   *string `sil:"CHAR(16)"`
	F155  *string `sil:"CHAR(30)"`
	F180  *string `sil:"CHAR(20)"`
	F193  *string `sil:"CHAR(16)"`
	F213  *string `sil:"CHAR(13)"`
	F214  *string `sil:"CHAR(13)"`
	F215  *int    `sil:"INTEGER"`
	F218  *string `sil:"CHAR(4)"`
	F255  *string `sil:"CHAR(120)"`
	F270  *string `sil:"NUMBER(5,3)"`
	F940  *int    `sil:"INTEGER"`
	F941  *int    `sil:"INTEGER"`
	F1000 *string `sil:"CHAR(5)"`
	F1001 int     `sil:"INTEGER" default:"1"`
	F1002 *string `sil:"NUMBER(5,3)"`
	F1004 *int    `sil:"INTEGER"`
	F1118 *string `sil:"CHAR(9)"`
	F1119 *string `sil:"CHAR(30)"`
	F1168 *int    `sil:"INTEGER"`
	F1699 *string `sil:"NUMBER(5,0)"`
	F1736 *int    `sil:"INTEGER"`
	F1737 *string `sil:"CHAR(13)"`
	F1738 *string `sil:"CHAR(13)"`
	F1744 *string `sil:"NUMBER(8,4)"`
	F1781 *string `sil:"CHAR(12)"`
	F1782 *string `sil:"CHAR(12)"`
	F1783 *int    `sil:"INTEGER"`
	F1784 *int    `sil:"INTEGER"`
	F1939 *string `sil:"CHAR(30)"`
	F1940 *string `sil:"CHAR(60)"`
	F1941 *string `sil:"CHAR(30)"`
	F1942 *string `sil:"CHAR(120)"`
	F1957 *string `sil:"CHAR(100)"`
	F1958 *string `sil:"CHAR(20)"`
	F1959 *string `sil:"CHAR(20)"`
	F1960 *string `sil:"CHAR(10)"`
	F1962 *string `sil:"CHAR(600)"`
	F1964 *string `sil:"CHAR(4)"`
	F2600 *string `sil:"CHAR(2)"`
	F2693 *string `sil:"CHAR(13)"`
	F2789 *string `sil:"CHAR(13)"`
	F2931 *int    `sil:"INTEGER"`
}
