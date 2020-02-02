package silread

// token is an identifier for a given token
type token int

// the tokens to look for
// TODO: shoudl add in INSERT and CREATE as tokens/types that hold the info on each line within
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
