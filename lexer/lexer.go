package lexer

import "github.com/MoonShining/condition/token"

type Lexer struct {
	input        []rune
	position     int // current position in input(current char)
	readPosition int // current reading position in input(after current char)
	ch           rune
}

func New(in string) *Lexer {
	l := &Lexer{input: []rune(in)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	p := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[p:l.position])
}

func (l *Lexer) readNumber() string {
	p := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[p:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()

	switch l.ch {
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.Token{Type: token.GTE, Literal: ">="}
		} else {
			t = newToken(token.GT, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.Token{Type: token.LTE, Literal: "<="}
		} else {
			t = newToken(token.LT, l.ch)
		}
	case '=':
		// look ahead to decide
		if l.peekChar() == '=' {
			l.readChar()
			t = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '[':
		t = newToken(token.LBRACKET, l.ch)
	case ']':
		t = newToken(token.RBRACKET, l.ch)
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			t = token.Token{Type: token.AND, Literal: "&&"}
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			t = token.Token{Type: token.OR, Literal: "||"}
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return t
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(typ token.TokenType, ch rune) token.Token {
	return token.Token{Type: typ, Literal: string(ch)}
}
