package parser

import (
	"fmt"

	"github.com/kitagry/monkey/token"
)

type ParserError struct {
	token   token.Token
	message string
}

func newParserError(msg string, tok token.Token) *ParserError {
	return &ParserError{
		token:   tok,
		message: msg,
	}
}

func (p *ParserError) Error() string {
	return fmt.Sprintf("%d:%d %s", p.token.Row, p.token.Col, p.message)
}
