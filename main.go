package main

import (
	"bufio"
	"fmt"
	"github.com/cynic89/go-interpreter/calc"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	interpreter := calc.Interpreter{}
	fmt.Print("calc>")
	text, _ := reader.ReadString('\n')

	result, err := interpreter.Eval(strings.Trim(text, "\n"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(result)

}
