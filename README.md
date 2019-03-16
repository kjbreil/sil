# SIL

## Create a SIL File in Go

Changing fairly often right now but starting to stabilize on needed API. 
High level goal is to have a type define the SIL structure and add data and get []bytes representing a complete SIL file that processes without errors.

Like JSON or XML tags are being used to define the structure of the SIL using a type. Currently there are two tags, sil and default. The sil tag is used to define the data type and default will fill in data that is missing but required. Pointers will be optional values but right now they are treated the same as non pointers and processed normally, I would recommend creating any types using pointers for optional values because they work right now and will eventually allow for more compact and smarter SIL file creation.

Features:
- [ ] Use GO type as definition
	- [x] Allow pointers to be used for optional elements
	- [x] single quote for most data types and no quotes for integers
		- [x] confirm data is integer for integer types
	- [ ] Validate that element name is a proper SIL type
	- [ ] Validate that data passed for Rows matches type used for Make
- [ ] Create View
	- [ ] Define table name with struct tag
	- [x] View Header
		- [x] respect default tag
		- [x] check for unsafe elements, correct or error as needed
	- [ ] View Data
		- [x] Create View data
		- [ ] only insert optional elements into SIL file when they are used

## Example

```Go
package main

// Define type of batch - incomplete example, OBJ table has more required fields
type OBJ struct {
	F01 string `sil:"CHAR(13)"` // Product code, max 13 digits BUT is stored as text in LOC
	F16 int    `sil:"INTEGER"`
	F17 *int   `sil:"INTEGER"`
}

func main() {
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
}
```