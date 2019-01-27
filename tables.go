package sil

func (s *SIL) TableCLK() {
	s.Table.Name = "CLK"
	s.Table.Columns = []Column{
		{
			Name: "F1185 INTEGER",
			Type: "F1001 INTEGER",
		},
		{
			Name: "F1126 INTEGER",
			Type: "F1571 CHAR(60)",
		},
		{
			Name: "F27 CHAR(14)",
			Type: "F170 INTEGER",
		},
		{
			Name: "F253 DATE(7)",
			Type: "F902 CHAR(8)",
		},
		{
			Name: "F940 INTEGER",
			Type: "F941 INTEGER",
		},
		{
			Name: "F1000 CHAR(5)",
			Type: "F1056 CHAR(4)",
		},
		{
			Name: "F1127 CHAR(30)",
			Type: "F1141 CHAR(200)",
		},
		{
			Name: "F1142 INTEGER",
			Type: "F1143 CHAR(30)",
		},
		{
			Name: "F1144 CHAR(30)",
			Type: "F1145 DATE(7)",
		},
		{
			Name: "F1146 CHAR(10)",
			Type: "F1148 CHAR(14)",
		},
		{
			Name: "F1176 CHAR(3)",
			Type: "F1264 DATE(7)",
		},
		{
			Name: "F1552 CHAR(4)",
			Type: "F1553 DATE(7)",
		},
		{
			Name: "F1554 DATE(7)",
			Type: "F1555 DATE(7)",
		},
		{
			Name: "F1556 INTEGER",
			Type: "F1557 CHAR(40)",
		},
		{
			Name: "F1558 CHAR(40)",
			Type: "F1559 CHAR(40)",
		},
		{
			Name: "F1560 CHAR(20)",
			Type: "F1561 CHAR(20)",
		},
		{
			Name: "F1562 CHAR(15)",
			Type: "F1563 CHAR(20)",
		},
		{
			Name: "F1564 CHAR(20)",
			Type: "F1565 CHAR(20)",
		},
		{
			Name: "F1566 NUMBER(10,0)",
			Type: "F1567 INTEGER",
		},
		{
			Name: "F1568 INTEGER",
			Type: "F1569 INTEGER",
		},
		{
			Name: "F1570 NUMBER(10,0)",
			Type: "F1585 CHAR(20)",
		},
		{
			Name: "F1586 CHAR(4)",
			Type: "F1587 CHAR(20)",
		},
		{
			Name: "F1588 CHAR(4)",
			Type: "F1589 CHAR(15)",
		},
		{
			Name: "F1590 CHAR(200)",
			Type: "F1964 CHAR(4)",
		},
		{
			Name: "F2587 DATE(7)",
			Type: "F2597 CHAR(20)",
		},
		{
			Name: "F2692 CHAR(1)",
			Type: "F2806 CHAR(40)",
		},
		{
			Name: "F2827 CHAR(10)",
			Type: "F2828 NUMBER(8,4)",
		},
		{
			Name: "F2829 NUMBER(8,4)",
			Type: "F2830 NUMBER(8,4)",
		},
		{
			Name: "F2831 NUMBER(8,4)",
			Type: "F2832 CHAR(20)",
		},
		{
			Name: "F2833 INTEGER",
			Type: "F2844 CHAR(200)",
		},
	}
}
