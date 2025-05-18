package internal

import (
	"fmt"
	"strconv"
)

func ToPostfix(tokens []*Token) []*Token {
	queue := newQueue()
	stack := newStack()

	for _, token := range tokens {
		switch token.Type {
		case NUMBER_TOKEN:
			queue.enqueue(token)

		case OPERATOR_TOKEN:
			for !stack.isEmpty() {
				top := stack.peek()
				if top.Type != OPERATOR_TOKEN || getOperatorPrecedence(top.Value) < getOperatorPrecedence(token.Value) {
					break
				}
				queue.enqueue(stack.pop())
			}
			stack.push(token)
		case LPAREN_TOKEN:
			stack.push(token)

		case RPAREN_TOKEN:
			for !stack.isEmpty() {
				top := stack.pop()
				if top.Type == LPAREN_TOKEN {
					break
				}
				queue.enqueue(top)
			}
		}
	}

	for !stack.isEmpty() {
		queue.enqueue(stack.pop())
	}
	return queue.Values
}

func getOperatorPrecedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}

	return 0
}

func EvaluatePostfix(tokens []*Token) (float64, error) {
	stack := newStack()
	queue := &queue{Values: tokens}

	for !queue.isEmpty() {
		current := queue.dequeue()
		switch current.Type {
		case NUMBER_TOKEN:
			stack.push(current)
		case OPERATOR_TOKEN:
			rightToken := stack.pop()
			leftToken := stack.pop()

			if rightToken == nil {
				rightToken = NewToken(NUMBER_TOKEN, "0", 0)
			}
			if leftToken == nil {
				leftToken = NewToken(NUMBER_TOKEN, "0", 0)
			}

			a, _ := strconv.ParseFloat(leftToken.Value, 64)
			b, _ := strconv.ParseFloat(rightToken.Value, 64)

			var result float64
			switch current.Value {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("cannot divide by zero")
				}
				result = a / b
			}

			stack.push(NewToken(NUMBER_TOKEN, strconv.FormatFloat(result, 'f', -1, 64), leftToken.Position))
		}
	}

	res, _ := strconv.ParseFloat(stack.pop().Value, 64)
	return res, nil
}
