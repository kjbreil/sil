# SIL

## Create a SIL File in Go

Changing fairly often right now but starting to stabilize on needed API. 
Working towards passing a SIL type and getting []bytes back of the created SIL file

Mostly stripped of CLK specific code, just not tested with another type yet. Refactoring code needs work and cleaning up.

A tag is needed for the F code, this way the type elements can be human readable

```Go

// Define type of batch
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

// I'll write more soon

```