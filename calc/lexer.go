package calc

import (
	"fmt"
	"strconv"
)

type Lexer struct {
	pos  int
	expr string
}

func NewLexer(e string) Lexer {
	return Lexer{expr: e}
}

func (lexer *Lexer) NextToken() (Token, error) {
	if lexer.reachedEnd() {
		return Eof(), nil
	}

	currentChar := lexer.currentChar()

	switch {
	case currentChar == " ":
		lexer.advance()
		lexer.NextToken()
	case currentChar == "+":
		lexer.advance()
		return Plus(), nil
	case currentChar == "-":
		lexer.advance()
		return Minus(), nil
	case isDigit(currentChar):
		return lexer.getIntegerToken()
	default:
		return Token{}, fmt.Errorf("Unexpected Character %s", currentChar)
	}
	return Token{}, fmt.Errorf("Unknown error")
}

func (lexer *Lexer) currentChar() string {
	return lexer.expr[lexer.pos : lexer.pos+1]
}

func (lexer *Lexer) reachedEnd() bool {
	return lexer.pos == len(lexer.expr)-1
}

func (lexer *Lexer) advance() {
	lexer.pos++
}

func (lexer *Lexer) getIntegerToken() (Token, error) {
	num := ""
	for !lexer.reachedEnd() && isDigit(lexer.currentChar()) {
		num = num + lexer.currentChar()
		lexer.advance()
	}

	integerVal, _ := strconv.Atoi(num)
	return NewToken(INTEGER, integerVal), nil
}

func isDigit(char string) bool {
	if _, err := strconv.Atoi(char); err != nil {
		return false
	}
	return true
}
