package calc

import "fmt"

type Phrase interface {
	eval() Result
}

type Result struct {
	val interface{}
}

type ArithmeticPhrase struct {
	left  int
	right int
	op    string
}

func (a ArithmeticPhrase) eval() Result {
	fmt.Println(a)
	if a.op == "+" {
		return Result{a.left + a.right}
	}

	if a.op == "-" {
		return Result{a.left - a.right}
	}

	return Result{}
}
