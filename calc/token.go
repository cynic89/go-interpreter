package calc

import "strconv"

type Token struct {
	kind TokenType
	value TokenVal
}

type TokenVal string
type TokenType string

const (
	INTEGER TokenType = "Integer"
	PLUS TokenType = "Plus"
	EOF TokenType = "EOF"
	WHITESPACE TokenType = "Whitespace"

	NONE TokenVal = ""
)

func NewToken(kind TokenType, val string) Token  {
	return Token{kind, TokenVal(val)}
}

func Eof() Token{
	return Token {EOF, NONE}
}

func Plus() Token{
	return Token{PLUS, "+"}
}

func Whitespace() Token{
	return Token{WHITESPACE, " "}
}


func (t Token) GetValue() interface{}  {
	if t.kind == INTEGER {
		if v, err := strconv.Atoi(string(t.value)); err == nil{
			return v
		}
	}
	return t.value
}