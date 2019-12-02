package calc

type Interpreter struct {
	parser *Parser
}

func NewInterpreter(p *Parser) Interpreter {
	return Interpreter{parser: p}
}

func (i Interpreter) Eval() (Result, error) {
	result, err := i.parser.expr()
	if err != nil {
		return Result{}, err
	}

	return Result{result}, nil
}
