package sil

// User is easy to fill CLK struct
type User struct {
	Number int
	First  string
	Last   string
	Short  string
	Level  int
}

// CLK is the structure of a CLK insert
// required fields are regular variables, optional are pointers
type CLK struct {
	F1185 int     // User order number
	F1001 int     // Record status
	F1126 int     // User number
	F1571 *string // User Email address
	F27   *string
	F170  *string
	F253  string // Last change date
	F902  string // Batch identifier
	F940  *string
	F941  *string
	F1000 string  // Target Identifier
	F1056 *string // Terminal store
	F1127 string  // User short name
	F1141 *string
	F1142 int // User restriction level
	F1143 *string
	F1144 *string
	F1145 *string
	F1146 *string
	F1148 *string
	F1176 *string
	F1264 *string
	F1552 *string
	F1553 *string
	F1554 *string
	F1555 *string
	F1556 *string
	F1557 *string
	F1558 *string
	F1559 *string
	F1560 *string
	F1561 *string
	F1562 *string
	F1563 *string
	F1564 *string
	F1565 *string
	F1566 *string
	F1567 *string
	F1568 *string
	F1569 *string
	F1570 *string
	F1585 *string
	F1586 *string
	F1587 *string
	F1588 *string
	F1589 *string // User Employee number
	F1590 *string
	F1964 string  // Store responsible
	F2587 *string // Clock in/out
	F2597 *string
	F2692 *string
	F2806 *string
	F2827 *string
	F2828 *string
	F2829 *string
	F2830 *string
	F2831 *string
	F2832 *string
	F2833 *string
	F2844 *string
}
