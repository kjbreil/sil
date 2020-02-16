package silread

// token is an identifier for a given token
type token int

// the tokens to look for
const (
	ILLEGAL token = iota
	EOF
	WS

	IDENT // fields

	COMMA
	OPEN
	CLOSE
	SEMICOLON
	SINGLE

	CRLF
)
