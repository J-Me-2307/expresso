package internal

import (
	"fmt"
	"strings"
	"unicode"
)

func (t *Token) String() string {
	return fmt.Sprintf("{Type: %s, Value: %s, Position: %d}", t.Type, t.Value, t.Position)
}

func NewToken(t TokenType, v string, p int) *Token {
	return &Token{
		Type:     t,
		Value:    v,
		Position: p,
	}
}

func Tokenize(expression string) []*Token {
	expression = strings.ReplaceAll(expression, " ", "")
	lenght := len(expression)
	runes := []rune(expression)
	i := 0

	tokens := make([]*Token, 0)

	for i < lenght {
		value := ""
		tokenType := INVALID_TOKEN
		commaCount := 0

		current := runes[i]

		if current == '.' || unicode.IsDigit(current) {
			tokenType = NUMBER_TOKEN

			for i < lenght {
				current = runes[i]

				if unicode.IsDigit(current) {
					value += string(current)
					i++
				} else if current == '.' {
					value += string(current)
					commaCount++
					i++
				} else {
					break
				}
			}

		} else if current == '+' || current == '*' || current == '/' {
			value += string(current)
			tokenType = OPERATOR_TOKEN
			i++
		} else if current == '-' {
			isFirst := i == 0
			var prev *Token

			if !isFirst {
				prev = tokens[len(tokens)-1]
			}

			if isFirst || prev.Type == LPAREN_TOKEN || (prev.Type == OPERATOR_TOKEN && prev.Value != "-") {
				tokenType = NUMBER_TOKEN
				value += string(current)
				i++

				for i < lenght {
					current = runes[i]

					if unicode.IsDigit(current) {
						value += string(current)
						i++
					} else if current == '.' {
						value += string(current)
						commaCount++
						if commaCount > 1 {
							tokenType = INVALID_TOKEN
							i++
							break
						}
						i++
					} else {
						break
					}
				}
				if len(value) == 1 {
					tokenType = OPERATOR_TOKEN
				}
			} else {
				value = string(current)
				tokenType = OPERATOR_TOKEN
				i++
			}

		} else if current == '(' {
			value += string(current)
			tokenType = LPAREN_TOKEN
			i++
		} else if current == ')' {
			value += string(current)
			tokenType = RPAREN_TOKEN
			i++
		} else {
			value = string(current)
			i++
		}

		if value == "." || value == "-." || commaCount > 1 {
			tokenType = INVALID_TOKEN
		}

		tokens = append(tokens, NewToken(tokenType, value, i))
	}
	return tokens
}
