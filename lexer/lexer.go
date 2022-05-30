package lexer

import (
	"github.com/coopstools/basic/token"
)

type Lexer struct {
	Input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) Lexer {
	l := Lexer{Input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()
	tType := token.LookupChr(l.ch)

	switch {
	case tType != token.Type(""):
		t = token.Token{Type: tType, Literal: string(l.ch)}
	case l.ch == '"':
		t = token.Token{Type: token.STRING, Literal: l.readString()}
	case isLetter(l.ch):
		identifier := l.readIdentifier()
		return token.Token{Type: token.LookupIdent(identifier), Literal: identifier}
	case isDigit(l.ch):
		return token.Token{Type: token.INT, Literal: l.readNumber()}
	default:
		t = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}
	l.readChar()
	return t
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.Input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.Input[position:l.position]
}

func (l *Lexer) readString() string {
	l.readChar()
	position := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	return l.Input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_' || ch == '$' || ch == '#' || ch == '!'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	l.ch = 0
	if l.readPosition < len(l.Input) {
		l.ch = l.Input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
