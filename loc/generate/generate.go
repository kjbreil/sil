package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

// TypeFile holds the info for each type file
type TypeFile struct {
	Name      string
	FileName  string
	TableName string
	Data      []Element
}

// Element is the info about each element in the type
type Element struct {
	Name      string
	Type      string
	Sil       string
	Arguments []string
	Default   string
	Required  bool
}

func main() {
	// get the records from the csv
	records := read()
	// first loop over the records and group them into a map
	structure := make(map[string][][]string)
	for r := range records {
		structure[records[r][0]] = append(structure[records[r][0]], records[r])
	}

	// make each file
	for k := range structure {
		// convert the table name into lowercase snake for filename
		// and Camel case for the type name
		tf := TypeFile{
			Name:      strcase.ToCamel(strings.ToLower(k)),
			TableName: k,
			FileName:  strcase.ToSnake(k),
		}
		// this is horrible - should use pointers to reduce overhead
		for i := range structure[k] {
			elem := Element{
				Name:      strcase.ToCamel(structure[k][i][2]),
				Type:      dataType(structure[k][i][6]),
				Sil:       structure[k][i][1],
				Arguments: arguments(structure[k][i][1]),
				Default:   hasDefault(structure[k][i][1]),
				Required:  isRequired(structure[k][i][1]),
			}
			tf.Data = append(tf.Data, elem)
		}
		if tf.Name == "CfgTab" {
			tf.write()
			// fmt.Println(tf.marshal())
		}
	}

}

func read() [][]string {
	f, err := os.Open("./RB_FIELDS.csv")
	if err != nil {
		log.Fatal(err)
	}
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func dataType(dt string) string {
	switch dt {
	case "dtInteger":
		return "int"
	default:
		return "string"
	}
}

// marshal turns the TypeFile into a string
func (tf *TypeFile) marshal() []byte {
	packageName := "loc"
	// start the file
	out := fmt.Sprintf("package %s\n\n", packageName)

	out = out + fmt.Sprintf("// %s is the %s definition\n", tf.Name, tf.TableName)

	out = out + fmt.Sprintf("type %s struct {\n", tf.Name)
	for i := range tf.Data {
		var pNeed string
		// add an astricks if the item is not required making the type a pointer
		if !tf.Data[i].Required {
			pNeed = "*"
		}

		// parse and insert arguments
		var args string
		if len(tf.Data[i].Arguments) > 0 {
			args = ","
			for _, ea := range tf.Data[i].Arguments {
				args = args + ea + ","
			}
			// remove the last comma
			args = strings.TrimSuffix(args, ",")
		}

		// default string
		var def string
		if tf.Data[i].Default != "" {
			def = fmt.Sprintf(" default:\"%s\"", tf.Data[i].Default)
		}

		out = out + fmt.Sprintf("    %s %s%s `sil:\"%s%s\"%s`\n", tf.Data[i].Name, pNeed, tf.Data[i].Type, tf.Data[i].Sil, args, def)
	}
	out = out + fmt.Sprint("}\n")

	return []byte(out)
}

// when more arguments are used this should spin off to a zeropad function and
// this function fires that and then the new one off
func arguments(silCode string) []string {
	zeropad := []string{
		"F01",
	}

	for _, er := range zeropad {
		if silCode == er {
			return []string{"zeropad"}
		}
	}

	return []string{}

}

// isRequired is very basic for now, should come from a seperate file
// really LOC should be using their required tag to make this easier since even
// F01 does not contain it on a default install
func isRequired(silCode string) bool {
	required := []string{
		"F01",
		"F164",
		"F253",
		"F1000",
		"F1001",
		"F1031",
		"F1033",
		"F1034",
		"F1056",
		"F1264",
		"F2845",
		"F2846",
		"F2847",
	}

	for _, er := range required {
		if silCode == er {
			return true
		}
	}

	return false
}

func (tf *TypeFile) write() {
	filename := fmt.Sprintf("./gen/%s.go", tf.FileName)

	ioutil.WriteFile(filename, tf.marshal(), 0666)
}

func hasDefault(silCode string) string {
	switch silCode {
	case "F1001":
		return "1"
	case "F253":
		return "NOW"
	case "F1264":
		return "NOW"
	default:
		return ""

	}
}
