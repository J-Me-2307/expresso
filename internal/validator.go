package internal

import (
	"fmt"
)

type ParseError struct {
	Message  string
	Position int
	Got      string
	Expected string
}

func (e *ParseError) Error() string {
	base := fmt.Sprintf("parse error at position %d: %s", e.Position, e.Message)

	if e.Got != "" && e.Expected != "" {
		base += fmt.Sprintf(" (got '%s', expected %s)", e.Got, e.Expected)
	}

	return base
}

func ValidateTokens(tokens []*Token) []error {
	errs := make([]error, 0)

	if len(tokens) == 0 {
		return []error{fmt.Errorf("empty expression")}
	}

	// Check if first token is operator
	if tokens[0].Type == OPERATOR_TOKEN && tokens[0].Value != "-" {
		errs = append(errs, &ParseError{"Expression cannot start with an operator", tokens[0].Position, tokens[0].Value, "number or '('"})
	}

	// Check if last token is operator
	if tokens[len(tokens)-1].Type == OPERATOR_TOKEN {
		errs = append(errs, &ParseError{"Expression cannot end with an operator", len(tokens), tokens[len(tokens)-1].Value, "number or ')'"})
	}

	parentheseStack := make([]*Token, 0)

	for i, token := range tokens {
		// Check for invalid tokens
		if token.Type == INVALID_TOKEN {
			errs = append(errs, &ParseError{"Unmatched closing parentheses", token.Position, "", ""})
			continue
		}

		// Check if the prev token is an operator
		if i > 0 && token.Type == OPERATOR_TOKEN {
			prev := tokens[i-1]
			if prev.Type == OPERATOR_TOKEN {
				errs = append(errs, &ParseError{"Unexpected operator", token.Position, token.Value, "number or '('"})
			}
			continue
		}

		// If the token is a '(' push it to the stack
		if token.Type == LPAREN_TOKEN {
			parentheseStack = append(parentheseStack, token)
			continue
		}

		// If the token is a ')' pop the last '(' from the stack
		if token.Type == RPAREN_TOKEN {
			if len(parentheseStack) > 0 {
				parentheseStack = parentheseStack[:len(parentheseStack)-1]
			} else {
				errs = append(errs, &ParseError{"parse error at position %d: Unmatched closing parentheses", token.Position, "", ""})
			}
		}
	}

	if len(parentheseStack) > 0 {
		for _, token := range parentheseStack {
			errs = append(errs, &ParseError{"parse error at position %d: Unmatched opening parentheses", token.Position, "", ""})
		}
	}
	return errs
}
