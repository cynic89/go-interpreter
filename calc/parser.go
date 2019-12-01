package calc

import "fmt"

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(l *Lexer) Parser {
	return Parser{lexer: l}
}

func (p *Parser) Parse() (Phrase, error) {
	left, err := p.lexer.NextToken()
	if err != nil {
		return nil, err
	}
	p.currentToken = left

	err = p.eat(INTEGER)
	if err != nil {
		return nil, err
	}

	op := p.currentToken
	if op.kind == PLUS {
		err = p.eat(PLUS)
	} else {
		err = p.eat(MINUS)
	}
	if err != nil {
		return nil, err
	}

	right := p.currentToken
	err = p.eat(INTEGER)
	if err != nil {
		return nil, err
	}

	return ArithmeticPhrase{left.value.(int), right.value.(int), op.value.(string)}, nil

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

	return fmt.Errorf("Token types do not match %s != %s", p.currentToken.kind, tokenType)
}
