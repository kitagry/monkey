package lexer

import (
	"github.com/kitagry/monkey/token"
)

type Lexer struct {
	input string
	// 入力における現在の位置(現在の文字を指し示す)
	position int
	// これから読み込む位置
	readPosition int
	// 現在検査中の文字
	ch byte
	// 現在読み込み中の行
	curRow int
	// 現在読み込み中の列
	curCol int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, curRow: 1, curCol: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCIIの"NUL"に対応
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.curRow++
		l.curCol = 0
	} else {
		l.curCol++
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Row: l.curRow, Col: l.curCol}
			tok.Literal = l.readTwoChars()
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.curRow, l.curCol)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Row: l.curRow, Col: l.curCol}
			tok.Literal = l.readTwoChars()
		} else {
			tok = newToken(token.BANG, l.ch, l.curRow, l.curCol)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.curRow, l.curCol)
	case ':':
		tok = newToken(token.COLON, l.ch, l.curRow, l.curCol)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.curRow, l.curCol)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.curRow, l.curCol)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.curRow, l.curCol)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.curRow, l.curCol)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.curRow, l.curCol)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.curRow, l.curCol)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.curRow, l.curCol)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.curRow, l.curCol)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.curRow, l.curCol)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.curRow, l.curCol)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.curRow, l.curCol)
	case '<':
		tok = newToken(token.LT, l.ch, l.curRow, l.curCol)
	case '>':
		tok = newToken(token.GT, l.ch, l.curRow, l.curCol)
	case '"':
		tok.Row = l.curRow
		tok.Col = l.curCol
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Row = l.curRow
		tok.Col = l.curCol
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Row = l.curRow
			tok.Col = l.curCol
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.ch) {
			tok.Row = l.curRow
			tok.Col = l.curCol
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch, l.curRow, l.curCol)
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte, row, col int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Row: row, Col: col}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readTwoChars() string {
	ch := l.ch
	l.readChar()
	return string(ch) + string(l.ch)
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
