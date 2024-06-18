package silread

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// scanner represents a lex scanner

type scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func newScanner(r io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(r)}
}

func (s *scanner) scan() *part {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	switch ch {
	case eof:
		return &part{
			tok: EOF,
			val: "",
		}
	case ',':
		return &part{
			tok: COMMA,
			val: ",",
		}
	case '(':
		return &part{
			tok: OPEN,
			val: "(",
		}
	case ')':
		return &part{
			tok: CLOSE,
			val: ")",
		}
	case ';':
		return &part{
			tok: SEMICOLON,
			val: ";",
		}
	case '\'':
		return &part{
			tok: SINGLE,
			val: "'",
		}
	case '\r': // crlf newline detection, scans ahead to look for a newline
		ch = s.read()
		if ch != '\n' {
			log.Fatalf("carrage return without newline\n")
		}

		return &part{
			tok: CRLF,
			val: "\r\n",
		}
	case '\n': // this should never be reached, all newlines need to have a carrage return before it
		log.Fatalf("carrage return without newline\n")
	}

	s.unread()

	return s.scanIdent()
}

func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (s *scanner) scanIdent() *part {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	_, err := buf.WriteRune(s.read())

	if err != nil {
		log.Panicf("reading the character buffer faild %v", err)
	}

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && !isIncludeSpecial(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return &part{
		tok: IDENT,
		val: buf.String(),
	}
}

func isIncludeSpecial(ch rune) bool {

	switch ch {
	case '_':
		return true
	case '.':
		return true
	case '+':
		return true
	case '/':
		return true
	case ':':
		return true
	default:
		return false
	}

}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() *part {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	_, err := buf.WriteRune(s.read())

	if err != nil {
		log.Panicf("reading the character buffer faild %v", err)
	}

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			_, err := buf.WriteRune(ch)
			if err != nil {
				log.Panicf("reading the character buffer faild %v", err)
			}
		}
	}

	return &part{
		tok: WS,
		val: buf.String(),
	}
}

// unread places the previously read rune back on the reader.
func (s *scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return ch >= '0' && ch <= '9' }
