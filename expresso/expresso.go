package expresso

import (
	"errors"

	"github.com/J-Me-2307/expresso/internal"
)

// ValidationError aggregates multiple validation errors encountered during token validation.
type ValidationError struct {
	Errors []error
}

// Error returns a concatenated error message of all validation errors.
func (v ValidationError) Error() string {
	return errors.Join(v.Errors...).Error()
}

// Unwrap allows unwrapping the underlying combined errors for use with errors.Is and errors.As.
func (v ValidationError) Unwrap() error {
	return errors.Join(v.Errors...)
}

// Evaluate parses and evaluates a mathematical expression given as a string.
//
// Steps:
//  1. Tokenizes the input expression.
//  2. Validates the generated tokens. If validation fails, returns a ValidationError.
//  3. Converts the tokens to postfix notation (Reverse Polish Notation).
//  4. Evaluates the postfix expression and returns the result.
//
// Returns:
//   - The calculated result as a float64.
//   - An error if token validation fails or evaluation encounters an issue (e.g., division by zero).
func Evaluate(expression string) (float64, error) {
	tokens := internal.Tokenize(expression)

	if errs := internal.ValidateTokens(tokens); len(errs) > 0 {
		return 0, ValidationError{errs}
	}

	postfixExpression := internal.ToPostfix(tokens)

	res, err := internal.EvaluatePostfix(postfixExpression)
	if err != nil {
		return 0, err
	}

	return res, nil
}
