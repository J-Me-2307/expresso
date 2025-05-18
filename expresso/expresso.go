package expresso

import (
	"errors"

	"github.com/J-Me-2307/expresso/internal"
)

type ValidationError struct {
	Errors []error
}

func (v ValidationError) Error() string {
	return errors.Join(v.Errors...).Error()
}

func (v ValidationError) Unwrap() error {
	return errors.Join(v.Errors...)
}

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
