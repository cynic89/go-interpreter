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

	var tokens []Token

	token, err := p.lexer.NextToken()
	if err != nil {
		return nil, err
	}

	p.currentToken = token
	tokens = append(tokens, p.currentToken)

	err = p.eat(INTEGER)
	if err != nil {
		return nil, err
	}

	tSeq, err := p.parseSequence()
	if err != nil {
		return nil, err
	}
	tokens = append(tokens, tSeq...)

	for p.currentToken.kind != EOF {
		tSeq, err := p.parseSequence()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tSeq...)
	}

	fmt.Println(tokens)

	return ArithmeticPhrase{tokens: tokens}, nil

}

func (p *Parser) parseSequence() (tokens []Token, err error) {
	tokens = append(tokens, p.currentToken)

	op := p.currentToken
	if op.kind == PLUS {
		err = p.eat(PLUS)
	} else {
		err = p.eat(MINUS)
	}
	if err != nil {
		return nil, err
	}

	tokens = append(tokens, p.currentToken)
	err = p.eat(INTEGER)
	if err != nil {
		return nil, err
	}

	return tokens, err
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
