package calc

type Phrase interface {
	eval() Result
}

type Result struct {
	val interface{}
}

type ArithmeticPhrase struct {
	tokens []Token
}

func (a ArithmeticPhrase) eval() Result {
	var result int
	for i, token := range a.tokens {

		if i == 0 {
			result = a.tokens[i].value.(int)
		}

		if token.kind == ADD {
			result = result + a.tokens[i+1].value.(int)
		}

		if token.kind == SUB {
			result = result - a.tokens[i+1].value.(int)
		}
	}
	return Result{result}
}
