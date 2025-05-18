package internal

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
