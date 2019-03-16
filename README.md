# SIL

## Create a SIL File in Go

Changing fairly often right now but starting to stabilize on needed API. 
Working towards passing a SIL type and getting []bytes back of the created SIL file

Like JSON or XML types in golang tags are used to define the structure from a go type. Currently there are two tags, sil and default. The sil tag triggers that it will be used to create the sil file data and default will fill in data that is missing but required. Pointers will be optional values but right now they are treated the same as non pointers and processed normally. 

Features:
- [] Use GO type as definition
	- [x] Allow pointers to be used for optional elements
	- [x] single quote for most data types and no quotes for integers
		- [x] confirm data is integer for integer types
	- [ ] Validate that element name is a proper SIL type
	- [ ] Validate that data passed for Rows matches type used for Make
- [ ] Create View
	- [ ] Define table name with struct tag
	- [x] View Header
		- [x] respect default tag
		- [x] check for unsafe edits to either error or correct problems
	- [ ] View Data
		- [x] Create View data
		- [ ] only insert optional elements into SIL file when they are used


```Go

// Define type of batch - incomplete example, OBJ table has more required fields
type OBJ struct {
	F01 string `sil:"CHAR(13)"` // Product code, max 13 digits BUT is stored as text in LOC
	F16 int    `sil:"INTEGER"`
	F17 *int   `sil:"INTEGER"`
}
// Need to pass the name along with the type of data that will be passed
s := Make("OBJ", OBJ{})
// this will be a pointer so assigning now to point to later
n := 1
// assign a value with data in the pointer
s.View.Data = append(s.View.Data, OBJ{
	F01: "0000000009087",
	F16: 17,
	F17: &n,
})

// leave the pointer out, a row will still be inserted
s.View.Data = append(s.View.Data, OBJ{
	F01: "0000000009902",
	F16: 17,
})
// return []bytes of the sil file
b, err := s.Bytes()
if err != nil {
	fmt.Println(err)
}
// print the SIL file to console
fmt.Println(string(b))

```