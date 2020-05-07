package lexer_test

import (
	"testing"

	"github.com/kitagry/monkey/lexer"
	"github.com/kitagry/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if ( 5 < 10 ) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
macro(x, y) { x + y; };
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedRow     int
		expectedCol     int
	}{
		{token.LET, "let", 1, 1},
		{token.IDENT, "five", 1, 5},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},
		{token.LET, "let", 2, 1},
		{token.IDENT, "ten", 2, 5},
		{token.ASSIGN, "=", 2, 9},
		{token.INT, "10", 2, 11},
		{token.SEMICOLON, ";", 2, 13},
		{token.LET, "let", 4, 1},
		{token.IDENT, "add", 4, 5},
		{token.ASSIGN, "=", 4, 9},
		{token.FUNCTION, "fn", 4, 11},
		{token.LPAREN, "(", 4, 13},
		{token.IDENT, "x", 4, 14},
		{token.COMMA, ",", 4, 15},
		{token.IDENT, "y", 4, 17},
		{token.RPAREN, ")", 4, 18},
		{token.LBRACE, "{", 4, 20},
		{token.IDENT, "x", 5, 3},
		{token.PLUS, "+", 5, 5},
		{token.IDENT, "y", 5, 7},
		{token.SEMICOLON, ";", 5, 8},
		{token.RBRACE, "}", 6, 1},
		{token.SEMICOLON, ";", 6, 2},
		{token.LET, "let", 8, 1},
		{token.IDENT, "result", 8, 5},
		{token.ASSIGN, "=", 8, 12},
		{token.IDENT, "add", 8, 14},
		{token.LPAREN, "(", 8, 17},
		{token.IDENT, "five", 8, 18},
		{token.COMMA, ",", 8, 22},
		{token.IDENT, "ten", 8, 24},
		{token.RPAREN, ")", 8, 27},
		{token.SEMICOLON, ";", 8, 28},
		{token.BANG, "!", 9, 1},
		{token.MINUS, "-", 9, 2},
		{token.SLASH, "/", 9, 3},
		{token.ASTERISK, "*", 9, 4},
		{token.INT, "5", 9, 5},
		{token.SEMICOLON, ";", 9, 6},
		{token.INT, "5", 10, 1},
		{token.LT, "<", 10, 3},
		{token.INT, "10", 10, 5},
		{token.GT, ">", 10, 8},
		{token.INT, "5", 10, 10},
		{token.SEMICOLON, ";", 10, 11},
		{token.IF, "if", 12, 1},
		{token.LPAREN, "(", 12, 4},
		{token.INT, "5", 12, 6},
		{token.LT, "<", 12, 8},
		{token.INT, "10", 12, 10},
		{token.RPAREN, ")", 12, 13},
		{token.LBRACE, "{", 12, 15},
		{token.RETURN, "return", 13, 2},
		{token.TRUE, "true", 13, 9},
		{token.SEMICOLON, ";", 13, 13},
		{token.RBRACE, "}", 14, 1},
		{token.ELSE, "else", 14, 3},
		{token.LBRACE, "{", 14, 8},
		{token.RETURN, "return", 15, 2},
		{token.FALSE, "false", 15, 9},
		{token.SEMICOLON, ";", 15, 14},
		{token.RBRACE, "}", 16, 1},
		{token.INT, "10", 18, 1},
		{token.EQ, "==", 18, 4},
		{token.INT, "10", 18, 7},
		{token.SEMICOLON, ";", 18, 9},
		{token.INT, "10", 19, 1},
		{token.NOT_EQ, "!=", 19, 4},
		{token.INT, "9", 19, 7},
		{token.SEMICOLON, ";", 19, 8},
		{token.STRING, "foobar", 20, 1},
		{token.STRING, "foo bar", 21, 1},
		{token.LBRACKET, "[", 22, 1},
		{token.INT, "1", 22, 2},
		{token.COMMA, ",", 22, 3},
		{token.INT, "2", 22, 5},
		{token.RBRACKET, "]", 22, 6},
		{token.SEMICOLON, ";", 22, 7},
		{token.LBRACE, "{", 23, 1},
		{token.STRING, "foo", 23, 2},
		{token.COLON, ":", 23, 7},
		{token.STRING, "bar", 23, 9},
		{token.RBRACE, "}", 23, 14},
		{token.MACRO, "macro", 24, 1},
		{token.LPAREN, "(", 24, 6},
		{token.IDENT, "x", 24, 7},
		{token.COMMA, ",", 24, 8},
		{token.IDENT, "y", 24, 10},
		{token.RPAREN, ")", 24, 11},
		{token.LBRACE, "{", 24, 13},
		{token.IDENT, "x", 24, 15},
		{token.PLUS, "+", 24, 17},
		{token.IDENT, "y", 24, 19},
		{token.SEMICOLON, ";", 24, 20},
		{token.RBRACE, "}", 24, 22},
		{token.SEMICOLON, ";", 24, 23},
		{token.EOF, "", 25, 1},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Row != tt.expectedRow {
			t.Fatalf("test[%d] - row wrong. expected=%d, got=%d", i, tt.expectedRow, tok.Row)
		}

		if tok.Col != tt.expectedCol {
			t.Fatalf("test[%d] - col wrong. expected=%d, got=%d", i, tt.expectedCol, tok.Col)
		}
	}
}
