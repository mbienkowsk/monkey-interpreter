package lexer

import (
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int                 // current position in input (points to current char)
	readPosition int                 // current reading position in input (after current char)
	ch           byte                // current char under examination
	lPosition    token.LexerPosition // current line and column number
}

func (l *Lexer) LineNumber() int {
	return l.lPosition.LineNumber
}

func (l *Lexer) ColumnNumber() int {
	return l.lPosition.ColumnNumber
}

func (l *Lexer) readChar() {
	l.lPosition.Advance(l.ch)
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		lPosition: token.LexerPosition{
			LineNumber:   1,
			ColumnNumber: 0,
		},
	}
	l.readChar()
	return l
}

func newToken(t token.TokenType, ch byte, lp token.LexerPosition) token.Token {
	return token.Token{Type: t, Literal: string(ch), Position: lp}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "==", Position: l.lPosition}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.lPosition)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.lPosition)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.lPosition)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.lPosition)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.lPosition)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.lPosition)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.lPosition)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.lPosition)
	case '>':
		tok = newToken(token.GT, l.ch, l.lPosition)
	case '<':
		tok = newToken(token.LT, l.ch, l.lPosition)
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: "!=", Position: l.lPosition}
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch, l.lPosition)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch, l.lPosition)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.lPosition)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.lPosition)
	case '"':
		tok.Position = l.lPosition
		tok.Literal = l.readString()
		tok.Type = token.STRING
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.lPosition)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.lPosition)
	case ':':
		tok = newToken(token.COLON, l.ch, l.lPosition)
	case 0:
		tok.Position = l.lPosition
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok.Position = l.lPosition
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.lPosition)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	start_position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start_position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	start_position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start_position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	l.readChar() // skip the first quote
	start_position := l.position

	for l.ch != '"' && l.ch != 0 { // todo: error handling, escaping
		l.readChar()
	}
	end_position := l.position

	return l.input[start_position:end_position]
}
