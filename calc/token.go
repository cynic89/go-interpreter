package calc

type Token struct {
	kind  TokenType
	value TokenVal
}

type TokenVal interface{}
type TokenType string

const (
	INTEGER    TokenType = "Integer"
	PLUS       TokenType = "Plus"
	MINUS      TokenType = "Minus"
	EOF        TokenType = "EOF"
	WHITESPACE TokenType = "Whitespace"
)

func NewToken(kind TokenType, val interface{}) Token {
	return Token{kind, TokenVal(val)}
}

func Eof() Token {
	return Token{EOF, nil}
}

func Plus() Token {
	return Token{PLUS, "+"}
}

func Minus() Token {
	return Token{MINUS, "-"}
}

func Whitespace() Token {
	return Token{WHITESPACE, " "}
}

