package silread

import (
	"io"
)

// parser does the actual parsing of bytes into sil type
type parser struct {
	s   *scanner
	buf struct {
		pt part
		n  int // buffer size (max=1)
	}
}

type parsed []part

type part struct {
	tok token
	val string
}

// NewParser returns a new instance of Parser.
func newParser(r io.Reader) *parser {
	return &parser{s: newScanner(r)}
}

func (p *parser) parse() *parsed {
	var prsd parsed

	for {
		pt := p.scan()
		prsd = append(prsd, *pt)

		if pt.tok == EOF {
			break
		}
	}

	return &prsd
}

func (p *parser) scan() (pt *part) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return &p.buf.pt
	}

	// Otherwise read the next token from the scanner.
	pt = p.s.scan()

	// Save it to the buffer in case we unscan later.
	p.buf.pt = *pt

	return
}
