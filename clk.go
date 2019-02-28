package sil

// User is easy to fill CLK struct
type User struct {
	Number    int
	First     string
	Last      string
	Short     string
	Level     int
	Birthdate string
}

// CLK is the structure of a CLK insert
// All fields are required - this is because we need to fill in every
// column in a SIL file so pointers will eventually cause errors
// right now they will just cause pain
type CLK struct {
	F1185 int    `sil:"INTEGER"`  // User order number
	F1001 int    `sil:"INTEGER"`  // Record status
	F1126 int    `sil:"INTEGER"`  // User number
	F253  string `sil:"DATE(7)"`  // Last change date
	F902  string `sil:"CHAR(8)"`  // Batch identifier
	F1000 string `sil:"CHAR(5)"`  // Target Identifier
	F1127 string `sil:"CHAR(30)"` // User short name
	F1142 int    `sil:"INTEGER"`  // User restriction level
	F1143 string `sil:"CHAR(30)"`
	F1144 string `sil:"CHAR(30)"`
	F1145 string `sil:"DATE(7)"` // Operator Birthdate
	F1964 string `sil:"CHAR(4)"` // Store responsible
}
