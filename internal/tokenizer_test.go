package internal_test

import (
	"testing"

	"github.com/J-Me-2307/expresso/internal"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   []internal.Token
	}{
		{
			name:       "Simple Addition",
			expression: "3+4",
			expected: []internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "3", Position: 1},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 2},
				{Type: internal.NUMBER_TOKEN, Value: "4", Position: 3},
			},
		},
		{
			name:       "Negative Number",
			expression: "-5+2",
			expected: []internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "-5", Position: 2},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 3},
				{Type: internal.NUMBER_TOKEN, Value: "2", Position: 4},
			},
		},
		{
			name:       "Parentheses and Multiplication",
			expression: "(2+3)*4",
			expected: []internal.Token{
				{Type: internal.LPAREN_TOKEN, Value: "(", Position: 1},
				{Type: internal.NUMBER_TOKEN, Value: "2", Position: 2},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 3},
				{Type: internal.NUMBER_TOKEN, Value: "3", Position: 4},
				{Type: internal.RPAREN_TOKEN, Value: ")", Position: 5},
				{Type: internal.OPERATOR_TOKEN, Value: "*", Position: 6},
				{Type: internal.NUMBER_TOKEN, Value: "4", Position: 7},
			},
		},
		{
			name:       "Decimal Number",
			expression: "3.14+2.5",
			expected: []internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "3.14", Position: 4},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 5},
				{Type: internal.NUMBER_TOKEN, Value: "2.5", Position: 8},
			},
		},
		{
			name:       "Invalid Double Dot",
			expression: "3..14",
			expected: []internal.Token{
				{Type: internal.INVALID_TOKEN, Value: "3..14", Position: 5},
			},
		},
		{
			name:       "Negative Decimal",
			expression: "-3.5",
			expected: []internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "-3.5", Position: 4},
			},
		},
		{
			name:       "Trailing Operator",
			expression: "5+",
			expected: []internal.Token{
				{Type: internal.NUMBER_TOKEN, Value: "5", Position: 1},
				{Type: internal.OPERATOR_TOKEN, Value: "+", Position: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := internal.Tokenize(tt.expression)
			if len(tokens) != len(tt.expected) {
				t.Fatalf("expected %d tokens, got %d", len(tt.expected), len(tokens))
			}

			for i, token := range tokens {
				exp := tt.expected[i]
				if token.Type != exp.Type || token.Value != exp.Value || token.Position != exp.Position {
					t.Errorf("token %d mismatch:\n  got  {Type: %s, Value: %s, Position: %d}\n  want {Type: %s, Value: %s, Position: %d}",
						i, token.Type, token.Value, token.Position, exp.Type, exp.Value, exp.Position)
				}
			}
		})
	}
}
