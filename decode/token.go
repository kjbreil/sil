package decode

// token is an identifier for a given token
type token int

const (
	ILLEGAL token = iota
	EOF
	WS

	IDENT // fields

	COMMA
	OPEN
	CLOSE
	SEMICOLON

	CRLF
)
