package decode

import (
	"bytes"
	"log"
)

// Unmarshal SIL bytes into a interface{}
func Unmarshal(b []byte, v interface{}) {
	// open a reader using the bytes as the start
	// this can be improved to read directly from a file
	r := bytes.NewReader(b)

	// make a new parser with the reader
	p := newParser(r)
	// prsd is the parsed file token parts
	prsd := p.parse()

	// just displaying the parts for now
	// for _, ep := range *prsd {
	// 	fmt.Println(ep)
	// }

	err := prsd.decode()

	if err != nil {
		log.Fatalf("could not decode the parsed sil file: %v\n", err)
	}
}
