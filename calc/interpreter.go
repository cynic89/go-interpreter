package calc

import "fmt"

type Interpreter struct {
	parser *Parser
}

func NewInterpreter(p *Parser) Interpreter {
	return Interpreter{parser: p}
}

func (i Interpreter) Eval() (Result, error) {
	ast, err := i.parser.expr()
	fmt.Println(ast)
	if err != nil {
		return Result{}, err
	}
	result := ast.Accept(i)

	return Result{result}, nil
}

func (i Interpreter) VisitBinaryOp(b BinaryOp) int {
	switch b.op {
	case "+":
		return b.left.Accept(i) + b.right.Accept(i)
	case "-":
		return b.left.Accept(i) - b.right.Accept(i)
	case "*":
		return b.left.Accept(i) * b.right.Accept(i)
	case "/":
		return b.left.Accept(i) / b.right.Accept(i)
	}
	panic("Unknown Op")
}

func (i Interpreter) VisitNum(num Num) int {
	return num.val
}
