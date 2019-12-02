package calc

import (
	"fmt"
)

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(l *Lexer) (Parser, error) {
	token, err := l.NextToken()
	if err != nil {
		return Parser{}, err
	}
	return Parser{lexer: l, currentToken: token}, nil
}

/*

grammar

expr: term ( (ADD | SUB) term)*
term: factor ( (MUL | DIV) factor )*
factor: INTEGER

*/

// factor: INTEGER
func (p *Parser) factor() (int, error) {
	intVal, _ := p.currentToken.value.(int)
	err := p.eat(INTEGER)
	if err != nil {
		return -1, err
	}
	return intVal, nil

}

// term: factor ( (MUL | DIV) factor )*
func (p *Parser) term() (int, error) {
	result, err := p.factor()
	if err != nil {
		return -1, err
	}

	for p.currentToken.kind == MUL || p.currentToken.kind == DIV {
		if p.currentToken.kind == MUL {
			err := p.eat(MUL)
			if err != nil {
				return -1, err
			}

			operand, err := p.factor()
			if err != nil {
				return -1, err
			}

			result = result * operand
		} else if p.currentToken.kind == DIV {
			err := p.eat(DIV)
			if err != nil {
				return -1, err
			}

			operand, err := p.factor()
			if err != nil {
				return -1, err
			}

			result = result / operand
		}
	}

	return result, nil
}

// expr: term ( (ADD | SUB) term)*
func (p *Parser) expr() (int, error) {
	result, err := p.term()
	if err != nil {
		return -1, err
	}

	for p.currentToken.kind == ADD || p.currentToken.kind == SUB {
		if p.currentToken.kind == ADD {
			err := p.eat(ADD)
			if err != nil {
				return -1, err
			}

			operand, err := p.term()
			if err != nil {
				return -1, err
			}

			result = result + operand
		} else if p.currentToken.kind == SUB {
			err := p.eat(SUB)
			if err != nil {
				return -1, err
			}

			operand, err := p.term()
			if err != nil {
				return -1, err
			}

			result = result - operand
		}
	}

	return result, nil
}

func (p *Parser) eat(tokenType TokenType) error {
	if p.currentToken.kind == tokenType {
		token, err := p.lexer.NextToken()
		if err != nil {
			return err
		}
		p.currentToken = token
		return nil
	}

	return fmt.Errorf("Syntax Error: invalid syntax: `%s` , Expected `%s`", p.currentToken.kind, tokenType)
}
