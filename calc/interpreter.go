package calc

import (
	"fmt"
)

type Interpreter struct {
	currentToken Token
}

type Result struct {
	int
}

func (i Interpreter) Eval(expr string) (Result, error) {
	parser := NewParser(expr)
	token, err := parser.NextToken()

	if err != nil {
		return Result{}, err
	}
	i.currentToken = token
	left := i.currentToken
	err = i.eat(INTEGER, &parser)
	if err != nil {
		return Result{}, err
	}

	//_ := i.currentToken
	err = i.eat(PLUS, &parser)
	if err != nil {
		return Result{}, err
	}

	right := i.currentToken
	err = i.eat(INTEGER, &parser)
	if err != nil {
		return Result{}, err
	}
	leftInt := left.GetValue().(int)
	rightInt := right.GetValue().(int)
	return Result{leftInt + rightInt}, nil
}
func (i *Interpreter) eat(tokenType TokenType, parser *Parser) error {
	if tokenType == i.currentToken.kind {
		token, err := parser.NextToken()
		if err != nil {
			return err
		}
		i.currentToken = token
		return nil
	}
	return fmt.Errorf("Token types do not match %s != %s", tokenType, i.currentToken.kind)
}
