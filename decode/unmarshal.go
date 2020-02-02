package decode

import (
	"bytes"
	"fmt"
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

	d := prsd.decode()
	if len(d.err) > 0 {
		log.Fatalf("could not decode the parsed sil file: %v\n", d.err)
	}

	fmt.Println(d.data)

	// fmt.Println(d.p[len(d.p)-1], len(d.p), d.fcodes)

}
