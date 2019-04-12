package loc

// PRICE is the price table
// currently incomplete
type PRICE struct {
	F1000 string  `sil:"CHAR(5)" default:"PAL"`
	F1056 *string `sil:"CHAR(4)"`
	F01   string  `sil:"CHAR(13)"`
	F902  *string `sil:"CHAR(8)"`
	F126  *int    `sil:"INTEGER"`
	F07   *int    `sil:"INTEGER"`
	F30   *string `sil:"MONEY"`
	F31   *string `sil:"NUMBER(3,0)"`
	F32   *int    `sil:"INTEGER"`
	F33   *string `sil:"CHAR(8)"`
	F34   *string `sil:"CHAR(2)"`
	F35   *string `sil:"DATE(7)"`
	F36   *string `sil:"TIME(4)"`
	F37   *string `sil:"CHAR(1)"`
	F49   *string `sil:"NUMBER(5,3)"`
	F62   *string `sil:"NUMBER(2,0)"`
	F63   *string `sil:"MONEY"`
	F111  *string `sil:"MONEY"`
	F112  *string `sil:"NUMBER(5,3)"`
	F129  *string `sil:"DATE(7)"`
	F130  *string `sil:"TIME(4)"`
	F133  *string `sil:"DATE(7)"`
	F140  *string `sil:"MONEY"`
	F142  *string `sil:"NUMBER(3,0)"`
	F168  *string `sil:"MONEY"`
	F169  *int    `sil:"INTEGER"`
	F205  *string `sil:"MONEY"`
	F1001 *string `sil:"INTEGER"`
	F1005 *string `sil:"CHAR(4)"`
	F1713 *string `sil:"NUMBER(10,4)"`
	F1714 *string `sil:"CHAR(4)"`
	F1767 *string `sil:"NUMBER(10,4)"`
	F1768 *string `sil:"CHAR(4)"`
	F1769 *string `sil:"NUMBER(10,4)"`
	F1770 *string `sil:"CHAR(4)"`
	F1927 *string `sil:"NUMBER(8,4)"`
	F1934 *string `sil:"MONEY"`
	F1964 *string `sil:"CHAR(4)"`
}
