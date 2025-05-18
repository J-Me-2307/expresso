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

type queue struct {
	Values []*Token
}

func newQueue() *queue {
	return &queue{Values: make([]*Token, 0)}
}

func (q *queue) enqueue(token *Token) {
	q.Values = append(q.Values, token)
}

func (q *queue) dequeue() *Token {
	if q.isEmpty() {
		return nil
	}
	t := q.Values[0]
	q.Values = q.Values[1:]
	return t
}

func (q *queue) isEmpty() bool {
	return len(q.Values) == 0
}

func (q *queue) length() int {
	return len(q.Values)
}

type stack struct {
	Values []*Token
}

func newStack() *stack {
	return &stack{Values: make([]*Token, 0)}
}

func (s *stack) push(token *Token) {
	s.Values = append(s.Values, token)
}

func (s *stack) pop() *Token {
	if s.isEmpty() {
		return nil
	}
	length := s.length()
	t := s.Values[length-1]
	s.Values = s.Values[:length-1]
	return t
}

func (s *stack) peek() *Token {
	if s.isEmpty() {
		return nil
	}
	return s.Values[s.length()-1]
}

func (s *stack) isEmpty() bool {
	return len(s.Values) == 0
}

func (s *stack) length() int {
	return len(s.Values)
}
