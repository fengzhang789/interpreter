package lexer

import "interpreter/token"

type Lexer struct {
	input string
	position int // current position in input (points to curr char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination. ASCII VALUE IS PRINTED!
}

// NextToken() is a method of the Lexer type. Called by LexerInstance.NextToken()
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace();

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)
	
	case ',':
		tok = newToken(token.COMMA, l.ch)
	
	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)
	
	// END OF STRING
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
 		} else {
			tok - newToken(token.ILLEGAL, l.ch)
		}

	l.readChar()
	return tok
}

// SKIPS OVER A WHITESPACE.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
	l.readChar()
	}
	}

// newToken: takes a token type, a character ch, and creates a token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// New: creates a new Lexer struct
//   input: String
//   ouput: Lexer PTR
func New(input string) *Lexer {
	// creates a new lexer instance with the input field
	//   as input, and assigns the address of this instance
	//   to l. 
	l := &Lexer{input: input} 
	l.readChar();
	return l
}

// returns true if ch is a letter, false otherwise
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// a method on Lexer types that returns a string representing
//   the entire word or identifier. 
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
	}

// readChar: reads the next character and advance our position
//   in the input string. 
// input: Lexer PTR
func (l *Lexer) readChar() {
	// If this is true, there are no more characters to read
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}