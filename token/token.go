package token

type TokenType string

// Struct to hold a Token, each token has its own type, int, number, string etc. It also has its own value
//  which is the "Literal".
type Token struct {
	Type TokenType
	Literal string
}

// Dictionary of extra keywords used in our language
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

// Defining all the different types of token types in the language. 
//   ILLEGAL is used for a token we dont know about, and EOF means end of file.
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456
	// Operators
	ASSIGN = "="
	PLUS = "+"
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)
	