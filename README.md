# SIL

## Create a SIL File in Go

README Needs updating and is outdated right now

Changing fairly often right now but starting to stabilize on needed API. 
High level goal is to have a type define the SIL structure and add data and get []bytes representing a complete SIL file that processes without errors.

Like JSON or XML tags are being used to define the structure of the SIL using a type. Currently there are two tags, sil and default. The sil tag is used to define the data type and default will fill in data that is missing but required. Pointers will be optional values but right now they are treated the same as non pointers and processed normally, I would recommend creating any types using pointers for optional values because they work right now and will eventually allow for more compact and smarter SIL file creation.

### Goals:
- [ ] Type
	- [x] Marshal function
	- [ ] Unmarshal function
	- [x] Allow pointers to be used for optional elements
	- [] single quote for most data types and no quotes for integers
		- [] confirm data is integer for integer types
	- [ ] Validate that element name is a proper SIL type
	- [ ] Validate that data passed for Rows matches type used for Make
	- [ ] Accept time.Time as type for DATE(7)
- [ ] Header
	- [x] Batch number creation
- [ ] View
	- [ ] Define table name with struct tag
	- [x] View Header
		- [x] respect default tag
		- [x] check for unsafe elements, correct or error as needed
	- [ ] View Data
		- [x] Create View data
		- [x] Only insert optional elements into SIL file when they are used.

### Eventuals
- [x] Read SQL structure to create Go Types for tables automagically


## Example
Needs to be re-written