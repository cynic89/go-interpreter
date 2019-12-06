package calc

import "fmt"

type Visitor interface {
	VisitBinaryOp(b BinaryOp) int
	VisitNum(n Num) int
}

type AST interface {
	Accept(visitor Visitor) int
}

type Num struct {
	val int
}

type BinaryOp struct {
	op    Op
	left  AST
	right AST
}

type GenericNode struct {
}

type Op string

func (binaryOp BinaryOp) Accept(visitor Visitor) int {
	return visitor.VisitBinaryOp(binaryOp)
}

func (num Num) Accept(visitor Visitor) int {
	return visitor.VisitNum(num)
}

func (node GenericNode) Accept(visitor Visitor) int {
	fmt.Println("Do Nothing")
	return -1
}
