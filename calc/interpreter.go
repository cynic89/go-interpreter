package calc

type Interpreter struct {
	parser *Parser
}

func NewInterpreter(p *Parser) Interpreter {
	return Interpreter{parser: p}
}

func (i Interpreter) Eval() (Result, error) {
	phrase, err := i.parser.Parse()
	if err != nil {
		return Result{}, err
	}

	result := phrase.eval()
	return result, nil
}
