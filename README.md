# SIL

## Create a SIL File in Go

Changing fairly often right now but starting to stabilize on needed API. 
Working towards passing a SIL type and getting []bytes back of the created SIL file

Mostly stripped of CLK specific code, just not tested with another type yet. Refactoring code needs work and cleaning up.

A tag is needed for the F code, this way the type elements can be human readable

Features:
- [x] Use GO type as definition
	- [x] Allow pointers to be used for optional elements
	- [x] single quote for most data types and no quotes for integers
		- [x] confirm data is integer for integer types
- [] Create View
	- [] Define table name with struct tag
	- [x] View Header
		- [x] respect default tag

		- [x] check for unsafe edits to either error or correct problems
	- [] View Data
		- [] only insert optional elements into SIL file when they are used


```Go

// Define type of batch - incomplete example, OBJ table has more required fields
type OBJ struct {
	F01 string `sil:"CHAR(13)"` // Product code, max 13 digits BUT is stored as text in LOC
	F16 int    `sil:"INTEGER"`
	F17 *int   `sil:"INTEGER"`
}

// I'll write more soon

```