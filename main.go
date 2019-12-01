package main

import (
	"bufio"
	"fmt"
	"github.com/cynic89/go-interpreter/calc"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("calc>")
	text, _ := reader.ReadString('\n')
	lexer := calc.NewLexer(text)
	parser := calc.NewParser(&lexer)
	interpreter := calc.NewInterpreter(&parser)

	result, err := interpreter.Eval()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)

}
