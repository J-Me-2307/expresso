package internal

import (
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
				if top.Type == OPERATOR_TOKEN || getOperatorPrecedence(top.Value) < getOperatorPrecedence(token.Value) {
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

func EvaluatePostfix(tokens []*Token) float64 {
	stack := newStack()
	queue := &queue{Values: tokens}

	for !queue.isEmpty() {
		current := queue.dequeue()
		switch current.Type {
		case NUMBER_TOKEN:
			stack.push(current)
		case OPERATOR_TOKEN:
			leftToken := stack.pop()
			rightToken := stack.pop()
			a, _ := strconv.ParseFloat(leftToken.Value, 64)
			b, _ := strconv.ParseFloat(rightToken.Value, 64)
			var c float64
			switch current.Value {
			case "+":
				c = b + a
			case "-":
				c = b - a
			case "*":
				c = b * a
			case "/":
				c = b / a
			}
			stack.push(NewToken(NUMBER_TOKEN, strconv.FormatFloat(c, 'E', -1, 64), rightToken.Position))
		}
	}

	res, _ := strconv.ParseFloat(stack.pop().Value, 64)
	return res
}
