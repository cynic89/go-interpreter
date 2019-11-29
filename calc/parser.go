package calc

import (
	"fmt"
	"strconv"
)

type Parser struct {
	expr expr
	pos  int
}

type expr string

func NewParser(e string) Parser {
	return Parser{expr(e), 0}
}

func (p *Parser) Parse() ([]Token, error) {
	var tokens []Token
	for !p.reachedEnd() {
		token, err := p.NextToken()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (p *Parser) NextToken() (token Token, error error) {
	defer func() {
		p.pos++

		if error == nil{
			if token.kind == INTEGER {
				nt, ne := p.NextToken()
				if ne != nil {
					error = ne
					return
				}

				if nt.kind != INTEGER {
					p.pos--
					return
				}

				token.value = token.value + nt.value
			}

			if token.kind == WHITESPACE {
				nt, ne := p.NextToken()
				if ne != nil {
					error = ne
					return
				}
				token = nt
			}
		}

	}()

	if p.pos == len(p.expr) {
		token = Eof()
		return
	}

	char := p.expr.nextChar(p.pos)

	if _, err := strconv.Atoi(char); err == nil {
		token = NewToken(INTEGER, char)
		return
	}

	if "+" == char {
		token = Plus()
		return
	}

	if " " == char {
		token = Whitespace()
		return
	}

	error = fmt.Errorf("Unexpected Token %s", char)
	return

}

func (p *Parser) reachedEnd() bool {
	if p.pos == len(p.expr) {
		return true
	}
	return false
}

func (e expr) nextChar(pos int) string {
	return string(e[pos : pos+1])
}
