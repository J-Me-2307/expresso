package internal

type TokenType string

const (
	NUMBER_TOKEN   TokenType = "Number"
	OPERATOR_TOKEN TokenType = "Operator"
	LPAREN_TOKEN   TokenType = "Left parenthesis"
	RPAREN_TOKEN   TokenType = "Right parenthesis"
	INVALID_TOKEN  TokenType = "Invalid"
)

type Token struct {
	Type     TokenType
	Value    string
	Position int
}
